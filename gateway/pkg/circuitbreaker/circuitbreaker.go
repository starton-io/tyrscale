package circuitbreaker

import (
	"sync"
	"time"

	"github.com/sony/gobreaker"
)

type Settings struct {
	Enabled                bool   `json:"enabled,omitempty" validate:"omitempty"`
	Name                   string `json:"name,omitempty"`
	MaxRequests            uint32 `json:"max_requests,omitempty" validate:"required_if=Enabled True"`
	MaxConsecutiveFailures uint32 `json:"max_consecutive_failures,omitempty" validate:"required_if=Enabled True"`
	Interval               uint32 `json:"interval,omitempty" validate:"required_if=Enabled True"`
	Timeout                uint32 `json:"timeout,omitempty" validate:"required_if=Enabled True"`
}

type DefaultProxyCircuitBreaker struct {
	mu             sync.Mutex
	settings       *gobreaker.Settings
	CircuitBreaker map[string]*gobreaker.CircuitBreaker
}

type ProxyCircuitBreaker interface {
	Get(key string) *gobreaker.CircuitBreaker
	Add(key string)
	Remove(key string)
	Clean()
}

func NewCircuitBreaker(settings Settings) ProxyCircuitBreaker {
	// set default values if not set
	if settings.MaxRequests == 0 {
		settings.MaxRequests = 3
	}
	if settings.MaxConsecutiveFailures == 0 {
		settings.MaxConsecutiveFailures = 3
	}
	if settings.Interval == 0 {
		settings.Interval = 120000
	}
	if settings.Timeout == 0 {
		settings.Timeout = 60000
	}

	gobreakerSettings := gobreaker.Settings{
		Name:        settings.Name,
		MaxRequests: settings.MaxRequests,
		Interval:    time.Duration(settings.Interval) * time.Millisecond,
		Timeout:     time.Duration(settings.Timeout) * time.Millisecond,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures >= settings.MaxConsecutiveFailures
		},
	}
	return &DefaultProxyCircuitBreaker{settings: &gobreakerSettings, CircuitBreaker: make(map[string]*gobreaker.CircuitBreaker)}
}

func (c *DefaultProxyCircuitBreaker) Add(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.CircuitBreaker[key] = gobreaker.NewCircuitBreaker(*c.settings)
}

func (c *DefaultProxyCircuitBreaker) Get(key string) *gobreaker.CircuitBreaker {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.CircuitBreaker[key]
}

func (c *DefaultProxyCircuitBreaker) Remove(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.CircuitBreaker, key)
}

func (c *DefaultProxyCircuitBreaker) Clean() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.CircuitBreaker = nil
}
