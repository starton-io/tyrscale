package consumer

import (
	"fmt"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/starton-io/tyrscale/gateway/pkg/balancer"
	"github.com/starton-io/tyrscale/gateway/pkg/circuitbreaker"
	"github.com/starton-io/tyrscale/gateway/pkg/handler"
	"github.com/starton-io/tyrscale/gateway/pkg/healthcheck"
	"github.com/starton-io/tyrscale/gateway/pkg/middleware"
	"github.com/starton-io/tyrscale/gateway/pkg/plugin"
	"github.com/starton-io/tyrscale/gateway/pkg/proxy"
	"github.com/starton-io/tyrscale/gateway/pkg/reverseproxy"
	"github.com/starton-io/tyrscale/gateway/pkg/route"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	pb "github.com/starton-io/tyrscale/manager/pkg/pb/route"
	pbUpstream "github.com/starton-io/tyrscale/manager/pkg/pb/upstream"
	"google.golang.org/protobuf/proto"
)

type RouteHandler struct {
	ProxyRouter   route.IRouter
	PluginManager plugin.IPluginManager
}

func NewRouteHandler(proxyRouter route.IRouter, pluginManager plugin.IPluginManager) *RouteHandler {
	return &RouteHandler{
		ProxyRouter:   proxyRouter,
		PluginManager: pluginManager,
	}
}

func (h *RouteHandler) HandleRouteCreated(msg *message.Message) error {
	routeProxy := &pb.RouteModel{}
	if err := proto.Unmarshal(msg.Payload, routeProxy); err != nil {
		return fmt.Errorf("error unmarshalling route: %v", err)
	}
	logger.Infof("RouteCreated: %v", routeProxy)

	proxyController := proxy.NewProxyController(balancer.LoadBalancerStrategy(routeProxy.LoadBalancerStrategy), map[string]string{
		"route_uuid": routeProxy.GetUuid(),
	})

	if routeProxy.CircuitBreaker != nil && routeProxy.CircuitBreaker.Enabled {
		logger.Infof("CircuitBreaker: %v", routeProxy.CircuitBreaker)
		proxyController.CircuitBreaker = circuitbreaker.NewTwoStepCircuitBreaker(circuitbreaker.Settings{
			Name:                   routeProxy.GetUuid(),
			MaxConsecutiveFailures: routeProxy.CircuitBreaker.MaxConsecutiveFailures,
			Interval:               routeProxy.CircuitBreaker.Interval,
			Timeout:                routeProxy.CircuitBreaker.Timeout,
			MaxRequests:            routeProxy.CircuitBreaker.MaxRequests,
		})
	}
	handler, err := handler.NewFactory(proxyController)
	if err != nil {
		return fmt.Errorf("error creating handler: %v", err)
	}
	route := route.NewRoute(
		routeProxy.Uuid,
		routeProxy.Host,
		routeProxy.Path,
		reverseproxy.NewReverseProxyHandler(handler),
		proxyController,
	)

	// setup healthcheck
	if routeProxy.HealthCheck != nil && routeProxy.HealthCheck.Enabled {
		logger.Infof("HealthCheck: %v", routeProxy.HealthCheck)
		route.HealthCheckConfig = &healthcheck.HealthCheckConfig{
			Enabled:                    routeProxy.HealthCheck.Enabled,
			Type:                       healthcheck.HealthCheckType(routeProxy.HealthCheck.Type),
			CombinedWithCircuitBreaker: routeProxy.HealthCheck.CombinedWithCircuitBreaker,
			Interval:                   routeProxy.HealthCheck.Interval,
			Timeout:                    routeProxy.HealthCheck.Timeout,
		}
		if routeProxy.HealthCheck.Request != nil {
			route.HealthCheckConfig.Request = &healthcheck.Request{
				Method:     routeProxy.HealthCheck.Request.Method,
				StatusCode: uint32(routeProxy.HealthCheck.Request.StatusCode),
				Headers:    routeProxy.HealthCheck.Request.Headers,
				Body:       routeProxy.HealthCheck.Request.Body,
			}
		}
	}

	logger.Infof("Adding route: %s", route.Host+route.Path)
	return h.ProxyRouter.Upsert(route)
}

