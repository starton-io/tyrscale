package balancer

import (
	"errors"
	"sync"
)

type Server struct {
	Uuid   string
	Weight int
}

type roundRobinWeight struct {
	sync.Mutex

	servers []*Server

	// index indicates the server selected last time , and i is initialized to -1
	index int
	// cw is the current weight
	cw int
	// maxWeight is the maximum weight of all servers
	maxWeight int
	// greatest common divisor of all server weights
	gcd int
	// current server list count
	count int
}

func NewWeightRoundRobin() IBalancer {
	rb := &roundRobinWeight{
		servers: []*Server{},
		count:   0,
		index:   -1,
		cw:      0,
	}
	return rb
}

func (r *roundRobinWeight) Balance() ([]string, error) {
	r.Lock()
	defer r.Unlock()

	if r.count == 0 {
		return []string{}, nil
	}

	if r.count == 1 {
		return []string{r.servers[0].Uuid}, nil
	}
	for {
		r.index = (r.index + 1) % r.count
		if r.index == 0 {
			r.cw = r.cw - r.gcd
			if r.cw <= 0 {
				r.cw = r.maxWeight
				if r.cw == 0 {
					return []string{}, nil
				}
			}
		}
		if r.servers[r.index].Weight >= r.cw {
			return []string{r.servers[r.index].Uuid}, nil
		}
	}
}

type ServerOption func(*Server)

func (r *roundRobinWeight) AddServer(server *Server, opts ...ServerOption) {
	r.Lock()
	defer r.Unlock()

	for _, opt := range opts {
		opt(server)
	}

	if server.Weight > 0 {
		if r.gcd == 0 {
			r.gcd = server.Weight
			r.maxWeight = server.Weight
			r.index = -1
			r.cw = 0
		} else {
			r.gcd = gcd(r.gcd, server.Weight)
			if r.maxWeight < server.Weight {
				r.maxWeight = server.Weight
			}
		}
	}
	r.servers = append(r.servers, server)
	r.count++
}

func (r *roundRobinWeight) GetStrategy() LoadBalancerStrategy {
	return BalancerWeightRoundRobin
}

func (r *roundRobinWeight) UpdateWeight(uuid string, newWeight int) error {
	r.Lock()
	defer r.Unlock()

	if newWeight <= 0 {
		return errors.New("weight must be positive")
	}
	updated := false
	founded := false
	for _, server := range r.servers {
		if server.Uuid == uuid {
			founded = true
			if server.Weight != newWeight {
				server.Weight = newWeight
				updated = true
			}
			break
		}
	}
	if !founded {
		return errors.New("UUID not found")
	}

	if updated {
		// Recalculate gcd and maxWeight only if a weight was actually updated
		r.gcd = r.servers[0].Weight
		r.maxWeight = r.servers[0].Weight
		for _, server := range r.servers {
			r.gcd = gcd(r.gcd, server.Weight)
			r.maxWeight = max(r.maxWeight, server.Weight)
		}
		r.index = -1
		r.cw = 0
	}

	return nil
}

func (r *roundRobinWeight) RemoveServer(uuid string) error {
	r.Lock()
	defer r.Unlock()
	for i, s := range r.servers {
		if uuid == s.Uuid {
			r.servers = append(r.servers[:i], r.servers[i+1:]...)
			r.count--
			if len(r.servers) > 0 {
				r.gcd = r.servers[0].Weight
				r.maxWeight = r.servers[0].Weight
				for _, server := range r.servers {
					r.maxWeight = max(r.maxWeight, server.Weight)
					r.gcd = gcd(r.gcd, server.Weight)
				}
			} else {
				r.gcd = 0
			}
		}
	}
	return ErrNoUpstreams
}

func (r *roundRobinWeight) RemoveAll() {
	r.Lock()
	defer r.Unlock()

	r.servers = r.servers[:0]
	r.count = 0
	r.cw = 0
	r.index = -1
	r.gcd = 0
	r.maxWeight = 0
}

func (r *roundRobinWeight) Reset() {
	r.Lock()
	defer r.Unlock()

	r.index = -1
	r.cw = 0
}

//func gcd(a, b int) int {
//	for b != 0 {
//		a, b = b, a%b
//	}
//	return a
//}

func gcd(x, y int) int {
	var t int
	for {
		t = (x % y)
		if t > 0 {
			x = y
			y = t
		} else {
			return y
		}
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
