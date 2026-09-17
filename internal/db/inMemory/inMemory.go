package inmemory

import (
	"context"
	"errors"
	"go-auth/internal/db"
	"log"
	"sync"
)

var _ db.RoleRepo = &Memory{}

type Memory struct {
	mu           *sync.Mutex
	permPerRoute map[string][]string
	permPerRole  map[string]db.Role
}

func NewMemoryDB(routeMap map[string][]string, roleMap map[string]db.Role, mu *sync.Mutex) *Memory {
	return &Memory{
		mu:           mu,
		permPerRoute: routeMap,
		permPerRole:  roleMap,
	}
}

func (m *Memory) GetPermissionsFromRoute(route string) ([]string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	v, ok := m.permPerRoute[route]
	if !ok {
		log.Printf("Permission not found - Route: %v", route)
		return []string{}, errors.New("Permission not found!")
	}
	return v, nil
}

func (m *Memory) GetRole(ctx context.Context, roleId string) (db.Role, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	v, ok := m.permPerRole[roleId]
	if !ok {
		log.Printf("Role Not Found! - Role: %v", roleId)
		return db.Role{}, errors.New("Role Not Found!")
	}
	return v, nil
}
