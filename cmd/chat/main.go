package main

import (
	"context"
	"fmt"
	chatapplication "github.com/mbobrovskyi/ddd-chat-management-go/internal/chat/api"
	chatdomain "github.com/mbobrovskyi/ddd-chat-management-go/internal/chat/domain"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/chat/repositories"
	chatpubsub "github.com/mbobrovskyi/ddd-chat-management-go/internal/chat/subscription"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/chat/websocket"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/application"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/connector"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/pubsub"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/infrastructure/configs"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/infrastructure/database/redis"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/infrastructure/logger/logrus"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/infrastructure/server"
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

	chatPublisher := pubsub.NewPublisher(redisClient, cfg.ChatPubSubPrefix)

	chatRepository := repositories.NewChatRepository()
	messageRepository := repositories.NewMessageRepository()

	chatService := chatdomain.NewService(chatRepository)
	messageService := chatdomain.NewMessageService(messageRepository, chatPublisher)

	chatEventHandler := websocket.NewMessageEventHandler(messageService)
	chatConnector := connector.NewConnector(chatEventHandler, connector.Config{Logger: log})

	chatSubscriberHandler := chatpubsub.NewChatSubscriberHandler(messageService, chatConnector)
	chatSubscriber := pubsub.NewSubscriber(log, redisClient, chatSubscriberHandler, cfg.ChatPubSubPrefix)

	mainController := application.NewMainController(version)
	chatController := chatapplication.NewChatController(chatService, chatConnector)

	httpServer := server.NewHttpServer(
		cfg,
		log,
		application.NewErrorHandler(cfg, log).Handle,
		[]server.Controller{mainController, chatController},
	)

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		if err := chatConnector.Start(ctx); err != nil {
			log.Errorf("Error on running connector: %s", err.Error())
			return err
		}

		log.Info("Connector gracefully stopped")

		return nil
	})

	eg.Go(func() error {
		if err := chatSubscriber.Subscribe(ctx, chatdomain.GetAllPubSubEventTypes()); err != nil {
			log.Errorf("Error on running pubsub subscriber: %s", err.Error())
			return err
		}

		log.Info("Connector gracefully stopped")

		return nil
	})

	eg.Go(func() error {
		if err := httpServer.Start(ctx); err != nil {
			log.Errorf("Error on running http server: %s", err.Error())
			return err
		}

		log.Info("Server gracefully stopped")

		return nil
	})

	if err = eg.Wait(); err != nil {
		log.Error(err)
	}
}
