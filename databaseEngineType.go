package golangdb

/* Defines a type of database engine */
type DatabaseEngine int

/* Constants used to specify a type of database connection */
const (
	NONE DatabaseEngine = iota
	SQLITE
	MYSQL
	MSSQL
)

var databaseEngineNames = map[DatabaseEngine]string{
	NONE:   "None",
	SQLITE: "SQLite",
	MYSQL:  "MySQL",
	MSSQL:  "MSSQL",
}

func (this *DatabaseEngine) ToString() string {
	return databaseEngineNames[*this]
}