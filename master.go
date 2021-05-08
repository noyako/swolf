package swolf

import (
	"database/sql"
	"fmt"
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

func (m *masterDB) Get(id string) (string, error) {
	t, err := m.db.Begin()
	if err != nil {
		return "", err
	}
	defer t.Commit()

	key := m.mapper.IDToField(id)
	rows, err := t.Query("SELECT $1 FROM $2 WHERE $3=$4", m.tenantColumn, m.tableName, m.keyColumn, key)

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

	return "", fmt.Errorf("Tenant for id=%s (key=%s) not found", id, key)
}

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

	// stmt, err := t.Prepare("INSERT INTO $1($2, $3) VALUES($4, $5)")
	stmt, err := t.Prepare("INSERT INTO main_table(tenant_id, tenant_db) VALUES($1, $2)")
	if err != nil {
		return "", err
	}

	key := m.mapper.IDToField(id)
	value := m.mapper.KeyToValue(key)
	// res, err := stmt.Exec(m.tableName, m.keyColumn, m.tenantColumn, key, value)
	res, err := stmt.Exec(key, value)
	if err != nil {
		return "", err
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		return "", err
	}
	if rowCnt == 0 {
		return "", fmt.Errorf("Could not insert for id=%s (key=%s, value=%s)", id, key, value)
	}

	return m.mapper.FieldToDatabase(value), nil
}

func (m *masterDB) Delete(id string) error {
	t, err := m.db.Begin()
	if err != nil {
		return err
	}
	defer t.Commit()

	stmt, err := t.Prepare("DELETE FROM $1 WHERE $2=$3")
	if err != nil {
		return err
	}

	key := m.mapper.IDToField(id)
	res, err := stmt.Exec(m.tableName, key)
	if err != nil {
		return err
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowCnt == 0 {
		return fmt.Errorf("Could not delete for id=%s (key=%s)", id, key)
	}

	return nil
}
