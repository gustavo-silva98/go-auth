package authService

import (
	inmemory "go-auth/internal/db/inMemory"
	"strings"
	"sync"
	"testing"
)

func NewTestAuthService() *AuthService {
	perms := map[string][]string{}
	perms["route"] = []string{"1", "2", "3"}
	var mu sync.Mutex
	db := inmemory.NewMemoryDB(perms, &mu)
	as := NewAuthService(db)
	return &as
}
func TestHasPermission(t *testing.T) {
	t.Run("User dont have perms", func(t *testing.T) {
		userPerm := []string{"4"}
		as := NewTestAuthService()
		hasPerms, err := as.HasPermission("routeDoNotExist", userPerm)
		if err == nil {
			t.Error("Failed to generate error getting permissions")
		}
		if hasPerms {
			t.Error("Failed to generate false status on search")
		}
	})
	t.Run("User dont have all Perms", func(t *testing.T) {
		userPerm := []string{"1"}
		as := NewTestAuthService()
		hasPerms, err := as.HasPermission("route", userPerm)
		if err == nil {
			t.Error("Failed to generate error getting permissions")
		}
		if strings.TrimSpace(err.Error()) != "Failed to validate permission 2" {
			t.Errorf("Error message != from expected: %s", err.Error())
		}
		if hasPerms {
			t.Error("Status false on searching perms for route/user")
		}
	})
	t.Run("User have all Perms", func(t *testing.T) {
		userPerm := []string{"1", "2", "3", "4"}
		as := NewTestAuthService()
		hasPerms, err := as.HasPermission("route", userPerm)
		if err != nil {
			t.Errorf("Error getting search user/route perms: %v", err)
		}
		if !hasPerms {
			t.Error("Status false on searching perms for route/user")
		}
	})
}
