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
	"github.com/starton-io/tyrscale/gateway/pkg/config"
	"github.com/starton-io/tyrscale/gateway/pkg/consumer"
	"github.com/starton-io/tyrscale/gateway/pkg/healthcheck"
	"github.com/starton-io/tyrscale/gateway/pkg/middleware"
	"github.com/starton-io/tyrscale/gateway/pkg/route"
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

	router := route.NewRouter(route.WithPort(int32(cfg.ProxyHttpPort)), route.WithHealthCheckManager(healthcheck.NewHealthCheckManager()))
	listMiddleware := []middleware.MiddlewareFunc{middleware.HeadersMiddleware, middleware.NewLogger(logger.Logger)}
	setupMiddleware := middleware.MiddlewareComposer(listMiddleware)

	// init redis pubsub
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisURI,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})
	subConfig := pubsub.NewRedisSub(redisstream.SubscriberConfig{
		ConsumerGroup: "gateway",
		Client:        redisClient,
		OldestId:      "$",
	}, pubsub.NewGlobalPrefix(cfg.RedisStreamGlobalPrefix))

	routeHandler := consumer.NewRouteHandler(router)
	msgRouter, err := message.NewRouter(message.RouterConfig{}, nil)
	if err != nil {
		logger.Fatalf("Failed to create new router: %v", err)
	}

	consumer := consumer.NewConsumer(routeHandler, subConfig, msgRouter)

	// ---------------------------------------------------------------------------
	// -------------------- Test Load Balancer -----------------------------------
	// ---------------------------------------------------------------------------

	//proxyController := proxy.ProxyController{
	//	ClientManager: &proxy.DefaultClientManager{
	//		MapClient: make(map[string]*proxy.UpstreamClient),
	//	},
	//	Balancer: balancer.NewBalancer(balancer.BalancerPriority),
	//	CircuitBreaker: circuitbreaker.NewCircuitBreaker(circuitbreaker.Settings{
	//		Name:                   "test",
	//		MaxRequests:            1,
	//		MaxConsecutiveFailures: 5,
	//		Interval:               120,
	//		Timeout:                60,
	//	}),
	//}

	//myHandler := handler.NewFailoverHandler(proxyController)
	//proxy := reverseproxy.NewReverseProxyHandler(myHandler)

	//router.Add(&route.Route{
	//	Host:            "starton-local.com",
	//	Path:            "/test",
	//	ReverseProxy:    proxy,
	//	ProxyController: proxyController,
	//})

	//proxyController.AddUpstream(&upstream.UpstreamPublishUpsertModel{
	//	Uuid:   "1",
	//	Host:   "api.starton.io",
	//	Path:   "/v3/listener-evm/health/network",
	//	Port:   443,
	//	Scheme: "https",
	//	Weight: 10,
	//})

	//proxyController.AddUpstream(&upstream.UpstreamPublishUpsertModel{
	//	Uuid:   "2",
	//	Host:   "api.starton.io",
	//	Path:   "/v3/listener-evm/healthfdfdf",
	//	Port:   443,
	//	Scheme: "https",
	//	Weight: 100,
	//})

	//proxyController.AddUpstream(&upstream.UpstreamPublishUpsertModel{
	//	Uuid:   "3",
	//	Host:   "api.starton.titi",
	//	Path:   "/v3/listener-evm/health",
	//	Port:   443,
	//	Scheme: "https",
	//	Weight: 100,
	//})

	//proxyController.AddUpstream(&upstream.UpstreamPublishUpsertModel{
	//	Uuid:   "4",
	//	Host:   "api.starton.tutu",
	//	Path:   "/v3/listener-evm/health",
	//	Port:   443,
	//	Scheme: "https",
	//	Weight: 100,
	//})

	//proxyController.AddUpstream(&upstream.UpstreamPublishUpsertModel{
	//	Uuid:   "5",
	//	Host:   "api.starton.testfddf",
	//	Path:   "/v3/listener-evm/health",
	//	Port:   443,
	//	Scheme: "https",
	//	Weight: 100,
	//})

	//proxyController2 := &proxy.ProxyController{
	//	MapClient: make(map[string]*proxy.UpstreamClient),
	//	Balancer:  balancer.NewBalancer(balancer.BalancerTypeWeightRoundRobin),
	//}
	//proxyController2.AddUpstream(&upstream.UpstreamPublishModel{
	//	Uuid:   "3",
	//	Host:   "api.starton.com",
	//	Path:   "/v3/listener-evm/health",
	//	Port:   443,
	//	Scheme: "https",
	//	Weight: 50,
	//})
	//proxyController2.AddUpstream(&upstream.UpstreamPublishModel{
	//	Uuid:   "4",
	//	Host:   "api.starton.toto",
	//	Path:   "/v3/listener-evm/health",
	//	Port:   443,
	//	Scheme: "https",
	//	Weight: 50,
	//})

	//failoverHandler := &proxy.FailoverHandler{}
	////////defaultHandler := &proxy.DefaultProxyHandler{}

	//proxyRouter := proxy.NewProxy(failoverHandler, proxyController)
	////////proxyRouter2 := proxy.NewProxy(defaultHandler, proxyController2)

	//router.Add(&route.Route{Host: "starton-local.com", Path: "/test", Proxy: proxyRouter})
	//router.Add(&route.Route{Host: "starton-local.com", Path: "/test2", Proxy: proxyRouter2})

	consumer.Start(context.Background())

	// ---------------------------------------------------------------------------
	// ---------------------------------------------------------------------------

	// ---------------------------------------------------------------------------
	// -------------------- Initialize Proxy -------------------------------------
	// ---------------------------------------------------------------------------
	proxyinit := initializer.NewProxyInitializer(cfg.TyrscaleApiUrl, router)
	err = proxyinit.Initialize(context.Background())
	if err != nil {
		logger.Fatalf("Failed to initialize proxy: %v", err)
	}

	consumer.Resume()

	svr := server.NewServer(cfg, logger.Logger, server.WithRouter(router), server.WithMiddleware(setupMiddleware))
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		err := flowmatic.Do(
			func() error {
				return svr.Run()
			},
		)
		if err != nil {
			logger.Fatalf("Failed to run Gateway : %v", err)
		}
	}()
	<-signalChan
	consumer.Shutdown(context.Background())
}
