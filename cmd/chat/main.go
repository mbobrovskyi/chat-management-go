package main

import (
	"context"
	"fmt"
	"github.com/mbobrovskyi/chat-management-go/config"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/application/http"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/application/pubsub"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/application/websocket"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/abstracts"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/aggregates/connection"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/infrastructure/repositories"
	"github.com/mbobrovskyi/chat-management-go/pkg/application/common"
	commonhttp "github.com/mbobrovskyi/chat-management-go/pkg/application/http"
	"github.com/mbobrovskyi/chat-management-go/pkg/infrastructure/database/redisfactory"
	"github.com/mbobrovskyi/chat-management-go/pkg/infrastructure/httpserver"
	"github.com/mbobrovskyi/chat-management-go/pkg/infrastructure/logger/logrusfactory"
	"github.com/mbobrovskyi/chat-management-go/pkg/infrastructure/pubsub/event"
	"github.com/mbobrovskyi/chat-management-go/pkg/infrastructure/pubsub/publisher"
	"github.com/mbobrovskyi/chat-management-go/pkg/infrastructure/pubsub/subscriber"
	"github.com/mbobrovskyi/chat-management-go/pkg/infrastructure/startable"
	"github.com/mbobrovskyi/chat-management-go/pkg/infrastructure/userclient"
	"github.com/mbobrovskyi/connector/pkg/connector"
	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(fmt.Errorf("error on init config: %w", err))
	}

	fileVersion, err := os.ReadFile("VERSION")
	if err != nil {
		panic(fmt.Errorf("error on read VERSION file: %w", err))
	}

	version := string(fileVersion)

	log, err := logrusfactory.NewLogger(cfg.LogLevel)
	if err != nil {
		panic(fmt.Errorf("error on init logger: %w", err))
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	defer cancel()

	var redisClient *redis.Client
	var chatPubSubChan chan event.Event

	var chatPublisher publisher.Publisher

	switch cfg.PubSubType {
	case config.MemoryPubSub:
		chatPubSubChan = make(chan event.Event)
		defer close(chatPubSubChan)
		chatPublisher = publisher.NewMemoryPublisher(chatPubSubChan)
	case config.RedisPubSub:
		redisClient, err = redisfactory.New(ctx, cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.Db)
		if err != nil {
			log.Fatal(fmt.Errorf("error on connection to redis: %w", err))
		}
		chatPublisher = publisher.NewRedisPublisher(redisClient, cfg.ChatPubSubPrefix)
	default:
		log.Fatal("invalid pub/sub type")
	}

	var chatRepository abstracts.ChatRepository
	var messageRepository abstracts.MessageRepository

	switch cfg.DBType {
	case config.MemoryDBType:
		chatRepository = repositories.NewMemoryChatRepository()
		messageRepository = repositories.NewMemoryMessageRepository()
	default:
		log.Fatal("invalid memory type")
	}

	baseErrorHandler := common.NewBaseErrorHandler(log)

	chatConnectorHandler := websocket.NewChatConnectorEventHandler(chatRepository, messageRepository, chatPublisher)
	connectorConfig := connector.Config{Logger: log, ErrorHandler: baseErrorHandler}
	chatConnector := connector.NewT[*connection.Connection](chatConnectorHandler, connectorConfig)

	chatSubscriberHandler := pubsub.NewChatSubscriberHandler(chatConnector)

	var chatSubscriber startable.Startable

	switch cfg.PubSubType {
	case config.MemoryPubSub:
		chatSubscriber = subscriber.NewMemorySubscriber(
			log, chatPubSubChan, domain.GetAllPubSubEventTypes(), chatSubscriberHandler)
	case config.RedisPubSub:
		chatSubscriber = subscriber.NewRedisSubscriber(
			log, redisClient, cfg.ChatPubSubPrefix, domain.GetAllPubSubEventTypes(), chatSubscriberHandler)
	default:
		log.Fatal("invalid pub/sub type")
	}

	userContract := userclient.NewUserContract()

	authMiddleware := commonhttp.NewAuthMiddleware(userContract)

	mainController := commonhttp.NewMainController(version)
	chatController := http.NewChatController(authMiddleware, chatRepository, messageRepository, chatConnector)

	httpServer := httpserver.NewHttpServer(
		&cfg.BaseConfig,
		log,
		commonhttp.NewErrorHandler(cfg, log).Handle,
		[]httpserver.Controller{mainController, chatController},
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
		if err := chatSubscriber.Start(ctx); err != nil {
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
