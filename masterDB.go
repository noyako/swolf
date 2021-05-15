package swolf

import (
	"database/sql"
	"fmt"

	"github.com/noyako/swolf/builder"
)

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

// Get returns database name for tenant.
func (m *masterDB) Get(id string) (string, error) {
	t, err := m.db.Begin()
	if err != nil {
		return "", err
	}
	defer t.Commit()

	key := m.mapper.IDToField(id)
	builder := builder.NewBuilder().
		Select(m.tenantColumn).
		From(m.tableName).
		Where(builder.Eq(m.keyColumn, "$1")).
		Build()
	rows, err := t.Query(builder, key)

	if err != nil {
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		var databaseName string
		err := rows.Scan(&databaseName)
		if err != nil {
			return "", err
		}

		return m.mapper.FieldToDatabase(databaseName), nil
	}

	err = rows.Err()
	if err != nil {
		return "", err
	}

	return "", fmt.Errorf(errorTenantNotFound, id, key)
}

// Create creates database name for tenant.
func (m *masterDB) Create(id string) (string, error) {
	err := m.db.Ping()
	if err != nil {
		panic(err)
	}
	t, err := m.db.Begin()
	if err != nil {
		return "", err
	}
	defer t.Commit()

	req := builder.NewBuilder().
		Insert(m.tableName, m.keyColumn, m.tenantColumn).
		Values("$1", "$2").
		Build()

	stmt, err := t.Prepare(req)
	if err != nil {
		return "", err
	}

	key := m.mapper.IDToField(id)
	value := m.mapper.KeyToValue(key)
	res, err := stmt.Exec(key, value)
	if err != nil {
		return "", err
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		return "", err
	}
	if rowCnt == 0 {
		return "", fmt.Errorf(errorCannotInsert, id, key, value)
	}

	return m.mapper.FieldToDatabase(value), nil
}

// Delete deletes database name for tenant.
func (m *masterDB) Delete(id string) (string, error) {
	t, err := m.db.Begin()
	if err != nil {
		return "", err
	}
	defer t.Commit()

	req := builder.NewBuilder().
		Delete().
		From(m.tableName).
		Where(builder.Eq(m.keyColumn, "$1")).
		Build()

	stmt, err := t.Prepare(req)
	if err != nil {
		return "", err
	}

	key := m.mapper.IDToField(id)
	value := m.mapper.FieldToDatabase(key)
	res, err := stmt.Exec(key)
	if err != nil {
		return "", err
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		return "", err
	}
	if rowCnt == 0 {
		return "", fmt.Errorf(errorCannotDelete, id, key)
	}

	return value, nil
}
