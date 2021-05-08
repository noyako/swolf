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
