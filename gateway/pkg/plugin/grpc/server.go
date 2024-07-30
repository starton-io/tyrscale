package grpc

import (
	"github.com/starton-io/tyrscale/gateway/pkg/plugin"
	pb "github.com/starton-io/tyrscale/gateway/proto/gen/go/plugin"
	"google.golang.org/grpc"
)

func RegisterHandlers(server *grpc.Server, storage plugin.IPluginStorage, serviceManager plugin.IPluginManager) {
	handler := NewPluginHandler(storage, serviceManager)
	pb.RegisterPluginServiceServer(server, handler)
}
