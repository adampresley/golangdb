package golangdb

import (
	"database/sql"
	"fmt"
)

/*
DatabaseConnection contains data necessary to establish a connection
to a database server. Most fields are useful for standard SQL
servers. The DatabaseFile field is used by SQLite.
*/
type DatabaseConnection struct {
	Engine   DatabaseEngine
	Address  string
	Port     int
	Database string
	UserName string
	Password string
}

/*
Connect to a database. This method will use the specified
engine to determine how to connect.
*/
func (this *DatabaseConnection) Connect(connectionName string) error {
	var db *sql.DB
	var err error

	switch this.Engine {
	case SQLITE:
		db, err = ConnectSQLite(this)

	case MYSQL:
		db, err = ConnectMySQL(this)

	case MSSQL:
		db, err = ConnectMSSQL(this)

	case TESTDB:
		db, err = ConnectTestDB(this)
	}

	if err != nil {
		return err
	}

	Db[connectionName] = db
	return nil
}

/*
This method converts a DatabaseConnection structure into a string
appropriate for connecting to a database server.
*/
func (this *DatabaseConnection) ToString() string {
	switch this.Engine {
	case SQLITE:
		return this.toSQLiteString()

	case MYSQL:
		return this.toMySQLString()

	case MSSQL:
		return this.toMSSQLString()

	case TESTDB:
		return ""
	}

	return ""
}

func (this *DatabaseConnection) toMySQLString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?autocommit=true", this.UserName, this.Password, this.Address, this.Port, this.Database)
}

func (this *DatabaseConnection) toMSSQLString() string {
	return fmt.Sprintf("Server=%s;Port=%d;User Id=%s;Password=%s;Database=%s", this.Address, this.Port, this.UserName, this.Password, this.Database)
}

func (this *DatabaseConnection) toSQLiteString() string {
	return this.Database
}
