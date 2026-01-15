# learn_golang

### quick command
- go env
- go env GOPATH
- go run <filename.go>

### 🌐 Official Go Websites

Go Language (Official Site)
https://go.dev

- The official Go programming language website. Contains downloads, documentation, tutorials, and news.

Go Packages Index
https://pkg.go.dev

- Official documentation hub for Go packages. Useful for browsing APIs, examples, and version details.

Go Standard Library
https://pkg.go.dev/std

- Documentation for Go’s built-in standard library (e.g., net/http, database/sql, fmt, time).

Go Language Specification
https://go.dev/ref/spec

- Defines the syntax, semantics, and core behavior of the Go language. Useful for deep understanding of how Go works.

SQL Database Drivers for Go
https://go.dev/wiki/SQLDrivers

- A curated list of available SQL drivers compatible with Go’s database/sql package.

### 📦 Go Packages Used
### Database & SQL

database/sql
https://pkg.go.dev/database/sql

- Go’s standard database interface.
- Provides generic SQL database access (connection pooling, querying, transactions) independent of the underlying driver.

sqlx (jmoiron/sqlx)
https://pkg.go.dev/github.com/jmoiron/sqlx

- An extension of database/sql that adds:
- Struct scanning
- Named queries
- Cleaner and more convenient database code

Microsoft SQL Server Driver (go-mssqldb)
https://pkg.go.dev/github.com/denisenkom/go-mssqldb

- A SQL Server driver implementation for Go, used together with database/sql or sqlx.

MySQL Driver (go-sql-driver/mysql)
https://pkg.go.dev/github.com/go-sql-driver/mysql

- A popular and stable MySQL driver for Go, compatible with database/sql.

### Web & Routing

gorilla/mux
https://pkg.go.dev/github.com/gorilla/mux

- A powerful HTTP request router and URL matcher.
- Commonly used to build REST APIs with:
- Path parameters
- HTTP method matching
- Middleware support

Logging

zap (Uber)
https://pkg.go.dev/go.uber.org/zap

- A fast, structured, production-ready logging library.
- Designed for high performance with JSON logs and leveled logging (Info, Debug, Error, etc.).

### Configuration

viper (spf13/viper)
https://pkg.go.dev/github.com/spf13/viper

- A configuration management library that supports:
- Config files (YAML, JSON, TOML)
- Environment variables
- Default values
- Hot reloading (optional)
