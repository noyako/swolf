package connector

import (
	"database/sql"
	"fmt"
)

// Connect to database with diven name and credentials as connection variable
func Connect(driver, connection, database string) (*sql.DB, error) {
	return sql.Open(driver, fmt.Sprintf(connection, database))
}
