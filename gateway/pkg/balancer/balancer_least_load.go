package balancer

import (
	"errors"
	"sort"
	"sync"
)

type LeastLoad struct {
	mu      sync.Mutex
	servers []*Server
}

func NewLeastLoad() *LeastLoad {
	return &LeastLoad{
		servers: make([]*Server, 0),
	}
}

func (l *LeastLoad) AddServer(server *Server, opts ...ServerOption) {
	l.mu.Lock()
	defer l.mu.Unlock()
	for _, opt := range opts {
		opt(server)
	}
	l.servers = append(l.servers, server)
	l.resetWithoutLock()
}

func (l *LeastLoad) RemoveServer(uuid string) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	for i, server := range l.servers {
		if server.Uuid == uuid {
			l.servers = append(l.servers[:i], l.servers[i+1:]...)
			l.resetWithoutLock()
			return nil
		}
	}
	return nil
}

func (l *LeastLoad) GetStrategy() LoadBalancerStrategy {
	return BalancerLeastLoad
}

func (l *LeastLoad) RemoveAll() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.servers = []*Server{}
	l.resetWithoutLock()
}

func (l *LeastLoad) Reset() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.resetWithoutLock()
}

func (l *LeastLoad) resetWithoutLock() {
	for _, server := range l.servers {
		server.Weight = 0
	}
}

func (l *LeastLoad) Balance() ([]string, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if len(l.servers) == 0 {
		return nil, errors.New("no servers available")
	}

	sort.Slice(l.servers, func(i, j int) bool {
		return l.servers[i].Weight < l.servers[j].Weight
	})
	servers := []string{}
	for i, server := range l.servers {
		if i == 0 {
			server.Weight++
		}
		servers = append(servers, server.Uuid)
	}
	return servers, nil
}

func (l *LeastLoad) UpdateWeight(uuid string, weight int) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	for _, server := range l.servers {
		if server.Uuid == uuid {
			server.Weight = weight
			return nil
		}
	}
	return errors.New("server not found")
}
