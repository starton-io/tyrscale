package service

import (
	"context"
	"errors"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/starton-io/tyrscale/gateway/pkg/balancer"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/pubsub"
	"github.com/starton-io/tyrscale/go-kit/pkg/tracer"
	"github.com/starton-io/tyrscale/manager/api/proxy/route/dto"
	recommendationDto "github.com/starton-io/tyrscale/manager/api/recommendation/dto"

	"github.com/starton-io/tyrscale/manager/api/proxy/route/repository"
	recommendationRepository "github.com/starton-io/tyrscale/manager/api/recommendation/repository"
	pb "github.com/starton-io/tyrscale/manager/pkg/pb/route"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/proto"
)

var (
	routeDeletedTopic = "route_deleted"
	routeUpdatedTopic = "route_updated"
	routeCreatedTopic = "route_created"
)

type RouteService struct {
	repo               repository.IRouteRepository
	recommendationRepo recommendationRepository.IRecommendationRepository
	pub                pubsub.IPub
}

type IRouteService interface {
	Create(ctx context.Context, req *dto.CreateRouteReq) (*dto.CreateRouteRes, error)
	//AttachPlugin(ctx context.Context, routeUUID string, req *dto.AttachPluginReq) error
	//DetachPlugin(ctx context.Context, routeUUID string, req *dto.DetachPluginReq) error
	//ListPlugins(ctx context.Context, routeUUID string) (*dto.Plugins, error)
	List(ctx context.Context, req *dto.ListRouteReq) ([]*pb.RouteModel, error)
	Update(ctx context.Context, routeUUID string, req *dto.UpdateRouteReq) error
	Delete(ctx context.Context, req *dto.DeleteRouteReq) error
}

func NewRouteService(repo repository.IRouteRepository, recommendationRepo recommendationRepository.IRecommendationRepository, pub pubsub.IPub) *RouteService {
	return &RouteService{repo: repo, recommendationRepo: recommendationRepo, pub: pub}
}

func (s *RouteService) Create(ctx context.Context, req *dto.CreateRouteReq) (*dto.CreateRouteRes, error) {
	if err := req.LoadBalancerStrategy.Validate(); err != nil {
		return nil, err
	}
	routeUuid := uuid.New().String()

	ctx, span := tracer.Start(ctx, "UpsertRouteService", trace.WithAttributes(attribute.String("app.route.uuid", routeUuid)))
	defer tracer.SafeEndSpan(span)

	// check if there is a route with the same host and path
	routes, err := s.repo.List(ctx, &dto.ListRouteReq{Host: req.Host, Path: req.Path})
	if err != nil {
		return nil, err
	}
	if len(routes) > 0 {
		return nil, errors.New("route already exists")
	}

	route := &pb.RouteModel{
		Uuid:                 routeUuid,
		Host:                 req.Host,
		Path:                 req.Path,
		LoadBalancerStrategy: req.LoadBalancerStrategy.String(),
	}

	if req.CircuitBreaker != nil {
		route.CircuitBreaker = &pb.CircuitBreaker{
			Enabled:     req.CircuitBreaker.Enabled,
			MaxRequests: req.CircuitBreaker.MaxRequests,
			Interval:    req.CircuitBreaker.Interval,
			Timeout:     req.CircuitBreaker.Timeout,
		}
	}

	if req.HealthCheck != nil {
		route.HealthCheck = &pb.HealthCheck{
			Type:                       string(req.HealthCheck.Type),
			Interval:                   req.HealthCheck.Interval,
			CombinedWithCircuitBreaker: req.HealthCheck.CombinedWithCircuitBreaker,
			Enabled:                    req.HealthCheck.Enabled,
			Timeout:                    req.HealthCheck.Timeout,
		}
		if req.HealthCheck.Request != nil {
			route.HealthCheck.Request = &pb.Request{
				Method:     req.HealthCheck.Request.Method,
				StatusCode: uint32(req.HealthCheck.Request.StatusCode),
				Headers:    req.HealthCheck.Request.Headers,
				Body:       req.HealthCheck.Request.Body,
			}
		}
	}

	if err := s.repo.Upsert(ctx, route); err != nil {
		return nil, err
	}
	routeBytes, err := proto.Marshal(route)
	if err != nil {
		return nil, err
	}
	msg := message.NewMessage(uuid.New().String(), routeBytes)
	err = s.pub.Publish(ctx, routeCreatedTopic, msg)
	if err != nil {
		return nil, err
	}

	return &dto.CreateRouteRes{Uuid: route.Uuid}, nil
}

