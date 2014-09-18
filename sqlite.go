package golangdb

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectSQLite(connectionInfo DatabaseConnection) (*sql.DB, error) {
	/*
	 * Create the connection
	 */
	log.Println("Connecting to SQLITE3 database", connectionInfo.ToString())

	db, err := sql.Open("sqlite3", connectionInfo.ToString())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func DropSQLite(connectionInfo DatabaseConnection) {
	os.Remove(connectionInfo.ToString())
}