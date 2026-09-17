package inmemory

import (
	"go-auth/internal/db"
	"slices"
	"sync"
	"testing"
)

func testNewMemory() *Memory {
	perms := map[string][]string{}
	roleMap := map[string]db.Role{}
	perms["route"] = []string{"1", "2", "3"}
	var mu sync.Mutex
	return NewMemoryDB(perms, roleMap, &mu)
}

func TestGetPermissionsFromRoute(t *testing.T) {
	t.Run("Permissão não encontrada", func(t *testing.T) {
		m := testNewMemory()
		perms, err := m.GetPermissionsFromRoute("routeNotFound")
		if err == nil {
			t.Error("Error simulating error from GetPermissionsFromRoute")
		}
		if len(perms) > 0 {
			t.Errorf("Perms len is > 0: %v", len(perms))
		}
	})
	t.Run("Permissão encontrada", func(t *testing.T) {
		permsAnswer := []string{"1", "2", "3"}
		m := testNewMemory()
		perms, err := m.GetPermissionsFromRoute("route")
		if err != nil {
			t.Errorf("Error on GetPermissionsFromRoute: %v", err)
		}
		if !slices.Equal(perms, permsAnswer) {
			t.Errorf("Slices are not equal")
		}
		if len(perms) != 3 {
			t.Errorf("Len of perms is different from 3: Len %v", len(perms))
		}
	})
}
