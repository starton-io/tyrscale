package healthcheck

import (
	"github.com/madflojo/tasks"
	"github.com/starton-io/tyrscale/gateway/pkg/circuitbreaker"
)

type HealthCheckManager struct {
	Scheduler      *tasks.Scheduler
	CircuitBreaker circuitbreaker.ProxyCircuitBreaker
}

type HealthCheckManagerInterface interface {
	AddHealthCheck(id string, healthCheck HealthCheckInterface)
	SetCircuitBreaker(circuitBreaker circuitbreaker.ProxyCircuitBreaker)
	GetInterval(id string) uint32
	RemoveHealthCheck(id string)
}

func NewHealthCheckManager() HealthCheckManagerInterface {
	return &HealthCheckManager{Scheduler: tasks.New()}
}

func (m *HealthCheckManager) SetCircuitBreaker(circuitBreaker circuitbreaker.ProxyCircuitBreaker) {
	m.CircuitBreaker = circuitBreaker
}

func (m *HealthCheckManager) AddHealthCheck(id string, healthCheck HealthCheckInterface) {
	if m.CircuitBreaker != nil {
		healthCheck.SetCircuitBreaker(m.CircuitBreaker)
	}
	_ = m.Scheduler.AddWithID(id, &tasks.Task{
		Interval: healthCheck.GetInterval(),
		TaskFunc: healthCheck.CheckHealth,
	})
}

func (m *HealthCheckManager) RemoveHealthCheck(id string) {
	m.Scheduler.Del(id)
}

func (m *HealthCheckManager) GetInterval(id string) uint32 {
	if len(m.Scheduler.Tasks()) > 0 {
		tasks := m.Scheduler.Tasks()
		if _, ok := tasks[id]; ok {
			return uint32(tasks[id].Interval)
		}
	}
	return 0
}

func (m *HealthCheckManager) CleanTasks() {
	m.Scheduler.Stop()
}
