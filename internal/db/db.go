package db

import (
	"context"
)

type DB interface {
	GetPermissionsFromRoute(route string) ([]string, error)
}
type RoleRepo interface {
	GetRole(ctx context.Context, roleId string) (Role, error)
}

type Role struct {
	Id          string
	Name        string
	Permissions []string
}
