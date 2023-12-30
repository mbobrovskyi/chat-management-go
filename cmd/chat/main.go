package main

import (
	"context"
	"fmt"
	"github.com/mbobrovskyi/chat-management-go/internal/application/connector"
	http2 "github.com/mbobrovskyi/chat-management-go/internal/application/http"
	"github.com/mbobrovskyi/chat-management-go/internal/application/pubsub"
	"github.com/mbobrovskyi/chat-management-go/internal/application/websocket"
	"github.com/mbobrovskyi/chat-management-go/internal/domain"
	service2 "github.com/mbobrovskyi/chat-management-go/internal/domain/services"
	"github.com/mbobrovskyi/chat-management-go/internal/infrastructure/configs"
	"github.com/mbobrovskyi/chat-management-go/internal/infrastructure/contracts"
	"github.com/mbobrovskyi/chat-management-go/internal/infrastructure/database/redis"
	"github.com/mbobrovskyi/chat-management-go/internal/infrastructure/logger/logrus"
	"github.com/mbobrovskyi/chat-management-go/internal/infrastructure/pubsub/publisher"
	"github.com/mbobrovskyi/chat-management-go/internal/infrastructure/pubsub/subscriber"
	repositories2 "github.com/mbobrovskyi/chat-management-go/internal/infrastructure/repositories"
	server2 "github.com/mbobrovskyi/chat-management-go/internal/infrastructure/server"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := configs.NewConfig()
	if err != nil {
		panic(fmt.Errorf("error on init config: %w", err))
	}

	fileVersion, err := os.ReadFile("VERSION")
	if err != nil {
		panic(fmt.Errorf("error on read VERSION file: %w", err))
	}

	version := string(fileVersion)

	log, err := logrus.NewLogger(cfg.LogLevel)
	if err != nil {
		panic(fmt.Errorf("error on init logger: %w", err))
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()

	//dbConn, err := postgres.NewPostgres(ctx, cfg.PostgresUri)
	//if err != nil {
	//	log.Fatal(fmt.Errorf("error on connection to postgres: %w", err))
	//}

	redisClient, err := redis.NewRedis(ctx, cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDb)
	if err != nil {
		log.Fatal(fmt.Errorf("error on connection to redis: %w", err))
	}

	chatPublisher := publisher.NewPublisher(redisClient, cfg.ChatPubSubPrefix)

	chatRepository := repositories2.NewMemoryChatRepository()
	messageRepository := repositories2.NewMemoryMessageRepository()

	chatService := service2.NewChatService(chatRepository)
	messageService := service2.NewMessageService(messageRepository, chatPublisher)

	chatEventHandler := websocket.NewMessageEventHandler(messageService)
	chatConnector := connector.NewConnector(chatEventHandler, connector.Config{Logger: log})

	chatSubscriberHandler := pubsub.NewChatSubscriberHandler(messageService, chatConnector)
	chatSubscriber := subscriber.NewSubscriber(log, redisClient, chatSubscriberHandler, cfg.ChatPubSubPrefix)

	userContract := contracts.NewUserContract()

	mainController := http2.NewMainController(version)
	authMiddleware := http2.NewAuthMiddleware(userContract)
	chatController := http2.NewChatController(authMiddleware, chatService, messageService, chatConnector)

	httpServer := server2.NewHttpServer(
		cfg,
		log,
		http2.NewErrorHandler(cfg, log).Handle,
		[]server2.Controller{mainController, chatController},
	)

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		if err := chatConnector.Start(ctx); err != nil {
			log.Errorf("Error on running connector: %s", err.Error())
			return err
		}

		log.Info("Chat connector gracefully stopped")

		return nil
	})

	eg.Go(func() error {
		if err := chatSubscriber.Start(ctx, domain.GetAllPubSubEventTypes()); err != nil {
			log.Errorf("Error on running pubsub subscriber: %s", err.Error())
			return err
		}

		log.Info("Chat subscriber gracefully stopped")

		return nil
	})

	eg.Go(func() error {
		if err := httpServer.Start(ctx); err != nil {
			log.Errorf("Error on running http server: %s", err.Error())
			return err
		}

		log.Info("HTTP server gracefully stopped")

		return nil
	})

	if err = eg.Wait(); err != nil {
		log.Error(err)
	}
}
