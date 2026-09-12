package authService

import (
	"fmt"
	"go-auth/internal/db"
	"log"
	"slices"
)

type AuthService struct {
	db db.DB
}

func NewAuthService(db db.DB) AuthService {
	return AuthService{
		db: db,
	}
}

func (as *AuthService) HasPermission(route string, userPerms []string) (bool, error) {
	routePerms, err := as.db.GetPermissionsFromRoute(route)
	if err != nil {
		return false, err
	}
	for _, routePerm := range routePerms {
		if !slices.Contains(userPerms, routePerm) {
			log.Printf("User don't have permission %v", routePerm)
			return false, fmt.Errorf("Failed to validate permission %v\n", routePerm)
		}
	}
	return true, nil
}
