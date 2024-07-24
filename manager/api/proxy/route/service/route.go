package service

import (
	"context"
	"errors"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/starton-io/tyrscale/gateway/pkg/balancer"
	"github.com/starton-io/tyrscale/gateway/pkg/plugin"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/pubsub"
	"github.com/starton-io/tyrscale/go-kit/pkg/tracer"
	"github.com/starton-io/tyrscale/manager/api/proxy/route/dto"
	recommendationDto "github.com/starton-io/tyrscale/manager/api/recommendation/dto"

	pluginService "github.com/starton-io/tyrscale/manager/api/proxy/plugin/service"
	"github.com/starton-io/tyrscale/manager/api/proxy/route/repository"
	recommendationRepository "github.com/starton-io/tyrscale/manager/api/recommendation/repository"
	pb "github.com/starton-io/tyrscale/manager/pkg/pb/route"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/proto"
)

var (
	routeDeletedTopic   = "route_deleted"
	routeUpdatedTopic   = "route_updated"
	routeCreatedTopic   = "route_created"
	pluginAttachedTopic = "plugin_attached"
	pluginDetachedTopic = "plugin_detached"
)

type RouteService struct {
	repo               repository.IRouteRepository
	recommendationRepo recommendationRepository.IRecommendationRepository
	pub                pubsub.IPub
	pluginService      pluginService.IPluginService
}

type IRouteService interface {
	Create(ctx context.Context, req *dto.CreateRouteReq) (*dto.CreateRouteRes, error)
	AttachPlugin(ctx context.Context, routeUUID string, req *dto.AttachPluginReq) error
	DetachPlugin(ctx context.Context, routeUUID string, req *dto.DetachPluginReq) error
	List(ctx context.Context, req *dto.ListRouteReq) ([]*pb.RouteModel, error)
	Update(ctx context.Context, routeUUID string, req *dto.UpdateRouteReq) error
	Delete(ctx context.Context, req *dto.DeleteRouteReq) error
}

