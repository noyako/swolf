package swolf

import (
	"database/sql"

	"github.com/aoyako/swolf/builder"
	"github.com/aoyako/swolf/connector"
)

type tenantDB struct {
	db         *sql.DB
	template   string
	connection string
	driver     string
}

func newTenantDB(db *sql.DB, connection, drv, tmpl string) *tenantDB {

	return &tenantDB{
		db:       db,
		template: tmpl,
		driver:   drv,
	}
}

func (t *tenantDB) Get(name string) (*sql.DB, error) {
	conn, err := connector.Connect(t.driver, t.connection, name)
	return conn, err
}

func (t *tenantDB) Create(name string) (*sql.DB, error) {
	req := builder.NewBuilder().Create(name).Build()
	_, err := t.db.Exec(req)
	if err != nil {
		return nil, err
	}

	return connector.Connect(t.driver, t.connection, name)
}

func (t *tenantDB) Delete(name string) error {
	req := builder.NewBuilder().Drop(name).Build()
	_, err := t.db.Exec(req)
	return err
}