//func (s *RouteService) AttachPlugin(ctx context.Context, routeUUID string, req *dto.AttachPluginReq) error {
//	routes, err := s.List(ctx, &dto.ListRouteReq{Uuid: routeUUID})
//	if err != nil {
//		return err
//	}
//	if len(routes) == 0 {
//		return errors.New("route not found")
//	}
//	route := routes[0]
//
//	// Validate the plugin
//	jsonConfig, err := json.Marshal(req.Config)
//	if err != nil {
//		return err
//	}
//	_, err = s.pluginService.Validate(ctx, &pbPlugin.ValidatePluginRequest{
//		Name:    req.Name,
//		Type:    string(req.Type),
//		Payload: jsonConfig,
//	})
//	if err != nil {
//		return err
//	}
//
//	//pluginList, err := s.ListPlugins(ctx, routeUUID)
//	//if err != nil {
//	//	return err
//	//}
//
//	//cachePlugins := pluginList.Plugins[string(req.Type)]
//	reqPlugin := &dto.Plugin{
//		Name:        req.Name,
//		Description: req.Description,
//		Config:      req.Config,
//		Priority:    req.Priority,
//	}
//
//	switch req.Type {
//	case plugin.PluginTypeMiddleware:
//		route.Plugins = &pb.Plugins{
//			Middleware: make([]*pb.Plugin, 0),
//		}
//		if err := s.addPluginsv2(reqPlugin, &route.Plugins.Middleware); err != nil {
//			return err
//		}
//	case plugin.PluginTypeRequestInterceptor:
//		route.Plugins = &pb.Plugins{
//			RequestInterceptor: make([]*pb.Plugin, 0),
//		}
//		if err := s.addPluginsv2(reqPlugin, &route.Plugins.RequestInterceptor); err != nil {
//			return err
//		}
//	case plugin.PluginTypeResponseInterceptor:
//		route.Plugins = &pb.Plugins{
//			ResponseInterceptor: make([]*pb.Plugin, 0),
//		}
//		if err := s.addPluginsv2(reqPlugin, &route.Plugins.ResponseInterceptor); err != nil {
//			return err
//		}
//	default:
//		return errors.New("plugin type not supported")
//	}
//
//	// Update the route with the new plugin
//	if err := s.repo.Upsert(ctx, route); err != nil {
//		return err
//	}
//
//	//jsonConfig, err = json.Marshal(req.Config)
//	//if err != nil {
//	//	return err
//	//}
//
//	pluginPublish := &pb.PublishPlugin{
//		RouteHost:         route.Host,
//		RoutePath:         route.Path,
//		PluginDescription: req.Description,
//		PluginConfig:      jsonConfig,
//		PluginName:        req.Name,
//		PluginType:        string(req.Type),
//		PluginPriority:    int32(req.Priority),
//	}
//
//	pluginPublishBytes, err := proto.Marshal(pluginPublish)
//	if err != nil {
//		return err
//	}
//	msg := message.NewMessage(uuid.New().String(), pluginPublishBytes)
//	return s.pub.Publish(ctx, pluginAttachedTopic, msg)
//}

