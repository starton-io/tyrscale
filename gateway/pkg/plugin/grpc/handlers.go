package grpc

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/starton-io/tyrscale/gateway/pkg/plugin"
	pb "github.com/starton-io/tyrscale/gateway/proto/gen/go/plugin"
)

type PluginHandler struct {
	pb.UnimplementedPluginServiceServer
	service        plugin.IPluginStorage
	serviceManager plugin.IPluginManager
}

func NewPluginHandler(service plugin.IPluginStorage, serviceManager plugin.IPluginManager) *PluginHandler {
	return &PluginHandler{
		service:        service,
		serviceManager: serviceManager,
	}
}

func (h *PluginHandler) ListPlugins(ctx context.Context, req *pb.ListPluginsRequest) (*pb.ListPluginsResponse, error) {
	plugins, err := h.serviceManager.Store().ListPlugins(ctx)
	if err != nil {
		return nil, err
	}

	convertedPlugins := make(map[string]*pb.PluginList)
	for key, value := range plugins {
		convertedPlugins[key] = &pb.PluginList{Names: value}
	}
	//h.serviceManager.RegisterResponseInterceptor()

	return &pb.ListPluginsResponse{Plugins: convertedPlugins}, nil
}

func (h *PluginHandler) ValidatePlugin(ctx context.Context, req *pb.ValidatePluginRequest) (*empty.Empty, error) {
	err := h.serviceManager.ValidatePlugin(req.Name, req.Type, req.Payload)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