func (h *RouteHandler) HandleRouteUpdated(msg *message.Message) error {
	routeProxy := &pb.RouteModel{}
	if err := proto.Unmarshal(msg.Payload, routeProxy); err != nil {
		return fmt.Errorf("error unmarshalling route: %v", err)
	}
	logger.Infof("RouteUpdated: %v", routeProxy)
	route, err := h.ProxyRouter.GetRoute(routeProxy.Host, routeProxy.Path)
	if err != nil {
		return fmt.Errorf("error getting route: %v", err)
	}

	// update circuitbreaker if needed
	if routeProxy.CircuitBreaker != nil && routeProxy.CircuitBreaker.Enabled {
		logger.Infof("CircuitBreaker: %v", routeProxy.CircuitBreaker)
		route.ProxyController.CircuitBreaker = circuitbreaker.NewTwoStepCircuitBreaker(circuitbreaker.Settings{
			Name:                   routeProxy.GetUuid(),
			MaxConsecutiveFailures: routeProxy.CircuitBreaker.MaxConsecutiveFailures,
			Interval:               routeProxy.CircuitBreaker.Interval,
			Timeout:                routeProxy.CircuitBreaker.Timeout,
			MaxRequests:            routeProxy.CircuitBreaker.MaxRequests,
		})
	}

	// update healthcheck if needed
	if routeProxy.HealthCheck != nil && routeProxy.HealthCheck.Enabled {
		logger.Infof("HealthCheck: %v", routeProxy.HealthCheck)
		route.HealthCheckConfig = &healthcheck.HealthCheckConfig{
			Enabled:                    routeProxy.HealthCheck.Enabled,
			Type:                       healthcheck.HealthCheckType(routeProxy.HealthCheck.Type),
			CombinedWithCircuitBreaker: routeProxy.HealthCheck.CombinedWithCircuitBreaker,
			Interval:                   routeProxy.HealthCheck.Interval,
			Timeout:                    routeProxy.HealthCheck.Timeout,
		}
		if routeProxy.HealthCheck.Request != nil {
			route.HealthCheckConfig.Request = &healthcheck.Request{
				Method:     routeProxy.HealthCheck.Request.Method,
				StatusCode: uint32(routeProxy.HealthCheck.Request.StatusCode),
				Headers:    routeProxy.HealthCheck.Request.Headers,
				Body:       routeProxy.HealthCheck.Request.Body,
			}
		}
	}

	return h.ProxyRouter.Upsert(route)
}

func (h *RouteHandler) HandleRouteDeleted(msg *message.Message) error {
	routeProxy := &pb.RouteModel{}
	if err := proto.Unmarshal(msg.Payload, routeProxy); err != nil {
		return fmt.Errorf("error unmarshalling route: %v", err)
	}
	logger.Infof("Deleting route: %s", routeProxy.Host+routeProxy.Path)
	return h.ProxyRouter.Remove(routeProxy.Host, routeProxy.Path)
}

func (h *RouteHandler) HandleUpstreamUpserted(msg *message.Message) error {
	upstreamProxy := &pbUpstream.UpstreamPublishUpsertModel{}
	if err := proto.Unmarshal(msg.Payload, upstreamProxy); err != nil {
		return fmt.Errorf("error unmarshalling upstream: %v", err)
	}
	proxyController, err := h.ProxyRouter.GetProxyController(upstreamProxy.RouteHost, upstreamProxy.RoutePath)
	if err != nil {
		logger.Errorf("error getting proxy: %v", err)
		return fmt.Errorf("error getting proxy: %v", err)
	}
	_, ok := proxyController.ClientManager.GetClient(upstreamProxy.Uuid)
	if !ok {
		logger.Infof("Adding upstream: %s", upstreamProxy.Uuid)
		proxyController.AddUpstream(upstreamProxy)

		if len(proxyController.ClientManager.GetAllClients()) == 1 {
			route, err := h.ProxyRouter.GetRoute(upstreamProxy.RouteHost, upstreamProxy.RoutePath)
			if err != nil {
				logger.Errorf("error getting route: %v", err)
				return fmt.Errorf("error getting route: %v", err)
			}
			if route.HealthCheckConfig.Enabled {
				logger.Infof("Adding health check for route: %s", route.Host+route.Path)
				healthCheck, err := healthcheck.NewHealthCheck(proxyController.ClientManager, route.HealthCheckConfig)
				if proxyController.CircuitBreaker != nil && route.HealthCheckConfig.CombinedWithCircuitBreaker {
					healthCheck.SetCircuitBreaker(proxyController.CircuitBreaker)
				}
				if err != nil {
					logger.Errorf("error creating health check: %v", err)
					return fmt.Errorf("error creating health check: %v", err)
				}
				h.ProxyRouter.AddHealthCheck(route.Uuid, healthCheck)
			}
		}
		return nil
	}
	logger.Infof("Updating upstream: %s", upstreamProxy.Uuid)
	return proxyController.UpdateUpstream(upstreamProxy)
}