//func (s *RouteService) DetachPlugin(ctx context.Context, routeUUID string, req *dto.DetachPluginReq) error {
//	routes, err := s.List(ctx, &dto.ListRouteReq{Uuid: routeUUID})
//	if err != nil {
//		return err
//	}
//	if len(routes) == 0 {
//		return errors.New("route not found")
//	}
//	route := routes[0]
//
//	switch req.PluginType {
//	case plugin.PluginTypeMiddleware:
//		if err := s.removePluginsv2(&dto.Plugin{Name: req.PluginName}, &route.Plugins.Middleware); err != nil {
//			return err
//		}
//	case plugin.PluginTypeRequestInterceptor:
//		if err := s.removePluginsv2(&dto.Plugin{Name: req.PluginName}, &route.Plugins.RequestInterceptor); err != nil {
//			return err
//		}
//	case plugin.PluginTypeResponseInterceptor:
//		if err := s.removePluginsv2(&dto.Plugin{Name: req.PluginName}, &route.Plugins.ResponseInterceptor); err != nil {
//			return err
//		}
//	default:
//		return errors.New("plugin type not supported")
//	}
//
//	// Update the route with the new plugin
//	if err := s.repo.Upsert(ctx, route); err != nil {
//		return err
//	}
//
//	pluginPublish := &pb.PublishPlugin{
//		RouteHost:  route.Host,
//		RoutePath:  route.Path,
//		PluginName: req.PluginName,
//		PluginType: string(req.PluginType),
//	}
//
//	pluginPublishBytes, err := proto.Marshal(pluginPublish)
//	if err != nil {
//		return err
//	}
//	msg := message.NewMessage(uuid.New().String(), pluginPublishBytes)
//	return s.pub.Publish(ctx, pluginDetachedTopic, msg)
//}

//func (s *RouteService) handlePlugins(ctx context.Context, req *dto.Route, route *pb.RouteModel) error {
//	pluginList, err := s.pluginService.List(ctx)
//	if err != nil {
//		return err
//	}
//	cachePlugins := make(map[plugin.PluginType][]string)
//	for pluginType, pluginValue := range pluginList.Plugins {
//		switch pluginType {
//		case string(plugin.PluginTypeMiddleware):
//			cachePlugins[plugin.PluginTypeMiddleware] = append(cachePlugins[plugin.PluginTypeMiddleware], pluginValue.GetNames()...)
//		case string(plugin.PluginTypeRequestInterceptor):
//			cachePlugins[plugin.PluginTypeRequestInterceptor] = append(cachePlugins[plugin.PluginTypeRequestInterceptor], pluginValue.GetNames()...)
//		case string(plugin.PluginTypeResponseInterceptor):
//			cachePlugins[plugin.PluginTypeResponseInterceptor] = append(cachePlugins[plugin.PluginTypeResponseInterceptor], pluginValue.GetNames()...)
//		}
//	}
//
//	if req.Plugins != nil {
//		if err := s.addPlugins(req.Plugins.Middleware, cachePlugins[plugin.PluginTypeMiddleware], &route.Plugins.Middleware); err != nil {
//			return err
//		}
//		if err := s.addPlugins(req.Plugins.InterceptorResponse, cachePlugins[plugin.PluginTypeResponseInterceptor], &route.Plugins.InterceptorResponse); err != nil {
//			return err
//		}
//		if err := s.addPlugins(req.Plugins.InterceptorRequest, cachePlugins[plugin.PluginTypeRequestInterceptor], &route.Plugins.InterceptorRequest); err != nil {
//			return err
//		}
//	}
//
//	return nil
//}

