package authService

import "go-auth/internal/db"

type AuthService struct {
	db db.DB
}

func (as *AuthService) HasPermission(route string)
