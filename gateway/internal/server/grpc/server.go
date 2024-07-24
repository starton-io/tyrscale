package grpc

import (
	"fmt"
	"net"

	"github.com/starton-io/tyrscale/gateway/pkg/config"
	"github.com/starton-io/tyrscale/gateway/pkg/plugin"
	pluginGRPC "github.com/starton-io/tyrscale/gateway/pkg/plugin/grpc"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	engine        *grpc.Server
	cfg           *config.Schema
	pluginStorage plugin.IPluginStorage
}

func NewServer(cfg *config.Schema, pluginStorage plugin.IPluginStorage) *Server {
	engine := grpc.NewServer(
		grpc.ChainUnaryInterceptor(),
	)
	return &Server{
		engine:        engine,
		cfg:           cfg,
		pluginStorage: pluginStorage,
	}
}

func (s Server) Run() error {
	pluginGRPC.RegisterHandlers(s.engine, s.pluginStorage)
	reflection.Register(s.engine)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.GrpcPort))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	logger.Infof("grpc server started on port %d", s.cfg.GrpcPort)
	err = s.engine.Serve(lis)
	if err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}
