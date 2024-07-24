package consumer

import (
	"fmt"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/starton-io/tyrscale/gateway/pkg/balancer"
	"github.com/starton-io/tyrscale/gateway/pkg/circuitbreaker"
	"github.com/starton-io/tyrscale/gateway/pkg/handler"
	"github.com/starton-io/tyrscale/gateway/pkg/healthcheck"
	"github.com/starton-io/tyrscale/gateway/pkg/interceptor"
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
		proxyController.CircuitBreaker = circuitbreaker.NewCircuitBreaker(circuitbreaker.Settings{
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

	// setup middleware
	listMiddlewareWithPriority := make([]*middleware.MiddlewareWithPriority, 0)
	setupMiddleware := middleware.MiddlewareWithPriorityComposer(listMiddlewareWithPriority...)
	route.Middleware = setupMiddleware

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
		route.ProxyController.CircuitBreaker = circuitbreaker.NewCircuitBreaker(circuitbreaker.Settings{
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

	// update interceptorResponse if needed
	if routeProxy.Plugins != nil {
		interceptor := interceptor.NewInterceptorResponseChain()
		for _, p := range routeProxy.Plugins.InterceptorResponse {
			logger.Infof("Adding plugin response interceptor: %s", p.Name)
			interceptorResp, err := h.PluginManager.GetPluginRespInterceptor(p.Name)
			if err != nil {
				return fmt.Errorf("error getting interceptor: %v", err)
			}
			interceptor.AddOrdered(interceptorResp, int(p.Priority))
		}
		route.ProxyController.SetResponsesInterceptors(interceptor)
	}

	// update middleware if needed
	listMiddlewareWithPriority := []*middleware.MiddlewareWithPriority{}
	if routeProxy.Plugins != nil {
		for _, p := range routeProxy.Plugins.Middleware {
			logger.Infof("Adding Plugin Middleware: %s", p.Name)
			pluginMiddleware, err := h.PluginManager.GetPluginMiddleware(p.Name)
			if err != nil {
				logger.Errorf("failed to get plugin middleware %s: %w", p.Name, err)
				continue
				//return fmt.Errorf("failed to get plugin %s: %w", p.Name, err)
			}
			listMiddlewareWithPriority = append(listMiddlewareWithPriority, &middleware.MiddlewareWithPriority{
				Middleware: pluginMiddleware,
				Priority:   int(p.Priority),
			})
		}
	}
	route.SetListMiddleware(listMiddlewareWithPriority)

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

	switch pluginPayload.PluginType {
	case "Middleware":
		middlewareInterface, err := h.PluginManager.GetPluginMiddleware(pluginPayload.PluginName)
		if err != nil {
			logger.Errorf("error getting plugin middleware %s: %w", pluginPayload.PluginName, err)
			return fmt.Errorf("error getting plugin middleware %s: %w", pluginPayload.PluginName, err)
		}
		route.ListMiddleware = append(route.ListMiddleware, &middleware.MiddlewareWithPriority{
			Middleware: middlewareInterface,
			Priority:   int(pluginPayload.PluginPriority),
		})
		route.SetListMiddleware(route.ListMiddleware)
		logger.Infof("Added middleware: %s", pluginPayload.PluginName)
	case "InterceptorRequest":
		interceptorReqInterface, err := h.PluginManager.GetPluginReqInterceptor(pluginPayload.PluginName)
		if err != nil {
			logger.Errorf("error getting plugin interceptor %s: %w", pluginPayload.PluginName, err)
			return fmt.Errorf("error getting plugin interceptor %s: %w", pluginPayload.PluginName, err)
		}
		proxyInterceptor := proxyController.GetRequestsInterceptors()
		proxyInterceptor.AddOrdered(interceptorReqInterface, int(pluginPayload.PluginPriority))
		logger.Infof("Added request interceptor: %s", pluginPayload.PluginName)
	case "InterceptorResponse":
		interceptorRespInterface, err := h.PluginManager.GetPluginRespInterceptor(pluginPayload.PluginName)
		if err != nil {
			logger.Errorf("error getting plugin interceptor %s: %w", pluginPayload.PluginName, err)
			return fmt.Errorf("error getting plugin interceptor %s: %w", pluginPayload.PluginName, err)
		}
		proxyInterceptor := proxyController.GetResponsesInterceptors()
		proxyInterceptor.AddOrdered(interceptorRespInterface, int(pluginPayload.PluginPriority))
		logger.Infof("Added response interceptor: %s", pluginPayload.PluginName)
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

	switch pluginPayload.PluginType {
	case "InterceptorRequest":
		interceptorReqInterface, err := h.PluginManager.GetPluginReqInterceptor(pluginPayload.PluginName)
		if err != nil {
			logger.Errorf("error getting plugin interceptor %s: %w", pluginPayload.PluginName, err)
			return fmt.Errorf("error getting plugin interceptor %s: %w", pluginPayload.PluginName, err)
		}
		proxyInterceptor := proxyController.GetRequestsInterceptors()
		proxyInterceptor.Remove(interceptorReqInterface)
		logger.Infof("Removed request interceptor: %s", pluginPayload.PluginName)
	case "InterceptorResponse":
		interceptorRespInterface, err := h.PluginManager.GetPluginRespInterceptor(pluginPayload.PluginName)
		if err != nil {
			logger.Errorf("error getting plugin interceptor %s: %w", pluginPayload.PluginName, err)
			return fmt.Errorf("error getting plugin interceptor %s: %w", pluginPayload.PluginName, err)
		}
		proxyInterceptor := proxyController.GetResponsesInterceptors()
		proxyInterceptor.Remove(interceptorRespInterface)
		logger.Infof("Removed response interceptor: %s", pluginPayload.PluginName)
	case "Middleware":
		routeMiddleware := route.ListMiddleware
		for i, m := range routeMiddleware {
			if m.Name == pluginPayload.PluginName {
				route.ListMiddleware = append(route.ListMiddleware[:i], route.ListMiddleware[i+1:]...)
				break
			}
		}
		route.SetListMiddleware(route.ListMiddleware)
		logger.Infof("Removed middleware: %s", pluginPayload.PluginName)
	}

	return nil
}
