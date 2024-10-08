package balancer

import (
	"testing"
)

func TestPriority_AddServer(t *testing.T) {
	balancer := NewPriorityBalancer().(*Priority)
	server := &Server{Uuid: "server1", Weight: 10}

	balancer.AddServer(server)

	if len(balancer.servers) != 1 {
		t.Errorf("expected 1 server, got %d", len(balancer.servers))
	}
	if balancer.servers[0].Uuid != "server1" {
		t.Errorf("expected server UUID 'server1', got '%s'", balancer.servers[0].Uuid)
	}
}

func TestPriority_RemoveServer(t *testing.T) {
	balancer := NewPriorityBalancer().(*Priority)
	server := &Server{Uuid: "server1", Weight: 10}
	balancer.AddServer(server)

	err := balancer.RemoveServer("server1")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(balancer.servers) != 0 {
		t.Errorf("expected 0 servers, got %d", len(balancer.servers))
	}
}

func TestPriority_UpdateWeight(t *testing.T) {
	balancer := NewPriorityBalancer().(*Priority)
	server := &Server{Uuid: "server1", Weight: 10}
	balancer.AddServer(server)

	err := balancer.UpdateWeight("server1", 20)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if balancer.servers[0].Weight != 20 {
		t.Errorf("expected weight 20, got %d", balancer.servers[0].Weight)
	}
}

func TestPriority_RemoveAll(t *testing.T) {
	balancer := NewPriorityBalancer().(*Priority)
	server1 := &Server{Uuid: "server1", Weight: 10}
	server2 := &Server{Uuid: "server2", Weight: 20}
	balancer.AddServer(server1)
	balancer.AddServer(server2)

	balancer.RemoveAll()
	if len(balancer.servers) != 0 {
		t.Errorf("expected 0 servers, got %d", len(balancer.servers))
	}
}

func TestPriority_Reset(t *testing.T) {
	balancer := NewPriorityBalancer().(*Priority)
	server1 := &Server{Uuid: "server1", Weight: 10}
	server2 := &Server{Uuid: "server2", Weight: 20}
	balancer.AddServer(server1)
	balancer.AddServer(server2)

	balancer.Reset()
	if len(balancer.servers) != 0 {
		t.Errorf("expected 0 servers, got %d", len(balancer.servers))
	}
}

func TestPriority_Balance(t *testing.T) {
	balancer := NewPriorityBalancer().(*Priority)
	server1 := &Server{Uuid: "server1", Weight: 10}
	server2 := &Server{Uuid: "server2", Weight: 20}
	balancer.AddServer(server1)
	balancer.AddServer(server2)

	uuids, err := balancer.Balance()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(uuids) != 2 {
		t.Errorf("expected 2 UUIDs, got %d", len(uuids))
	}
	if uuids[0] != "server2" || uuids[1] != "server1" {
		t.Errorf("expected order ['server2', 'server1'], got %v", uuids)
	}
}

func TestLeastLoad_AddServer(t *testing.T) {
	ll := NewLeastLoad()
	server := &Server{Uuid: "server1"}
	ll.AddServer(server)

	if len(ll.servers) != 1 {
		t.Errorf("expected 1 server, got %d", len(ll.servers))
	}
}

func TestLeastLoad_RemoveServer(t *testing.T) {
	ll := NewLeastLoad()
	server := &Server{Uuid: "server1"}
	ll.AddServer(server)

	err := ll.RemoveServer("server1")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(ll.servers) != 0 {
		t.Errorf("expected 0 servers, got %d", len(ll.servers))
	}
}

func TestLeastLoad_Balance(t *testing.T) {
	ll := NewLeastLoad()
	server1 := &Server{Uuid: "server1"}
	server2 := &Server{Uuid: "server2"}
	ll.AddServer(server1)
	ll.AddServer(server2)

	servers, err := ll.Balance()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(servers) != 2 {
		t.Errorf("expected 2 servers, got %d", len(servers))
	}

	if servers[0] != "server1" {
		t.Errorf("expected server1 to be first, got %s", servers[0])
	}

	servers, err = ll.Balance()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(servers) != 2 {
		t.Errorf("expected 2 servers, got %d", len(servers))
	}
	if servers[0] != "server2" {
		t.Errorf("expected server2 to be first, got %s", servers[0])
	}
}

func TestLeastLoad_UpdateWeight(t *testing.T) {
	ll := NewLeastLoad()
	server := &Server{Uuid: "server1"}
	ll.AddServer(server)

	err := ll.UpdateWeight("server1", 10)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if server.Weight != 10 {
		t.Errorf("expected weight 10, got %d", server.Weight)
	}
}

