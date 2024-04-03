package dto

import (
	"github.com/starton-io/tyrscale/gateway/pkg/balancer"
	"github.com/starton-io/tyrscale/gateway/pkg/circuitbreaker"
	"github.com/starton-io/tyrscale/gateway/pkg/healthcheck"
)

type Route struct {
	Uuid                 string                        `json:"uuid" validate:"omitempty,uuid"`
	Host                 string                        `json:"host" validate:"required"`
	Path                 string                        `json:"path" validate:"omitempty"`
	LoadBalancerStrategy balancer.LoadBalancerStrategy `json:"load_balancer_strategy" validate:"required"`
	CircuitBreaker       *circuitbreaker.Settings      `json:"circuit_breaker,omitempty" validate:"omitempty"`
	HealthCheck          *HealthCheckConfig            `json:"health_check,omitempty" validate:"omitempty"`
}

// CircuitBreakerConfig holds the configuration for the circuit breaker.
//type CircuitBreakerConfig struct {
//	Settings *circuitbreaker.Settings `json:"settings,omitempty" validate:"required_if=Enabled true,dive"`
//}

type HealthCheckConfig struct {
	Enabled                    bool                        `json:"enabled" validate:"omitempty"`
	Type                       healthcheck.HealthCheckType `json:"type" validate:"required_if=Enabled true"`
	CombinedWithCircuitBreaker bool                        `json:"combined_with_circuit_breaker" validate:"omitempty"`
	Interval                   uint32                      `json:"interval" validate:"required_if=Enabled true"`
	Timeout                    uint32                      `json:"timeout" validate:"required_if=Enabled true"`
}

// TODO: Use this later when route handler update was implemented
type UpdateRouteReq struct {
	Uuid                 string                        `json:"uuid" validate:"required,uuid"`
	Host                 string                        `json:"host"`
	Path                 string                        `json:"path" validate:"omitempty"`
	LoadBalancerStrategy balancer.LoadBalancerStrategy `json:"load_balancer_strategy" validate:"omitempty"`
}

type CreateRouteRes struct {
	Uuid string `json:"uuid"`
}

type ListRouteRes struct {
	Routes []Route `json:"items"`
}

type ListRouteReq struct {
	Uuid                 string                        `query:"uuid" validate:"omitempty,uuid"`
	Host                 string                        `query:"host" validate:"omitempty,regexp=^[a-zA-Z0-9_-]+$"`
	Path                 string                        `query:"path" validate:"omitempty,regexp=^\\/[a-zA-Z0-9_-\\/]+$"`
	LoadBalancerStrategy balancer.LoadBalancerStrategy `query:"load_balancer_strategy" validate:"omitempty"`
}

type DeleteRouteReq struct {
	Uuid string `json:"uuid" validate:"required,uuid"`
}
