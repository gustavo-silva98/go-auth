package inmemory

import (
	"errors"
	"log"
	"sync"
)

type Memory struct {
	mu sync.Mutex
	db map[string][]string
}

func (m *Memory) GetPermissions(route string) ([]string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	v, ok := db[route]
	if !ok {
		log.Printf("Permission not found - Route: %v", route)
		return []string{}, errors.New("Permission not found!")
	}
	return v, nil
}