func TestLeastLoad_RemoveAll(t *testing.T) {
	ll := NewLeastLoad()
	server1 := &Server{Uuid: "server1"}
	server2 := &Server{Uuid: "server2"}
	ll.AddServer(server1)
	ll.AddServer(server2)

	ll.RemoveAll()

	if len(ll.servers) != 0 {
		t.Errorf("expected 0 servers, got %d", len(ll.servers))
	}
}

func TestLeastLoad_Reset(t *testing.T) {
	ll := NewLeastLoad()
	server1 := &Server{Uuid: "server1", Weight: 5}
	server2 := &Server{Uuid: "server2", Weight: 10}
	ll.AddServer(server1)
	ll.AddServer(server2)

	ll.Reset()

	for _, server := range ll.servers {
		if server.Weight != 0 {
			t.Errorf("expected weight 0, got %d", server.Weight)
		}
	}
}

func TestNewWeightRoundRobin(t *testing.T) {
	balancer := NewWeightRoundRobin()
	if balancer == nil {
		t.Fatal("Expected non-nil balancer")
	}
}

func TestAddServer(t *testing.T) {
	balancer := NewWeightRoundRobin().(*roundRobinWeight)
	server := &Server{Uuid: "server1", Weight: 5}
	balancer.AddServer(server)

	if len(balancer.servers) != 1 {
		t.Fatalf("Expected 1 server, got %d", len(balancer.servers))
	}
	if balancer.servers[0].Uuid != "server1" {
		t.Fatalf("Expected server UUID 'server1', got '%s'", balancer.servers[0].Uuid)
	}
}

func TestBalance(t *testing.T) {
	balancer := NewWeightRoundRobin().(*roundRobinWeight)
	server1 := &Server{Uuid: "server1", Weight: 5}
	server2 := &Server{Uuid: "server2", Weight: 1}
	balancer.AddServer(server1)
	balancer.AddServer(server2)

	// Perform multiple balance operations to test round-robin behavior
	expectedOrder := []string{"server1", "server1", "server1", "server1", "server1", "server2", "server1", "server1", "server1", "server1", "server1", "server2"}
	for i, expected := range expectedOrder {
		uuid, err := balancer.Balance()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if len(uuid) != 1 {
			t.Fatalf("Expected 1 UUID, got %d", len(uuid))
		}
		if uuid[0] != expected {
			t.Fatalf("At iteration %d, expected server UUID '%s', got '%s'", i, expected, uuid[0])
		}
	}
}

func TestUpdateWeight(t *testing.T) {
	balancer := NewWeightRoundRobin().(*roundRobinWeight)
	server := &Server{Uuid: "server1", Weight: 5}
	balancer.AddServer(server)

	err := balancer.UpdateWeight("server1", 10)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if balancer.servers[0].Weight != 10 {
		t.Fatalf("Expected weight 10, got %d", balancer.servers[0].Weight)
	}
}

func TestRemoveServer(t *testing.T) {
	balancer := NewWeightRoundRobin().(*roundRobinWeight)
	server := &Server{Uuid: "server1", Weight: 5}
	balancer.AddServer(server)

	err := balancer.RemoveServer("server1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(balancer.servers) != 0 {
		t.Fatalf("Expected 0 servers, got %d", len(balancer.servers))
	}
}

func TestRemoveAll(t *testing.T) {
	balancer := NewWeightRoundRobin().(*roundRobinWeight)
	server1 := &Server{Uuid: "server1", Weight: 5}
	server2 := &Server{Uuid: "server2", Weight: 1}
	balancer.AddServer(server1)
	balancer.AddServer(server2)

	balancer.RemoveAll()
	if len(balancer.servers) != 0 {
		t.Fatalf("Expected 0 servers, got %d", len(balancer.servers))
	}
}

func TestReset(t *testing.T) {
	balancer := NewWeightRoundRobin().(*roundRobinWeight)
	server1 := &Server{Uuid: "server1", Weight: 5}
	server2 := &Server{Uuid: "server2", Weight: 1}
	balancer.AddServer(server1)
	balancer.AddServer(server2)

	balancer.Reset()
	if balancer.index != -1 {
		t.Fatalf("Expected index -1, got %d", balancer.index)
	}
	if balancer.cw != 0 {
		t.Fatalf("Expected current weight 0, got %d", balancer.cw)
	}
}
