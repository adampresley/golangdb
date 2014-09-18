package golangdb

/* Defines a type of database engine */
type DatabaseEngine int

/* Constants used to specify a type of database connection */
const (
	_ DatabaseEngine = iota
	SQLITE
	MYSQL
	MSSQL
)

databaseEngineNames := []string{
	"None",
	"SQLite",
	"MySQL",
	"MSSQL"
}

func (this *DatabaseEngine) ToString() string {
	return databaseEngineNames[this]
}