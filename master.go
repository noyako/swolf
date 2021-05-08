package swolf

import "database/sql"

type masterDB struct {
	db        *sql.DB
	tableName string

	keyColumn    string
	tenantColumn string

	mapper FMapper
}

func newMasterDB(db *sql.DB, table, key, tenant string, mapper FMapper) *masterDB {
	return &masterDB{
		db:           db,
		tableName:    table,
		keyColumn:    key,
		tenantColumn: tenant,
		mapper:       mapper,
	}
}
