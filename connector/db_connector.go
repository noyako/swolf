package connector

import "database/sql"

// Connect to database with diven name and credentials as connection variable
func Connect(driver, connection, database string) (*sql.DB, error) {
	var db *sql.DB
	var err error

	switch driver {
	case "postgres":
		db, err = sql.Open(driver, connection+" dbname="+database)
	case "mysql":
		db, err = sql.Open(driver, connection+database)
	}

	return db, err
}
