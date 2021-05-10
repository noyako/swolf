package swolf

import "database/sql"

type masterController interface {
	Get(string) (string, error)
	Create(string) (string, error)
	Delete(string) (string, error)
}

type tenantController interface {
	Get(string) (*sql.DB, error)
	Create(string) (*sql.DB, error)
	Delete(string) error
}