//func (s *RouteService) addPluginsv2(reqPlugins *dto.Plugin, routePlugins *[]*pb.Plugin) error {
//	jsonConfig, err := json.Marshal(reqPlugins.Config)
//	if err != nil {
//		return err
//	}
//	*routePlugins = append(*routePlugins, &pb.Plugin{Name: reqPlugins.Name, Description: reqPlugins.Description, Config: jsonConfig, Priority: uint32(reqPlugins.Priority)})
//	return nil
//}
//
//func (s *RouteService) addPlugins(reqPlugins []*dto.Plugin, cachePlugins []string, routePlugins *[]*pb.Plugin) error {
//	for _, reqPlugin := range reqPlugins {
//		found := false
//		for _, cachePlugin := range cachePlugins {
//			if cachePlugin == reqPlugin.Name {
//				jsonConfig, err := json.Marshal(reqPlugin.Config)
//				if err != nil {
//					return err
//				}
//				*routePlugins = append(*routePlugins, &pb.Plugin{Name: reqPlugin.Name, Description: reqPlugin.Description, Config: jsonConfig, Priority: uint32(reqPlugin.Priority)})
//				found = true
//				break
//			}
//		}
//		if !found {
//			return errors.New("plugin not found: " + reqPlugin.Name)
//		}
//	}
//	return nil
//}
//
//func (s *RouteService) removePluginsv2(reqPlugin *dto.Plugin, routePlugins *[]*pb.Plugin) error {
//	for i, routePlugin := range *routePlugins {
//		if routePlugin.Name == reqPlugin.Name {
//			*routePlugins = append((*routePlugins)[:i], (*routePlugins)[i+1:]...)
//			return nil
//		}
//	}
//	return errors.New("plugin not found: " + reqPlugin.Name)
//}
//
//func (s *RouteService) removePlugins(reqPlugins []*dto.Plugin, cachePlugins []string, routePlugins *[]*pb.Plugin) error {
//	for _, reqPlugin := range reqPlugins {
//		found := false
//		for _, cachePlugin := range cachePlugins {
//			if cachePlugin == reqPlugin.Name {
//				// remove the plugin from the route
//				found = true
//				break
//			}
//		}
//		if !found {
//			return errors.New("plugin not found: " + reqPlugin.Name)
//		}
//		// delete the plugin from the route
//		for i, routePlugin := range *routePlugins {
//			if routePlugin.Name == reqPlugin.Name {
//				*routePlugins = append((*routePlugins)[:i], (*routePlugins)[i+1:]...)
//				break
//			}
//		}
//
//	}
//	return nil
//}

func (s *RouteService) Delete(ctx context.Context, req *dto.DeleteRouteReq) error {
	ctx, span := tracer.Start(ctx, "DeleteRouteService", trace.WithAttributes(attribute.String("app.route.uuid", req.Uuid)))
	defer tracer.SafeEndSpan(span)
	routes, err := s.repo.List(ctx, &dto.ListRouteReq{Uuid: req.Uuid})
	if err != nil {
		return err
	}
	if len(routes) == 0 {
		return errors.New("route not found")
	}
	recommendations, err := s.recommendationRepo.List(ctx, &recommendationDto.ListRecommendationReq{RouteUuid: req.Uuid})
	if err != nil {
		return err
	}
	if len(recommendations) > 0 {
		return errors.New("cannot delete route, it is associated to a recommendation")
	}
	err = s.repo.Delete(ctx, routes[0])
	if err != nil {
		return err
	}
	routeBytes, err := proto.Marshal(routes[0])
	if err != nil {
		return err
	}
	msg := message.NewMessage(uuid.New().String(), routeBytes)
	return s.pub.Publish(ctx, routeDeletedTopic, msg)
}

func (s *RouteService) List(ctx context.Context, req *dto.ListRouteReq) ([]*pb.RouteModel, error) {
	ctx, span := tracer.Start(ctx, "ListRouteService", trace.WithAttributes(attribute.String("app.route.uuid", req.Uuid)))
	defer tracer.SafeEndSpan(span)
	routes, err := s.repo.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return routes, nil
}

