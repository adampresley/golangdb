package golangdb

import (
	"database/sql"
)

/* Global database connection handle used throughout your application */
var Db = make(map[string]*sql.DB)