func (h *RouteHandler) HandleUpstreamDeleted(msg *message.Message) error {
	upstreamProxy := &pbUpstream.UpstreamPublishDeleteModel{}
	if err := proto.Unmarshal(msg.Payload, upstreamProxy); err != nil {
		return fmt.Errorf("error unmarshalling upstream: %v", err)
	}
	logger.Infof("Deleting upstream: %s, %s, %s", upstreamProxy.Uuid, upstreamProxy.RouteHost, upstreamProxy.RoutePath)
	proxyController, err := h.ProxyRouter.GetProxyController(upstreamProxy.RouteHost, upstreamProxy.RoutePath)
	if err != nil {
		return fmt.Errorf("error getting proxy: %v", err)
	}
	logger.Infof("Deleting upstream: %s", upstreamProxy.Uuid)
	proxyController.RemoveUpstream(upstreamProxy.Uuid)
	if len(proxyController.ClientManager.GetAllClients()) == 0 {
		route, err := h.ProxyRouter.GetRoute(upstreamProxy.RouteHost, upstreamProxy.RoutePath)
		if err != nil {
			return fmt.Errorf("error getting route: %v", err)
		}
		h.ProxyRouter.RemoveHealthCheck(route.Uuid)
	}
	return nil
}

func (h *RouteHandler) HandlePluginAttached(msg *message.Message) error {
	pluginPayload := &pb.PublishPlugin{}
	if err := proto.Unmarshal(msg.Payload, pluginPayload); err != nil {
		return fmt.Errorf("error unmarshalling plugin: %v", err)
	}
	route, err := h.ProxyRouter.GetRoute(pluginPayload.RouteHost, pluginPayload.RoutePath)
	if err != nil {
		return fmt.Errorf("error getting route: %v", err)
	}
	proxyController, err := h.ProxyRouter.GetProxyController(pluginPayload.RouteHost, pluginPayload.RoutePath)
	if err != nil {
		return fmt.Errorf("error getting proxy: %v", err)
	}
	logger.Infof("Adding plugin Name: %s, type: %s", pluginPayload.PluginName, pluginPayload.PluginType)

	switch pluginPayload.PluginType {
	case "Middleware":
		for _, m := range route.ListMiddleware {
			if m.Name == pluginPayload.PluginName {
				logger.Infof("Middleware already added: %s", pluginPayload.PluginName)
				return nil
			}
		}
		middlewareInterface, err := h.PluginManager.RegisterMiddleware(pluginPayload.PluginName, pluginPayload.PluginConfig)
		if err != nil {
			logger.Errorf("error registering plugin middleware %s: %w", pluginPayload.PluginName, err)
			return fmt.Errorf("error registering plugin middleware %s: %w", pluginPayload.PluginName, err)
		}
		route.ListMiddleware = append(route.ListMiddleware, &middleware.MiddlewareWithPriority{
			Name:       pluginPayload.PluginName,
			Middleware: middlewareInterface,
			Priority:   int(pluginPayload.PluginPriority),
		})
		route.SetListMiddleware(route.ListMiddleware)
		logger.Infof("Added middleware: %s", pluginPayload.PluginName)
	case "RequestInterceptor":
		//check if the interceptor is already added
		proxyInterceptor := proxyController.GetRequestsInterceptors()
		if proxyInterceptor.Has(pluginPayload.PluginName) {
			logger.Infof("Request interceptor already added: %s", pluginPayload.PluginName)
			return nil
		}
		interceptorReqInterface, err := h.PluginManager.RegisterRequestInterceptor(pluginPayload.PluginName, pluginPayload.PluginConfig)
		if err != nil {
			logger.Errorf("error registering plugin interceptor %s: %w", pluginPayload.PluginName, err)
			return fmt.Errorf("error registering plugin interceptor %s: %w", pluginPayload.PluginName, err)
		}
		proxyInterceptor.AddOrdered(interceptorReqInterface, pluginPayload.PluginName, int(pluginPayload.PluginPriority))
		logger.Infof("Added request interceptor: %s", pluginPayload.PluginName)
		proxyController.UpdateRequestInterceptors(proxyInterceptor)
	case "ResponseInterceptor":
		//check if the interceptor is already added
		proxyInterceptor := proxyController.GetResponsesInterceptors()
		if proxyInterceptor.Has(pluginPayload.PluginName) {
			logger.Infof("Response interceptor already added: %s", pluginPayload.PluginName)
			return nil
		}

		interceptorRespInterface, err := h.PluginManager.RegisterResponseInterceptor(pluginPayload.PluginName, pluginPayload.PluginConfig)
		if err != nil {
			logger.Errorf("error registering plugin interceptor %s: %w", pluginPayload.PluginName, err)
			return fmt.Errorf("error registering plugin interceptor %s: %w", pluginPayload.PluginName, err)
		}
		proxyInterceptor.AddOrdered(interceptorRespInterface, pluginPayload.PluginName, int(pluginPayload.PluginPriority))
		logger.Infof("Added response interceptor: %s", pluginPayload.PluginName)
		proxyController.UpdateResponseInterceptors(proxyInterceptor)
	}
	return nil
}

