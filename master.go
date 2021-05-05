package swolf

import "database/sql"

type master struct {
	db        *sql.DB
	tableName string
}
