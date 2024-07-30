package server

import (
	"fmt"

	"github.com/starton-io/tyrscale/gateway/pkg/config"
	"github.com/starton-io/tyrscale/gateway/pkg/middleware"
	"github.com/starton-io/tyrscale/gateway/pkg/middleware/types"
	"github.com/starton-io/tyrscale/gateway/pkg/route"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

type Server struct {
	Engine          *fasthttp.Server
	Router          route.IRouter
	ApplyMiddleware types.MiddlewareFunc
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

func (s *Server) Run() error {
	logger.Infof("Starting server on port %d", s.Cfg.ProxyHttpPort)
	routerHandler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/metrics":
			middleware.PrometheusHandler()(ctx)
		default:
			s.Router.ProxyRouter(ctx)
		}
	}
	s.Engine.Handler = s.ApplyMiddleware(routerHandler)
	return s.Engine.ListenAndServe(fmt.Sprintf(":%d", s.Cfg.ProxyHttpPort))
}
