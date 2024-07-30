package plugin

import (
	"context"
	"sync"
)

type IPluginStorage interface {
	AddPlugin(ctx context.Context, pluginName string, pluginType string) error
	ListPlugins(ctx context.Context) (map[string][]string, error)
	ListPluginsByType(ctx context.Context, pluginType string) ([]string, error)
}

type InMemoryPluginStorage struct {
	mu      sync.Mutex
	storage map[string][]string
}

func NewInMemoryPluginStorage() *InMemoryPluginStorage {
	return &InMemoryPluginStorage{
		storage: make(map[string][]string),
	}
}

func (s *InMemoryPluginStorage) AddPlugin(ctx context.Context, pluginName string, pluginType string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.storage[pluginType] == nil {
		s.storage[pluginType] = []string{}
	}
	s.storage[pluginType] = append(s.storage[pluginType], pluginName)
	return nil
}

func (s *InMemoryPluginStorage) ListPlugins(ctx context.Context) (map[string][]string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.storage, nil
}
func (s *InMemoryPluginStorage) ListPluginsByType(ctx context.Context, pluginType string) ([]string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.storage[pluginType], nil
}
