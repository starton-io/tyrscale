package consumer

import (
	"fmt"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/starton-io/tyrscale/gateway/pkg/balancer"
	"github.com/starton-io/tyrscale/gateway/pkg/circuitbreaker"
	"github.com/starton-io/tyrscale/gateway/pkg/handler"
	"github.com/starton-io/tyrscale/gateway/pkg/healthcheck"
	"github.com/starton-io/tyrscale/gateway/pkg/proxy"
	"github.com/starton-io/tyrscale/gateway/pkg/reverseproxy"
	"github.com/starton-io/tyrscale/gateway/pkg/route"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	pb "github.com/starton-io/tyrscale/manager/pkg/pb/route"
	pbUpstream "github.com/starton-io/tyrscale/manager/pkg/pb/upstream"
	"google.golang.org/protobuf/proto"
)

type RouteHandler struct {
	ProxyRouter route.IRouter
}

func NewRouteHandler(proxyRouter route.IRouter) *RouteHandler {
	return &RouteHandler{
		ProxyRouter: proxyRouter,
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
	handler := handler.NewFailoverHandler(proxyController)
	route := route.NewRoute(
		routeProxy.Uuid,
		routeProxy.Host,
		routeProxy.Path,
		reverseproxy.NewReverseProxyHandler(handler),
		proxyController,
	)
	if routeProxy.HealthCheck != nil && routeProxy.HealthCheck.Enabled {
		logger.Infof("HealthCheck: %v", routeProxy.HealthCheck)
		route.HealthCheckConfig = &healthcheck.HealthCheckConfig{
			Enabled:                    routeProxy.HealthCheck.Enabled,
			Type:                       healthcheck.HealthCheckType(routeProxy.HealthCheck.Type),
			CombinedWithCircuitBreaker: routeProxy.HealthCheck.CombinedWithCircuitBreaker,
			Interval:                   routeProxy.HealthCheck.Interval,
			Timeout:                    routeProxy.HealthCheck.Timeout,
		}
	}

	logger.Infof("Adding route: %s", route.Host+route.Path)
	return h.ProxyRouter.Add(route)
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
				healthCheck, err := healthcheck.NewHealthCheck(route.HealthCheckConfig.Type,
					proxyController.ClientManager,
					route.HealthCheckConfig.Interval,
					route.HealthCheckConfig.Timeout,
				)
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
