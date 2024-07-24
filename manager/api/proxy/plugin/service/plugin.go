package service

import (
	"context"
	"errors"
	"log"
	"sync"

	pb "github.com/starton-io/tyrscale/gateway/proto/gen/go/plugin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	once         sync.Once
	pluginClient pb.PluginServiceClient
)

type PluginService struct {
	listenerURL string
}

type IPluginService interface {
	List(ctx context.Context) (*pb.ListPluginsResponse, error)
}

func NewPluginService(listenerURL string) *PluginService {
	return &PluginService{listenerURL: listenerURL}
}

func initClient(listenerURL string) error {
	var err error
	once.Do(func() {
		conn, err := grpc.NewClient(listenerURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("failed to connect to gRPC server at %s: %v", listenerURL, err)
			return
		}
		pluginClient = pb.NewPluginServiceClient(conn)
	})
	return err
}

func (s *PluginService) List(ctx context.Context) (*pb.ListPluginsResponse, error) {
	err := initClient(s.listenerURL)
	if err != nil {
		return nil, err
	}
	if pluginClient == nil {
		return nil, errors.New("failed to connect to gRPC server")
	}
	listPlugins, err := pluginClient.ListPlugins(ctx, &pb.ListPluginsRequest{})
	if err != nil {
		return nil, err
	}
	return listPlugins, nil
}
