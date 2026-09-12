package inmemory

import (
	"errors"
	"log"
	"sync"
)

type Memory struct {
	mu *sync.Mutex
	db map[string][]string
}

func NewMemoryDB(permMap map[string][]string, mu *sync.Mutex) *Memory {
	return &Memory{
		mu: mu,
		db: permMap,
	}
}

func (m *Memory) GetPermissionsFromRoute(route string) ([]string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	v, ok := m.db[route]
	if !ok {
		log.Printf("Permission not found - Route: %v", route)
		return []string{}, errors.New("Permission not found!")
	}
	return v, nil
}
