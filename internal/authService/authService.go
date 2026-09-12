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

func (as *AuthService) HasPermission(route string, userPerms []string) (bool, error) {
	routePerms, err := as.db.GetPermissionsFromRoute(route)
	if err != nil {
		return false, err
	}
	for _, routePerm := range routePerms {
		if !slices.Contains(userPerms, routePerm) {
			log.Printf("Usuário não tem a permissão %v", routePerm)
			return false, fmt.Errorf("Falha ao validar a permissão %v\n", routePerm)
		}
	}
	return true, nil
}
