package swolf

import "database/sql"

// MasterController controls tenant data in master database
type MasterController interface {
	Get(string) (string, error)
	Create(string) (string, error)
	Delete(string) (string, error)
}

// TenantController controls databases for tenants
type TenantController interface {
	Get(string) (*sql.DB, error)
	Create(string) (*sql.DB, error)
	Delete(string) error
}
