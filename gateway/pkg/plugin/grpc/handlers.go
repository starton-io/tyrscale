package grpc

import (
	"context"

	"github.com/starton-io/tyrscale/gateway/pkg/plugin"
	pb "github.com/starton-io/tyrscale/gateway/proto/gen/go/plugin"
)

type PluginHandler struct {
	pb.UnimplementedPluginServiceServer
	service plugin.IPluginStorage
}

func NewPluginHandler(service plugin.IPluginStorage) *PluginHandler {
	return &PluginHandler{
		service: service,
	}
}

func (h *PluginHandler) ListPlugins(ctx context.Context, req *pb.ListPluginsRequest) (*pb.ListPluginsResponse, error) {
	plugins, err := h.service.ListPlugins(ctx)
	if err != nil {
		return nil, err
	}

	convertedPlugins := make(map[string]*pb.PluginList)
	for key, value := range plugins {
		convertedPlugins[key] = &pb.PluginList{Names: value}
	}

	return &pb.ListPluginsResponse{Plugins: convertedPlugins}, nil
}