func NewRouteService(repo repository.IRouteRepository, recommendationRepo recommendationRepository.IRecommendationRepository, pub pubsub.IPub, pluginService pluginService.IPluginService) *RouteService {
	return &RouteService{repo: repo, recommendationRepo: recommendationRepo, pub: pub, pluginService: pluginService}
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

func (s *RouteService) AttachPlugin(ctx context.Context, routeUUID string, req *dto.AttachPluginReq) error {
	routes, err := s.List(ctx, &dto.ListRouteReq{Uuid: routeUUID})
	if err != nil {
		return err
	}
	if len(routes) == 0 {
		return errors.New("route not found")
	}
	route := routes[0]

	// Validate the plugin
	pluginList, err := s.pluginService.List(ctx)
	if err != nil {
		return err
	}

	cachePlugins := pluginList.Plugins[string(req.PluginType)]
	reqPlugin := make([]*dto.Plugin, 0)
	reqPlugin = append(reqPlugin, &dto.Plugin{
		Name:     req.PluginName,
		Priority: req.Priority,
	})

	switch req.PluginType {
	case plugin.PluginTypeMiddleware:
		route.Plugins = &pb.Plugins{
			Middleware: make([]*pb.Plugin, 0),
		}
		if err := s.addPlugins(reqPlugin, cachePlugins.GetNames(), &route.Plugins.Middleware); err != nil {
			return err
		}
	case plugin.PluginTypeRequestInterceptor:
		route.Plugins = &pb.Plugins{
			InterceptorRequest: make([]*pb.Plugin, 0),
		}
		if err := s.addPlugins(reqPlugin, cachePlugins.GetNames(), &route.Plugins.InterceptorRequest); err != nil {
			return err
		}
	case plugin.PluginTypeResponseInterceptor:
		route.Plugins = &pb.Plugins{
			InterceptorResponse: make([]*pb.Plugin, 0),
		}
		if err := s.addPlugins(reqPlugin, cachePlugins.GetNames(), &route.Plugins.InterceptorResponse); err != nil {
			return err
		}
	default:
		return errors.New("plugin type not supported")
	}
	// Update the route with the new plugin
	if err := s.repo.Upsert(ctx, route); err != nil {
		return err
	}

	pluginPublish := &pb.PublishPlugin{
		RouteHost:      route.Host,
		RoutePath:      route.Path,
		PluginName:     req.PluginName,
		PluginType:     string(req.PluginType),
		PluginPriority: int32(req.Priority),
	}

	pluginPublishBytes, err := proto.Marshal(pluginPublish)
	if err != nil {
		return err
	}
	msg := message.NewMessage(uuid.New().String(), pluginPublishBytes)
	return s.pub.Publish(ctx, pluginAttachedTopic, msg)
}

func (s *RouteService) DetachPlugin(ctx context.Context, routeUUID string, req *dto.DetachPluginReq) error {
	routes, err := s.List(ctx, &dto.ListRouteReq{Uuid: routeUUID})
	if err != nil {
		return err
	}
	if len(routes) == 0 {
		return errors.New("route not found")
	}
	route := routes[0]

	// Validate the plugin
	pluginList, err := s.pluginService.List(ctx)
	if err != nil {
		return err
	}

	cachePlugins := pluginList.Plugins[string(req.PluginType)]

	switch req.PluginType {
	case plugin.PluginTypeMiddleware:
		if err := s.removePlugins([]*dto.Plugin{{Name: req.PluginName}}, cachePlugins.GetNames(), &route.Plugins.Middleware); err != nil {
			return err
		}
	case plugin.PluginTypeRequestInterceptor:
		if err := s.removePlugins([]*dto.Plugin{{Name: req.PluginName}}, cachePlugins.GetNames(), &route.Plugins.InterceptorRequest); err != nil {
			return err
		}
	case plugin.PluginTypeResponseInterceptor:
		if err := s.removePlugins([]*dto.Plugin{{Name: req.PluginName}}, cachePlugins.GetNames(), &route.Plugins.InterceptorResponse); err != nil {
			return err
		}
	default:
		return errors.New("plugin type not supported")
	}

	// Update the route with the new plugin
	if err := s.repo.Upsert(ctx, route); err != nil {
		return err
	}

	pluginPublish := &pb.PublishPlugin{
		RouteHost:  route.Host,
		RoutePath:  route.Path,
		PluginName: req.PluginName,
		PluginType: string(req.PluginType),
	}

	pluginPublishBytes, err := proto.Marshal(pluginPublish)
	if err != nil {
		return err
	}
	msg := message.NewMessage(uuid.New().String(), pluginPublishBytes)
	return s.pub.Publish(ctx, pluginDetachedTopic, msg)
}

func (s *RouteService) handlePlugins(ctx context.Context, req *dto.Route, route *pb.RouteModel) error {
	pluginList, err := s.pluginService.List(ctx)
	if err != nil {
		return err
	}
	cachePlugins := make(map[plugin.PluginType][]string)
	for pluginType, pluginValue := range pluginList.Plugins {
		switch pluginType {
		case string(plugin.PluginTypeMiddleware):
			cachePlugins[plugin.PluginTypeMiddleware] = append(cachePlugins[plugin.PluginTypeMiddleware], pluginValue.GetNames()...)
		case string(plugin.PluginTypeRequestInterceptor):
			cachePlugins[plugin.PluginTypeRequestInterceptor] = append(cachePlugins[plugin.PluginTypeRequestInterceptor], pluginValue.GetNames()...)
		case string(plugin.PluginTypeResponseInterceptor):
			cachePlugins[plugin.PluginTypeResponseInterceptor] = append(cachePlugins[plugin.PluginTypeResponseInterceptor], pluginValue.GetNames()...)
		}
	}

	if req.Plugins != nil {
		if err := s.addPlugins(req.Plugins.Middleware, cachePlugins[plugin.PluginTypeMiddleware], &route.Plugins.Middleware); err != nil {
			return err
		}
		if err := s.addPlugins(req.Plugins.InterceptorResponse, cachePlugins[plugin.PluginTypeResponseInterceptor], &route.Plugins.InterceptorResponse); err != nil {
			return err
		}
		if err := s.addPlugins(req.Plugins.InterceptorRequest, cachePlugins[plugin.PluginTypeRequestInterceptor], &route.Plugins.InterceptorRequest); err != nil {
			return err
		}
	}

	return nil
}

func (s *RouteService) addPlugins(reqPlugins []*dto.Plugin, cachePlugins []string, routePlugins *[]*pb.Plugin) error {
	for _, reqPlugin := range reqPlugins {
		found := false
		for _, cachePlugin := range cachePlugins {
			if cachePlugin == reqPlugin.Name {
				*routePlugins = append(*routePlugins, &pb.Plugin{Name: reqPlugin.Name, Priority: uint32(reqPlugin.Priority)})
				found = true
				break
			}
		}
		if !found {
			return errors.New("plugin not found: " + reqPlugin.Name)
		}
	}
	return nil
}

func (s *RouteService) removePlugins(reqPlugins []*dto.Plugin, cachePlugins []string, routePlugins *[]*pb.Plugin) error {
	for _, reqPlugin := range reqPlugins {
		found := false
		for _, cachePlugin := range cachePlugins {
			if cachePlugin == reqPlugin.Name {
				// remove the plugin from the route
				found = true
				break
			}
		}
		if !found {
			return errors.New("plugin not found: " + reqPlugin.Name)
		}
		// delete the plugin from the route
		for i, routePlugin := range *routePlugins {
			if routePlugin.Name == reqPlugin.Name {
				*routePlugins = append((*routePlugins)[:i], (*routePlugins)[i+1:]...)
				break
			}
		}

	}
	return nil
}

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
	routeReq := &dto.Route{
		Plugins: req.Plugins,
	}
	if err := s.handlePlugins(ctx, routeReq, route); err != nil {
		return err
	}

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
