package golangdb

import (
	"database/sql"
	"log"

	_ "github.com/erikstmartin/go-testdb"
)

/*
This method establishes a fake database connection.
The connection details are stored and sent via
a DatabaseConnection structure.
*/
func ConnectTestDB(connectionInfo *DatabaseConnection) (*sql.DB, error) {
	/*
	 * Create the connection
	 */
	log.Println("INFO - Connecting to Go-TestDB database...")

	db, err := sql.Open("testdb", connectionInfo.ToString())
	if err != nil {
		return nil, err
	}

	return db, nil
}