//func (s *RouteService) ListPlugins(ctx context.Context, routeUUID string) (*dto.Plugins, error) {
//	ctx, span := tracer.Start(ctx, "ListPluginsRouteService")
//	defer tracer.SafeEndSpan(span)
//	routes, err := s.repo.List(ctx, &dto.ListRouteReq{Uuid: routeUUID})
//	if err != nil {
//		return nil, err
//	}
//	plugins := &dto.Plugins{
//		InterceptorRequest:  make([]*dto.Plugin, 0),
//		InterceptorResponse: make([]*dto.Plugin, 0),
//		Middleware:          make([]*dto.Plugin, 0),
//	}
//	for _, route := range routes {
//		if route.Plugins != nil {
//			if route.Plugins.RequestInterceptor != nil {
//				for _, plugin := range route.Plugins.RequestInterceptor {
//					var dynamicConfig map[string]interface{}
//					if err := json.Unmarshal(plugin.Config, &dynamicConfig); err != nil {
//						return nil, fmt.Errorf("failed to unmarshal plugin config: %w", err)
//					}
//
//					plugins.InterceptorRequest = append(plugins.InterceptorRequest, &dto.Plugin{
//						Name:        plugin.Name,
//						Description: plugin.Description,
//						Config:      dynamicConfig,
//						Priority:    int(plugin.Priority),
//					})
//				}
//			}
//			if route.Plugins.ResponseInterceptor != nil {
//				for _, plugin := range route.Plugins.ResponseInterceptor {
//					var dynamicConfig map[string]interface{}
//					if err := json.Unmarshal(plugin.Config, &dynamicConfig); err != nil {
//						return nil, fmt.Errorf("failed to unmarshal plugin config: %w", err)
//					}
//					plugins.InterceptorResponse = append(plugins.InterceptorResponse, &dto.Plugin{
//						Name:        plugin.Name,
//						Description: plugin.Description,
//						Config:      dynamicConfig,
//						Priority:    int(plugin.Priority),
//					})
//				}
//			}
//			if route.Plugins.Middleware != nil {
//				for _, plugin := range route.Plugins.Middleware {
//					var dynamicConfig map[string]interface{}
//					if err := json.Unmarshal(plugin.Config, &dynamicConfig); err != nil {
//						return nil, fmt.Errorf("failed to unmarshal plugin config: %w", err)
//					}
//					plugins.Middleware = append(plugins.Middleware, &dto.Plugin{
//						Name:        plugin.Name,
//						Description: plugin.Description,
//						Config:      dynamicConfig,
//						Priority:    int(plugin.Priority),
//					})
//				}
//			}
//		}
//	}
//
//	// Convert the plugin configs to dynamic JSON
//	//var pluginResp *dto.Plugins
//	//for _, plugin := range plugins {
//	//	var dynamicConfig map[string]interface{}
//	//	if err := json.Unmarshal(plugin.Config, &dynamicConfig); err != nil {
//	//		return nil, fmt.Errorf("failed to unmarshal plugin config: %w", err)
//	//	}
//
//	//	// Convert the dynamic config back to JSON
//	//	//dynamicJSON, err := json.Marshal(dynamicConfig)
//	//	//if err != nil {
//	//	//	return nil, fmt.Errorf("failed to marshal dynamic config: %w", err)
//	//	//}
//	//	//pluginResp = append(pluginResp, &dto.Plugin{
//	//	//	Name:        plugin.Name,
//	//	//	Description: plugin.Description,
//	//	//	Config:      dynamicConfig,
//	//	//	Priority:    int(plugin.Priority),
//	//	//})
//	//}
//	return plugins, nil
//}

func (s *RouteService) Update(ctx context.Context, routeUUID string, req *dto.UpdateRouteReq) error {
	ctx, span := tracer.Start(ctx, "UpdateRouteService", trace.WithAttributes(attribute.String("app.route.routeUUID", routeUUID)))
	defer tracer.SafeEndSpan(span)

	routes, err := s.repo.List(ctx, &dto.ListRouteReq{Uuid: routeUUID})
	if err != nil {
		return err
	}
	if len(routes) == 0 {
		return errors.New("route not found")
	}
	route := routes[0]
	err = copier.CopyWithOption(route, req, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		return err
	}
	if err := balancer.LoadBalancerStrategy(route.LoadBalancerStrategy).Validate(); err != nil {
		return err
	}
	//routeReq := &dto.Route{
	//	Plugins: req.Plugins,
	//}
	//if err := s.handlePlugins(ctx, routeReq, route); err != nil {
	//	return err
	//}

	err = s.repo.Upsert(ctx, route)
	if err != nil {
		return err
	}
	routeBytes, err := proto.Marshal(route)
	if err != nil {
		return err
	}
	msg := message.NewMessage(uuid.New().String(), routeBytes)
	return s.pub.Publish(ctx, routeUpdatedTopic, msg)
}
