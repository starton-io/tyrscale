package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/carlmjohnson/flowmatic"
	"github.com/redis/go-redis/v9"
	"github.com/starton-io/tyrscale/gateway/internal/initializer"
	"github.com/starton-io/tyrscale/gateway/internal/server"
	"github.com/starton-io/tyrscale/gateway/internal/server/grpc"
	"github.com/starton-io/tyrscale/gateway/pkg/config"
	"github.com/starton-io/tyrscale/gateway/pkg/consumer"
	"github.com/starton-io/tyrscale/gateway/pkg/healthcheck"
	"github.com/starton-io/tyrscale/gateway/pkg/middleware"
	"github.com/starton-io/tyrscale/gateway/pkg/middleware/types"
	"github.com/starton-io/tyrscale/gateway/pkg/plugin"
	"github.com/starton-io/tyrscale/gateway/pkg/probes"
	"github.com/starton-io/tyrscale/gateway/pkg/route"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/kv"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/pubsub"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"go.uber.org/zap/zapcore"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	level, err := zapcore.ParseLevel(cfg.LogLevel)
	if err != nil {
		panic(err)
	}
	logger, err := logger.NewZapLogger(logger.WithLevel(level))
	if err != nil {
		panic(err)
	}
	cfgPlugin, err := config.LoadPluginConfig()
	if err != nil {
		panic(err)
	}
	pluginStorage := plugin.NewInMemoryPluginStorage()
	pluginManager := plugin.NewPluginManager(pluginStorage)
	_, err = pluginManager.LoadPlugins(cfgPlugin)
	if err != nil {
		panic(err)
	}
	// init grpc server
	grpcServer := grpc.NewServer(cfg, pluginStorage, pluginManager)

	router := route.NewRouter(route.WithPort(int32(cfg.ProxyHttpPort)), route.WithHealthCheckManager(healthcheck.NewHealthCheckManager()))
	listMiddleware := []types.MiddlewareFunc{middleware.HeadersMiddleware, middleware.NewLogger(logger.Logger)}
	setupMiddleware := middleware.ComposeMiddleware(listMiddleware...)

	// init redis client
	redisOptions := &redis.UniversalOptions{
		Addrs:    []string{cfg.RedisURI},
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	}
	if cfg.RedisMasterName != "" {
		redisOptions.MasterName = cfg.RedisMasterName
	}
	redisClient := redis.NewUniversalClient(redisOptions)

	// init redis pubsub
	subConfig := pubsub.NewRedisSub(redisstream.SubscriberConfig{
		ConsumerGroup: "gateway",
		Client:        redisClient,
		OldestId:      "$",
	}, pubsub.NewGlobalPrefix(cfg.RedisStreamGlobalPrefix))

	routeHandler := consumer.NewRouteHandler(router, pluginManager)
	msgRouter, err := message.NewRouter(message.RouterConfig{}, nil)
	if err != nil {
		logger.Fatalf("Failed to create new router: %v", err)
	}

	consumer := consumer.NewConsumer(routeHandler, subConfig, msgRouter)

	// Init probes
	kvDB, err := kv.NewRedis(redisClient, kv.WithGlobalPrefix(cfg.RedisDBGlobalPrefix))
	if err != nil {
		logger.Fatalf("failed to create redis store: %v", err)
	}
	probes := probes.NewHealthChecker(cfg, kvDB)

	consumer.Start(context.Background())

	// Initialize proxy
	proxyinit := initializer.NewProxyInitializer(cfg.TyrscaleApiUrl, router, pluginManager)
	err = proxyinit.Initialize(context.Background())
	if err != nil {
		logger.Fatalf("Failed to initialize proxy: %v", err)
	}

	consumer.Resume()

	svr := server.NewServer(cfg, logger.Logger,
		server.WithRouter(router),
		server.WithMiddleware(setupMiddleware),
		server.WithProbes(probes),
	)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		err := flowmatic.Do(
			func() error {
				return svr.Run()
			},
			func() error {
				return grpcServer.Run()
			},
		)
		if err != nil {
			logger.Fatalf("Failed to run Gateway : %v", err)
		}
	}()
	<-signalChan
	err = consumer.Shutdown(context.Background())
	if err != nil {
		logger.Fatalf("Failed to shutdown consumer: %v", err)
	}
}
