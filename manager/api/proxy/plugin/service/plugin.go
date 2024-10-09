package service

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"sync"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"github.com/starton-io/tyrscale/gateway/pkg/plugin"
	pb "github.com/starton-io/tyrscale/gateway/proto/gen/go/plugin"
	pubsub "github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/pubsub"
	"github.com/starton-io/tyrscale/manager/api/proxy/plugin/dto"
	pluginRepo "github.com/starton-io/tyrscale/manager/api/proxy/plugin/repository"
	dtoRoute "github.com/starton-io/tyrscale/manager/api/proxy/route/dto"
	routeRepo "github.com/starton-io/tyrscale/manager/api/proxy/route/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

var (
	once         sync.Once
	pluginClient pb.PluginServiceClient
)

const (
	pluginAttachedTopic = "plugin_attached"
	pluginDetachedTopic = "plugin_detached"
)

type PluginService struct {
	listenerURL string
	routeRepo   routeRepo.IRouteRepository
	pluginRepo  pluginRepo.IPluginRepository
	publisher   pubsub.IPub
}

type IPluginService interface {
	List(ctx context.Context) (*pb.ListPluginsResponse, error)
	ListFromRoute(ctx context.Context, routeUuid string) (*dto.Plugins, error)
	AttachPlugin(ctx context.Context, routeUuid string, plugin *dto.AttachPluginReq) error
	DetachPlugin(ctx context.Context, routeUuid string, req *dto.DetachPluginReq) error
	Validate(ctx context.Context, req *pb.ValidatePluginRequest) (*pb.ValidatePluginResponse, error)
}

func NewPluginService(listenerURL string, routeRepo routeRepo.IRouteRepository, pluginRepo pluginRepo.IPluginRepository, publisher pubsub.IPub) IPluginService {
	return &PluginService{listenerURL: listenerURL, routeRepo: routeRepo, pluginRepo: pluginRepo, publisher: publisher}
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

func (s *PluginService) Validate(ctx context.Context, req *pb.ValidatePluginRequest) (*pb.ValidatePluginResponse, error) {
	err := initClient(s.listenerURL)
	if err != nil {
		return nil, err
	}
	if pluginClient == nil {
		return nil, errors.New("failed to connect to gRPC server")
	}
	_, err = pluginClient.ValidatePlugin(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ValidatePluginResponse{}, nil
}

func (s *PluginService) ListFromRoute(ctx context.Context, routeUuid string) (*dto.Plugins, error) {
	routes, err := s.routeRepo.List(ctx, &dtoRoute.ListRouteReq{
		Uuid: routeUuid,
	})
	if err != nil {
		return nil, err
	}
	if len(routes) == 0 {
		return &dto.Plugins{}, nil
	}
	plugins, err := s.pluginRepo.List(ctx, &dto.ListPluginReq{
		RouteUuid: routeUuid,
	})
	if err != nil {
		return nil, err
	}
	var res = &dto.Plugins{
		InterceptorRequest:  make([]*dto.Plugin, 0),
		Middleware:          make([]*dto.Plugin, 0),
		InterceptorResponse: make([]*dto.Plugin, 0),
	}

	for _, pluginResp := range plugins {
		var payloadConfig map[string]interface{}
		_ = json.Unmarshal(pluginResp.Config, &payloadConfig)

		switch pluginResp.Type {
		case string(plugin.PluginTypeMiddleware):
			res.Middleware = append(res.Middleware, &dto.Plugin{
				Name:        pluginResp.Name,
				Description: pluginResp.Description,
				Config:      payloadConfig,
				Priority:    int(pluginResp.Priority),
			})
		case string(plugin.PluginTypeRequestInterceptor):
			res.InterceptorRequest = append(res.InterceptorRequest, &dto.Plugin{
				Name:        pluginResp.Name,
				Description: pluginResp.Description,
				Config:      payloadConfig,
				Priority:    int(pluginResp.Priority),
			})
		case string(plugin.PluginTypeResponseInterceptor):
			res.InterceptorResponse = append(res.InterceptorResponse, &dto.Plugin{
				Name:        pluginResp.Name,
				Description: pluginResp.Description,
				Config:      payloadConfig,
				Priority:    int(pluginResp.Priority),
			})
		}
	}

	return res, nil
}

func (s *PluginService) AttachPlugin(ctx context.Context, routeUuid string, plugin *dto.AttachPluginReq) error {
	err := initClient(s.listenerURL)
	if err != nil {
		return err
	}
	if pluginClient == nil {
		return errors.New("failed to connect to gRPC server")
	}
	routes, err := s.routeRepo.List(ctx, &dtoRoute.ListRouteReq{
		Uuid: routeUuid,
	})
	if err != nil {
		return err
	}
	if len(routes) == 0 {
		return errors.New("route not found")
	}

	payload, err := json.Marshal(plugin.Config)
	if err != nil {
		return err
	}
	req := &pb.ValidatePluginRequest{
		Name:    plugin.Name,
		Payload: payload,
		Type:    string(plugin.Type),
	}
	_, err = pluginClient.ValidatePlugin(ctx, req)
	if err != nil {
		return err
	}
	err = s.pluginRepo.Upsert(ctx, routeUuid, &pb.Plugin{
		Name:        plugin.Name,
		Config:      payload,
		Description: plugin.Description,
		Priority:    uint32(plugin.Priority),
		Type:        string(plugin.Type),
	})
	if err != nil {
		return err
	}

	publishPlugin := &pb.PublishPlugin{
		RouteHost:         routes[0].Host,
		RoutePath:         routes[0].Path,
		PluginName:        plugin.Name,
		PluginType:        string(plugin.Type),
		PluginConfig:      payload,
		PluginDescription: plugin.Description,
		PluginPriority:    int32(plugin.Priority),
	}

	payload, err = proto.Marshal(publishPlugin)
	if err != nil {
		return err
	}
	message := message.NewMessage(uuid.New().String(), payload)

	return s.publisher.Publish(ctx, pluginAttachedTopic, message)
}

func (s *PluginService) DetachPlugin(ctx context.Context, routeUuid string, req *dto.DetachPluginReq) error {
	plugins, err := s.pluginRepo.List(ctx, &dto.ListPluginReq{
		RouteUuid: routeUuid,
		Type:      req.Type,
		Name:      req.Name,
	})
	if err != nil {
		return err
	}
	if len(plugins) == 0 {
		return errors.New("plugin not found")
	}
	routes, err := s.routeRepo.List(ctx, &dtoRoute.ListRouteReq{
		Uuid: routeUuid,
	})
	if err != nil {
		return err
	}
	if len(routes) == 0 {
		return errors.New("route not found")
	}
	err = s.pluginRepo.Delete(ctx, routeUuid, &dto.DetachPluginReq{
		Name: req.Name,
		Type: req.Type,
	})
	if err != nil {
		return err
	}

	publishPlugin := &pb.PublishPlugin{
		RouteHost:  routes[0].Host,
		RoutePath:  routes[0].Path,
		PluginName: req.Name,
		PluginType: string(req.Type),
	}

	payload, err := proto.Marshal(publishPlugin)
	if err != nil {
		return err
	}
	message := message.NewMessage(uuid.New().String(), payload)

	return s.publisher.Publish(ctx, pluginDetachedTopic, message)
}
