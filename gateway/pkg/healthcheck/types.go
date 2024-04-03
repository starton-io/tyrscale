package healthcheck

import (
	"time"

	"github.com/starton-io/tyrscale/gateway/pkg/circuitbreaker"
)

type HealthCheckType string

const (
	EthBlockNumberType HealthCheckType = "eth_block_number"
)

type HealthCheckConfig struct {
	Enabled                    bool            `json:"enabled" validate:"omitempty"`
	Type                       HealthCheckType `json:"type" validate:"required_if=Enabled true"`
	CombinedWithCircuitBreaker bool            `json:"combined_with_circuit_breaker" validate:"omitempty"`
	Interval                   uint32          `json:"interval" validate:"required_if=Enabled true"`
	Timeout                    uint32          `json:"timeout" validate:"required_if=Enabled true"`
}

type HealthCheckInterface interface {
	CheckHealth() error
	GetInterval() time.Duration
	SetCircuitBreaker(circuitBreaker circuitbreaker.ProxyCircuitBreaker)
}
