package healthcheck

import (
	"time"

	"github.com/starton-io/tyrscale/gateway/pkg/circuitbreaker"
)

type HealthCheckType string

const (
	EthBlockNumberType HealthCheckType = "eth_block_number"
	EthSyncingType     HealthCheckType = "eth_syncing"
	CustomType         HealthCheckType = "custom"
)

type HealthCheckConfig struct {
	Enabled                    bool            `json:"enabled" validate:"omitempty"`
	Type                       HealthCheckType `json:"type" validate:"required_if=Enabled true"`
	CombinedWithCircuitBreaker bool            `json:"combined_with_circuit_breaker" validate:"omitempty"`
	Interval                   uint32          `json:"interval" validate:"required_if=Enabled true"`
	Timeout                    uint32          `json:"timeout" validate:"required_if=Enabled true"`
	Request                    *Request        `json:"request,omitempty" validate:"required_if=Type custom"`
}

type Request struct {
	StatusCode uint32            `json:"status_code"`
	Body       string            `json:"body"`
	Method     string            `json:"method"`
	Headers    map[string]string `json:"headers,omitempty"`
}

type HealthCheckInterface interface {
	CheckHealth() error
	GetInterval() time.Duration
	SetCircuitBreaker(circuitBreaker circuitbreaker.ProxyCircuitBreaker)
}
