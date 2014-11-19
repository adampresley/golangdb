package golangdb

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

/*
This method establishes a connection to a MSSQL database
server. The connection details are stored and sent via
a DatabaseConnection structure.
*/
func ConnectMSSQL(connectionInfo *DatabaseConnection) (*sql.DB, error) {
	/*
	 * Create the connection
	 */
	log.Println("INFO - Connecting to MSSQL database...")

	db, err := sql.Open("mssql", connectionInfo.ToString())
	if err != nil {
		return nil, err
	}

	return db, nil
}
