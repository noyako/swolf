package swolf

import (
	"database/sql"
)

type tenantDB struct {
	db         *sql.DB
	template   string
	connection string
}

func newTenantDB(db *sql.DB, connection, tmpl string) *tenantDB {

	return &tenantDB{
		db:       db,
		template: tmpl,
	}
}

func (t *tenantDB) Get(name string) (*sql.DB, error) {
	// TODO
	return nil, nil
}

func (t *tenantDB) Create(name string) (*sql.DB, error) {
	// TODO
	return nil, nil
}

func (t *tenantDB) Delete(name string) error {
	// TODO
	return nil
}
