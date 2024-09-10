package server

import (
	"encoding/json"
	"fmt"

	"github.com/starton-io/tyrscale/gateway/pkg/config"
	"github.com/starton-io/tyrscale/gateway/pkg/middleware"
	"github.com/starton-io/tyrscale/gateway/pkg/middleware/types"
	"github.com/starton-io/tyrscale/gateway/pkg/route"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"github.com/starton-io/tyrscale/manager/pkg/probes"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

type Server struct {
	Engine          *fasthttp.Server
	Router          route.IRouter
	ApplyMiddleware types.MiddlewareFunc
	probes          probes.HealthCheckApplication
	Cfg             *config.Schema
}

type Option func(*Server)

func NewServer(cfg *config.Schema, zapLogger *zap.Logger, opts ...Option) *Server {
	engine := &fasthttp.Server{
		Name:        cfg.ServerName,
		Concurrency: cfg.ProxyConcurrency,
	}
	s := &Server{
		Engine: engine,
		Cfg:    cfg,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func WithRouter(router route.IRouter) Option {
	return func(s *Server) {
		s.Router = router
	}
}

func WithMiddleware(middleware types.MiddlewareFunc) Option {
	return func(s *Server) {
		s.ApplyMiddleware = middleware
	}
}

func WithProbes(probes probes.HealthCheckApplication) Option {
	return func(s *Server) {
		s.probes = probes
	}
}

func (s *Server) Run() error {
	logger.Infof("Starting server on port %d", s.Cfg.ProxyHttpPort)
	routerHandler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/liveness":
			result := s.probes.LiveEndpoint()
			body, _ := json.Marshal(result)
			ctx.Response.Header.Set("Content-Type", "application/json")
			if result.Status {
				ctx.Response.SetBody(body)
				ctx.Response.SetStatusCode(fasthttp.StatusOK)
			} else {
				ctx.Response.SetBody(body)
				ctx.Response.SetStatusCode(fasthttp.StatusServiceUnavailable)
			}
		case "/readiness":
			result := s.probes.ReadyEndpoint()
			body, _ := json.Marshal(result)
			ctx.Response.Header.Set("Content-Type", "application/json")
			if result.Status {
				ctx.Response.SetBody(body)
				ctx.Response.SetStatusCode(fasthttp.StatusOK)
			} else {
				ctx.Response.SetBody(body)
				ctx.Response.SetStatusCode(fasthttp.StatusServiceUnavailable)
			}
		case "/metrics":
			middleware.PrometheusHandler()(ctx)
		default:
			s.Router.ProxyRouter(ctx)
		}
	}
	s.Engine.Handler = s.ApplyMiddleware(routerHandler)
	return s.Engine.ListenAndServe(fmt.Sprintf(":%d", s.Cfg.ProxyHttpPort))
}
