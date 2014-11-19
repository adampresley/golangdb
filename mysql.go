package golangdb

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

/*
This method establishes a connection to a MySQL database
server. The connection details are stored and sent via
a DatabaseConnection structure. Please note that this method
respects the environment variable max_connections, and
will default to "151" if it is not present.
*/
func ConnectMySQL(connectionInfo *DatabaseConnection) (*sql.DB, error) {
	/*
	 * Create the connection
	 */
	log.Println("INFO - Connecting to MySQL database...")

	db, err := sql.Open("mysql", connectionInfo.ToString())
	if err != nil {
		return nil, err
	}

	temp := os.Getenv("max_connections")
	if temp == "" {
		temp = "151"
	}

	maxConnections, _ := strconv.Atoi(temp)

	db.SetMaxIdleConns(maxConnections)
	db.SetMaxOpenConns(maxConnections)

	return db, nil
}