func (h *RouteHandler) HandlePluginDetached(msg *message.Message) error {
	pluginPayload := &pb.PublishPlugin{}
	if err := proto.Unmarshal(msg.Payload, pluginPayload); err != nil {
		return fmt.Errorf("error unmarshalling plugin: %v", err)
	}
	route, err := h.ProxyRouter.GetRoute(pluginPayload.RouteHost, pluginPayload.RoutePath)
	if err != nil {
		return fmt.Errorf("error getting route: %v", err)
	}
	proxyController, err := h.ProxyRouter.GetProxyController(pluginPayload.RouteHost, pluginPayload.RoutePath)
	if err != nil {
		return fmt.Errorf("error getting proxy: %v", err)
	}
	logger.Infof("Detaching plugin name: %s, type: %s", pluginPayload.PluginName, pluginPayload.PluginType)

	switch pluginPayload.PluginType {
	case "RequestInterceptor":
		proxyInterceptor := proxyController.GetRequestsInterceptors()
		proxyInterceptor.Remove(pluginPayload.PluginName)
		proxyController.UpdateRequestInterceptors(proxyInterceptor)
		logger.Infof("Removed request interceptor: %s", pluginPayload.PluginName)
	case "ResponseInterceptor":
		proxyInterceptor := proxyController.GetResponsesInterceptors()
		proxyInterceptor.Remove(pluginPayload.PluginName)
		proxyController.UpdateResponseInterceptors(proxyInterceptor)
		logger.Infof("Removed response interceptor: %s", pluginPayload.PluginName)
	case "Middleware":
		for i, m := range route.ListMiddleware {
			if m.Name == pluginPayload.PluginName {
				route.ListMiddleware = append(route.ListMiddleware[:i], route.ListMiddleware[i+1:]...)
				logger.Debugf("Removed middleware: %s", pluginPayload.PluginName)
				break
			}
		}
		logger.Debugf("route.ListMiddleware: %v", route.ListMiddleware)
		route.SetListMiddleware(route.ListMiddleware)
	}

	return nil
}
