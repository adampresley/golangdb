package golangdb

/* Defines a type of database engine */
type DatabaseEngine int

/* Constants used to specify a type of database connection */
const (
	NONE DatabaseEngine = iota
	SQLITE
	MYSQL
	MSSQL
	TESTDB
)

var databaseEngineNames = map[DatabaseEngine]string{
	NONE:   "None",
	SQLITE: "SQLite",
	MYSQL:  "MySQL",
	MSSQL:  "MSSQL",
	TESTDB: "Test (fake) DB",
}

func GetDatabaseEngineFromName(name string) DatabaseEngine {
	for k, v := range databaseEngineNames {
		if v == name {
			return k
		}
	}

	return NONE
}

func (this *DatabaseEngine) ToString() string {
	return databaseEngineNames[*this]
}