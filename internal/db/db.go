package db

import "context"

type DB interface {
	GetPermissionsFromRoute(route string) ([]string, error)
}
type RoleDB interface {
	GetRole(ctx context.Context, roleId string)
}

type Role struct {
	Id          string
	Name        string
	Permissions []string
}
