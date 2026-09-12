package db

type DB interface {
	GetPermissionsFromRoute(route string) ([]string, error)
}
