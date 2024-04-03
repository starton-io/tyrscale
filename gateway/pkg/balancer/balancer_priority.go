package balancer

import (
	"sort"
	"sync"
)

type Priority struct {
	servers []*Server
	sync.Mutex
}

func NewPriorityBalancer() IBalancer {
	return &Priority{}
}

func (f *Priority) AddServer(server *Server, opts ...ServerOption) {
	f.Lock()
	defer f.Unlock()

	for _, opt := range opts {
		opt(server)
	}
	f.servers = append(f.servers, server)
}

func (f *Priority) RemoveServer(uuid string) error {
	f.Lock()
	defer f.Unlock()

	for i, s := range f.servers {
		if s.Uuid == uuid {
			f.servers = append(f.servers[:i], f.servers[i+1:]...)
			break
		}
	}
	return nil
}

func (f *Priority) GetStrategy() LoadBalancerStrategy {
	return BalancerPriority
}

func (f *Priority) UpdateWeight(uuid string, weight int) error {
	f.Lock()
	defer f.Unlock()

	for _, server := range f.servers {
		if server.Uuid == uuid {
			server.Weight = weight
		}
	}
	return nil
}

func (f *Priority) RemoveAll() {
	f.Lock()
	defer f.Unlock()

	f.servers = f.servers[:0]
}

func (f *Priority) Reset() {
	f.Lock()
	defer f.Unlock()

	f.servers = f.servers[:0]
}

func (f *Priority) Balance() ([]string, error) {
	f.Lock()
	defer f.Unlock()

	// Sort the copied slice of servers by weight
	sort.Slice(f.servers, func(i, j int) bool {
		return f.servers[i].Weight > f.servers[j].Weight
	})

	// Extract UUIDs from the sorted slice of servers
	uuids := make([]string, len(f.servers))
	for i, server := range f.servers {
		uuids[i] = server.Uuid
	}

	return uuids, nil
}
