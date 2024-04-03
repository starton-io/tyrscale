package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	slogotel "github.com/samber/slog-otel"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/kv"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/pubsub"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"github.com/starton-io/tyrscale/go-kit/pkg/tracer"
	"github.com/starton-io/tyrscale/go-kit/pkg/validation"
	httpServer "github.com/starton-io/tyrscale/manager/internal/server/http"
	"github.com/starton-io/tyrscale/manager/pkg/config"
	"github.com/starton-io/tyrscale/manager/pkg/probes"
	"go.uber.org/zap"
)

//	@title			Tyrscale Manager API
//	@version		1.0
//	@description	This is the manager service for Tyrscale
//	@contact.name	API Support
//	@contact.url	https://starton.io
//	@contact.email	support@starton.io

// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisURI,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	var pubConfig pubsub.IPub
	pubConfig = pubsub.NewRedisPub(redisstream.PublisherConfig{
		Client: redisClient,
	}, pubsub.NewGlobalPrefix(cfg.RedisStreamGlobalPrefix))

	// Initialize a slice for AttrFromContext
	attrFromContext := []func(ctx context.Context) []slog.Attr{}

	// init tracer
	if cfg.OtlpEnabled {
		tp := tracer.InitTracing("manager", &tracer.OptTracer{
			Endpoint:        cfg.OtlpEndpoint,
			TracingProvider: nil,
		})
		if err := redisotel.InstrumentTracing(redisClient); err != nil {
			panic(err)
		}
		attrFromContext = append(attrFromContext, slogotel.ExtractOtelAttrFromContext([]string{}, "trace_id", "span_id"))
		pubConfig = pubsub.NewTracingPublisherDecorator(pubConfig)

		defer func() {
			if err := tp.Shutdown(context.Background()); err != nil {
				logger.Fatalf("Error shutting down tracer provider: %v", err)
			}
		}()
	}

	// init slogger
	logLevel, err := zap.ParseAtomicLevel(cfg.LogLevel)
	if err != nil {
		panic(err)
	}
	zapLogger, err := logger.NewZapLogger(logger.WithLevel(logLevel.Level()))
	if err != nil {
		panic(err)
	}

	slogger := logger.NewSLoggerFromZap(zapLogger.Logger, &logger.OptSLogger{
		ZapLevel:    logLevel.Level(),
		AttrFromCtx: attrFromContext,
	})

	kvDB, err := kv.NewRedis(redisClient, kv.WithGlobalPrefix(cfg.RedisDBGlobalPrefix))
	if err != nil {
		logger.Fatalf("failed to create redis store: %v", err)
	}

	validator := validation.New()
	probe := probes.NewHealthChecker(cfg, kvDB)

	// catch sigterm signal and close the server
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		httpServer := httpServer.NewServer(validator, kvDB,
			httpServer.WithProbes(probe),
			httpServer.WithPublisher(pubConfig),
			httpServer.WithSLogger(slogger),
		)
		if err := httpServer.Run(); err != nil {
			logger.Fatalf("failed to run http server: %v", err)
		}
	}()
	<-signalChan
	logger.Info("Graceful shutdown")
}
