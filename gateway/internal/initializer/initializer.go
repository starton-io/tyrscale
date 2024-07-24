package initializer

import (
	"context"
	"fmt"

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
	"github.com/starton-io/tyrscale/go-kit/pkg/ptr"
	upstreamPb "github.com/starton-io/tyrscale/manager/pkg/pb/upstream"
	tyrscaleSDK "github.com/starton-io/tyrscale/sdk/tyrscale-sdk-go"
)

type Initializer interface {
	Initialize(ctx context.Context) error
}

type ProxyInitializer struct {
	TyrscaleClient *tyrscaleSDK.APIClient
	pluginManager  plugin.IPluginManager
	router         route.IRouter
}

func NewProxyInitializer(url string, router route.IRouter, pluginManager plugin.IPluginManager) *ProxyInitializer {
	config := tyrscaleSDK.NewConfiguration()
	config.Servers = tyrscaleSDK.ServerConfigurations{
		{
			URL: url,
		},
	}
	client := tyrscaleSDK.NewAPIClient(config)
	return &ProxyInitializer{TyrscaleClient: client, router: router, pluginManager: pluginManager}
}

// TODO: Optimize the initialization process
func (i *ProxyInitializer) Initialize(ctx context.Context) error {
	routes, _, err := i.TyrscaleClient.RoutesAPI.ListRoutes(ctx).Execute()
	if err != nil {
		return fmt.Errorf("failed to list routes: %w", err)
	}
	if routes.Data == nil || len(routes.Data.Items) == 0 {
		logger.Warn("no routes found")
		return nil
	}

	for _, currentRoute := range routes.Data.Items {
		var healthCheckRoute healthcheck.HealthCheckInterface
		var healthCheckSettings *healthcheck.HealthCheckConfig
		logger.Infof("Adding Route: %s", currentRoute.GetUuid())

		listUpstream, _, err := i.TyrscaleClient.UpstreamsAPI.ListUpstreams(ctx, currentRoute.GetUuid()).Execute()
		if err != nil {
			return fmt.Errorf("failed to list upstreams for route %s: %w", currentRoute.GetUuid(), err)
		}

		proxyController := proxy.NewProxyController(balancer.LoadBalancerStrategy(currentRoute.LoadBalancerStrategy), map[string]string{
			"route_uuid": currentRoute.GetUuid(),
		})

		interceptorRespChain := interceptor.NewInterceptorResponseChain()
		if currentRoute.Plugins != nil {
			for _, p := range currentRoute.Plugins.InterceptorResponse {
				logger.Infof("Adding Plugin Interceptor Response: %s", p.Name)
				interceptorResp, err := i.pluginManager.GetPluginRespInterceptor(p.Name)
				if err != nil {
					logger.Errorf("failed to get plugin %s: %w", p.Name, err)
					//return fmt.Errorf("failed to get plugin %s: %w", p.Name, err)
					continue
				}
				interceptorRespChain.AddOrdered(interceptorResp, int(p.Priority))
			}
			proxyController.SetResponsesInterceptors(interceptorRespChain)
		}

		interceptorReqChain := interceptor.NewInterceptorRequestChain()
		if currentRoute.Plugins != nil {
			for _, p := range currentRoute.Plugins.InterceptorRequest {
				logger.Infof("Adding Plugin Interceptor Response: %s", p.Name)
				interceptorReq, err := i.pluginManager.GetPluginReqInterceptor(p.Name)
				if err != nil {
					logger.Errorf("failed to get plugin %s: %w", p.Name, err)
					//return fmt.Errorf("failed to get plugin %s: %w", p.Name, err)
					continue
				}
				interceptorReqChain.AddOrdered(interceptorReq, int(p.Priority))
			}
			proxyController.SetRequestsInterceptors(interceptorReqChain)
		}

		// Add circuit breaker if needed
		if cb := currentRoute.CircuitBreaker; cb != nil && cb.GetEnabled() {
			proxyController.CircuitBreaker = circuitbreaker.NewCircuitBreaker(circuitbreaker.Settings{
				Name:                   currentRoute.GetUuid(),
				Interval:               uint32(cb.GetInterval()),
				MaxConsecutiveFailures: uint32(cb.GetMaxConsecutiveFailures()),
				Timeout:                uint32(cb.GetTimeout()),
				MaxRequests:            uint32(cb.GetMaxRequests()),
			})
		}
		if currentRoute.HealthCheck != nil {
			healthCheckSettings = &healthcheck.HealthCheckConfig{
				Type:                       healthcheck.HealthCheckType(currentRoute.HealthCheck.GetType()),
				Interval:                   uint32(currentRoute.HealthCheck.GetInterval()),
				CombinedWithCircuitBreaker: currentRoute.HealthCheck.GetCombinedWithCircuitBreaker(),
				Enabled:                    currentRoute.HealthCheck.GetEnabled(),
				Timeout:                    uint32(currentRoute.HealthCheck.GetTimeout()),
			}
			if currentRoute.HealthCheck.Request != nil {
				healthCheckSettings.Request = &healthcheck.Request{
					Method:     currentRoute.HealthCheck.Request.GetMethod(),
					StatusCode: uint32(currentRoute.HealthCheck.Request.GetStatusCode()),
					Headers:    currentRoute.HealthCheck.Request.GetHeaders(),
					Body:       currentRoute.HealthCheck.Request.GetBody(),
				}
			}
		}

		proxyHandler, err := handler.NewFactory(proxyController)
		if err != nil {
			return fmt.Errorf("failed to create proxy handler: %w", err)
		}
		if currentRoute.Path == nil {
			currentRoute.Path = ptr.String("/")
		}

		listMiddlewareWithPriority := []*middleware.MiddlewareWithPriority{}
		if currentRoute.Plugins != nil {
			for _, p := range currentRoute.Plugins.Middleware {
				logger.Infof("Adding Plugin Middleware: %s", p.Name)
				pluginMiddleware, err := i.pluginManager.GetPluginMiddleware(p.Name)
				if err != nil {
					logger.Errorf("failed to get plugin middleware %s: %w", p.Name, err)
					continue
					//return fmt.Errorf("failed to get plugin %s: %w", p.Name, err)
				}
				listMiddlewareWithPriority = append(listMiddlewareWithPriority, &middleware.MiddlewareWithPriority{
					Name:       p.Name,
					Middleware: pluginMiddleware,
					Priority:   int(p.Priority),
				})
			}
		}
		//setupMiddleware := middleware.MiddlewareComposerWithPriority(listMiddlewareWithPriority)

		route := route.NewRoute(
			currentRoute.GetUuid(),
			currentRoute.Host,
			currentRoute.GetPath(),
			reverseproxy.NewReverseProxyHandler(proxyHandler),
			proxyController,
			route.WithHealthCheckConfig(healthCheckSettings),
			route.WithListMiddleware(listMiddlewareWithPriority),
		)

		if listUpstream.Data != nil {
			for idx, upstream := range listUpstream.Data.Items {
				if idx == 0 && currentRoute.HealthCheck != nil && currentRoute.HealthCheck.GetEnabled() {
					healthCheckRoute, err = healthcheck.NewHealthCheck(
						proxyController.ClientManager,
						healthCheckSettings,
					)
					i.router.AddHealthCheck(currentRoute.GetUuid(), healthCheckRoute)
					route.HealthCheckConfig = healthCheckSettings
					if proxyController.CircuitBreaker != nil && currentRoute.HealthCheck.GetCombinedWithCircuitBreaker() {
						healthCheckRoute.SetCircuitBreaker(proxyController.CircuitBreaker)
					}
					if err != nil {
						return fmt.Errorf("failed to create health check: %w", err)
					}
				}

				if upstream.Uuid == nil || upstream.Host == nil || upstream.Port == nil || upstream.Scheme == nil {
					continue
				}
				if upstream.Path == nil {
					upstream.Path = ptr.String("/")
				}

				logger.Infof("Adding Upstream: %s", *upstream.Uuid)
				upstreamModel := &upstreamPb.UpstreamPublishUpsertModel{
					Uuid:      upstream.GetUuid(),
					RouteHost: currentRoute.Host,
					Scheme:    upstream.GetScheme(),
					Host:      upstream.GetHost(),
					Path:      upstream.GetPath(),
					Port:      upstream.GetPort(),
					Weight:    float64(upstream.Weight),
				}
				proxyController.AddUpstream(upstreamModel)
			}
		}

		if err := i.router.Upsert(route); err != nil {
			return fmt.Errorf("failed to add route %s: %w", currentRoute.GetUuid(), err)
		}
	}

	return nil
}
