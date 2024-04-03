package balancer

import (
	"errors"
	"fmt"
)

var (
	ErrNoUpstreams = errors.New("no upstreams")
)

type LoadBalancerStrategy string

const (
	BalancerWeightRoundRobin LoadBalancerStrategy = "weight-round-robin"
	BalancerLeastLoad        LoadBalancerStrategy = "least-load"
	BalancerPriority         LoadBalancerStrategy = "failover-priority"
)

var validateBalancerValues = []LoadBalancerStrategy{
	BalancerWeightRoundRobin,
	BalancerLeastLoad,
	BalancerPriority,
}

func (b LoadBalancerStrategy) String() string {
	return string(b)
}

func (b LoadBalancerStrategy) Validate() error {
	for _, value := range validateBalancerValues {
		if b == value {
			return nil
		}
	}
	return fmt.Errorf("invalid balancer type: %s", b)
}

type IBalancer interface {
	Balance() ([]string, error)
	AddServer(server *Server, opts ...ServerOption)
	GetStrategy() LoadBalancerStrategy
	UpdateWeight(uuid string, weight int) error
	RemoveServer(uuid string) error
	RemoveAll()

	// Reset all current weights
	Reset()
}

func NewBalancer(b LoadBalancerStrategy) IBalancer {
	switch b {
	case BalancerWeightRoundRobin:
		return NewWeightRoundRobin()
	case BalancerPriority:
		return NewPriorityBalancer()
	case BalancerLeastLoad:
		return NewLeastLoad()
	default:
		return NewPriorityBalancer()
	}
}
