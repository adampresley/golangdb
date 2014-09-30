# Golang DB

This package is a simple abstraction I use to allow my projects to connect to various databases through a single interface. 

## Version 0.3

## MySQL
```golang
import (
    "log"

    "github.com/adampresley/golangdb"
)

connection := DatabaseConnection{
    Engine:   golangdb.MYSQL,
    Address:  "localhost",
    Port:     3306,
    Database: "test",
    UserName: "user",
    Password: "password",
}

if err := connection.Connect("main"); err != nil {
    log.Fatalf("Error connecting to database: %s", err.Error())
}

// Database handle now lives in golangdb.Db["main"]
```

## MSSQL
```golang
import (
    "log"

    "github.com/adampresley/golangdb"
)

connection := DatabaseConnection{
    Engine:   golangdb.MSSQL,
    Address:  "localhost",
    Port:     1433,
    Database: "test",
    UserName: "user",
    Password: "password",
}

if err := connection.Connect("main"); err != nil {
    log.Fatalf("Error connecting to database: %s", err.Error())
}

// Database handle now lives in golangdb.Db["main"]
```

## SQLite
```golang
import (
    "log"

    "github.com/adampresley/golangdb"
)

connection := DatabaseConnection{
    Engine:   golangdb.SQLITE,
    Database: "./test.db",
}

if err := connection.Connect("main"); err != nil {
    log.Fatalf("Error connecting to database: %s", err.Error())
}

// Database handle now lives in golangdb.Db["main"]
```

## TestDB
```golang
import (
    "log"

    "github.com/adampresley/golangdb"
)

connection := DatabaseConnection{
    Engine: golangdb.TESTDB,
}

if err := connection.Connect("test"); err != nil {
    log.Fatalf("Error connecting to database: %s", err.Error())
}

// Database handle now lives in golangdb.Db["test"]
```

## History

* 2014-09-29
    - Supports multiple connections by name
    - Added support for [go-testdb](https://github.com/erikstmartin/go-testdb)
* 2014-09-17
    - Initial release

## License
The MIT License (MIT)

Copyright (c) 2014 Adam Presley

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

