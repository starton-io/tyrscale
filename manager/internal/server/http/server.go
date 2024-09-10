package httpfiber

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/gofiber/contrib/otelfiber/v2"
	slogfiber "github.com/samber/slog-fiber"
	"github.com/starton-io/tyrscale/go-kit/pkg/errors"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/kv"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/pubsub"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"github.com/starton-io/tyrscale/manager/pkg/config"
	"go.uber.org/zap"

	_ "github.com/starton-io/tyrscale/manager/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/starton-io/tyrscale/go-kit/pkg/validation"
	networkHttp "github.com/starton-io/tyrscale/manager/api/network/port/http"
	pluginHttp "github.com/starton-io/tyrscale/manager/api/proxy/plugin/port/http"
	routeHttp "github.com/starton-io/tyrscale/manager/api/proxy/route/port/http"
	upstreamHttp "github.com/starton-io/tyrscale/manager/api/proxy/upstream/port/http"
	recommendationHttp "github.com/starton-io/tyrscale/manager/api/recommendation/port/http"
	rpcHttp "github.com/starton-io/tyrscale/manager/api/rpc/port/http"
	"github.com/starton-io/tyrscale/manager/pkg/probes"
)

type Server struct {
	app       *fiber.App
	cfg       *config.Schema
	validator validation.Validation
	dbKV      kv.IRedisStore
	publisher pubsub.IPub
	probes    probes.HealthCheckApplication
	slogger   *slog.Logger
}

type Option func(*Server)

func NewServer(validator validation.Validation, dbKV kv.IRedisStore, opts ...Option) *Server {
	cfg := config.GetConfig()
	app := fiber.New(fiber.Config{
		ErrorHandler:          errors.CustomErrorHandler,
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		DisableStartupMessage: true,
		ReadTimeout:           time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout:          time.Duration(cfg.WriteTimeout) * time.Second,
	})

	s := &Server{
		app:       app,
		cfg:       cfg,
		validator: validator,
		dbKV:      dbKV,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func WithKVDatabase(db kv.IRedisStore) Option {
	return func(s *Server) {
		s.dbKV = db
	}
}

func WithPublisher(pub pubsub.IPub) Option {
	return func(s *Server) {
		s.publisher = pub
	}
}

func WithProbes(probes probes.HealthCheckApplication) Option {
	return func(s *Server) {
		s.probes = probes
	}
}

func WithSLogger(logger *slog.Logger) Option {
	return func(s *Server) {
		s.slogger = logger
	}
}

func (s *Server) Run() error {

	s.app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	s.app.Use(otelfiber.Middleware())
	s.app.Use(slogfiber.New(s.slogger))
	//s.app.Use(recover.New())

	if err := s.MapRoutes(); err != nil {
		logger.Fatalf("MapRoutes Error: %v", err)
	}
	s.app.Get("/swagger/*", swagger.HandlerDefault) // Adjust according to your swagger setup

	s.app.Get("/readiness", func(c *fiber.Ctx) error {
		result := s.probes.ReadyEndpoint()
		if result.Status {
			return c.Status(fiber.StatusOK).JSON(result)
		}
		return c.Status(fiber.StatusServiceUnavailable).JSON(result)
	})

	s.app.Get("/liveness", func(c *fiber.Ctx) error {
		result := s.probes.LiveEndpoint()
		if result.Status {
			return c.Status(fiber.StatusOK).JSON(result)
		}
		return c.Status(fiber.StatusServiceUnavailable).JSON(result)
	})

	// Start http server
	logger.Info("HTTP server is listening on PORT: ", zap.Int("port", s.cfg.HttpPort))
	if err := s.app.Listen(fmt.Sprintf(":%d", s.cfg.HttpPort)); err != nil {
		logger.Fatalf("Running HTTP server: %v", err)
	}

	return nil
}

func (s *Server) MapRoutes() error {
	v1 := s.app.Group("/api/v1")
	networkHttp.Routes(v1, s.dbKV, s.validator, s.publisher)
	rpcHttp.Routes(v1, s.dbKV, s.validator, s.publisher)
	recommendationHttp.Routes(v1, s.dbKV, s.validator, s.publisher)
	upstreamHttp.Routes(v1, s.dbKV, s.validator, s.publisher)
	pluginHttp.Routes(v1, s.validator, s.cfg.GatewayUrl, s.dbKV, s.publisher)
	routeHttp.Routes(v1, s.dbKV, s.validator, s.publisher, s.cfg.GatewayUrl)
	return nil
}
