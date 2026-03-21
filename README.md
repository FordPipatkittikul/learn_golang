# Golang

## Command
- go mod init <foldername>
- go run <filename.go>
- go env
- go env GOPATH

## 🌐 Official Go Websites

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

## Go Packages Used

### Database & ORM
- **database/sql** – Go standard SQL interface  
  https://pkg.go.dev/database/sql

- **sqlx (jmoiron/sqlx)** – Extensions for `database/sql` (struct scan, named queries)  
  https://pkg.go.dev/github.com/jmoiron/sqlx

- **go-mssqldb** – Microsoft SQL Server driver for Go  
  https://pkg.go.dev/github.com/denisenkom/go-mssqldb

- **go-sql-driver/mysql** – MySQL driver for Go  
  https://pkg.go.dev/github.com/go-sql-driver/mysql

- **gorm** – ORM library for Go  
  https://pkg.go.dev/gorm.io/gorm
  
  [https://gorm.io/docs/conventions.html](https://gorm.io/docs/)

- **gorm MySQL driver** – MySQL dialect for GORM  
  https://pkg.go.dev/gorm.io/driver/mysql

---

### Web & HTTP
- **net/http** – Go standard HTTP server and client  
  https://pkg.go.dev/net/http

- **net/http/httptest** – HTTP testing utilities  
  https://pkg.go.dev/net/http/httptest

- **gorilla/mux** – HTTP router and URL matcher  
  https://pkg.go.dev/github.com/gorilla/mux

- **fiber v2** – High-performance web framework  
  https://pkg.go.dev/github.com/gofiber/fiber/v2

  https://docs.gofiber.io/

- **fiber jwt** – JWT middleware for Fiber  
  https://pkg.go.dev/github.com/gofiber/jwt/v2

- **fiber adaptor** – Adapter between `net/http` and Fiber  
  https://pkg.go.dev/github.com/gofiber/adaptor/v2

---

### Authentication & Security
- **bcrypt** – Password hashing  
  https://pkg.go.dev/golang.org/x/crypto/bcrypt

- **jwt-go** – JSON Web Token implementation  
  https://pkg.go.dev/github.com/dgrijalva/jwt-go

---

### Logging & Configuration
- **zap** – Structured, high-performance logging  
  https://pkg.go.dev/go.uber.org/zap

- **viper** – Configuration management  
  https://pkg.go.dev/github.com/spf13/viper

---

### Testing & Mocking
- **testify** – Assertions and testing helpers  
  https://pkg.go.dev/github.com/stretchr/testify

- **gomock** – Mocking framework for Go  
  https://pkg.go.dev/github.com/golang/mock/gomock

---

### Messaging, Cache & Resilience
- **redis v8** – Redis client  
  https://pkg.go.dev/github.com/go-redis/redis/v8

- **hystrix-go** – Circuit breaker implementation  
  https://pkg.go.dev/github.com/afex/hystrix-go/hystrix

- **sarama** – Apache Kafka client  
  https://pkg.go.dev/github.com/Shopify/sarama

---

### Utilities & Docs
- **uuid** – UUID generation  
  https://pkg.go.dev/github.com/google/uuid

- **godoc** – Go documentation tool  
  https://pkg.go.dev/golang.org/x/tools/cmd/godoc

---

### gRPC & Protobuf
- **grpc** – Remote procedure call framework  
  https://pkg.go.dev/google.golang.org/grpc

- **protobuf** – Protocol Buffers  
  https://pkg.go.dev/google.golang.org/protobuf

- **protoc-gen-go** – Protobuf Go code generator  
  https://pkg.go.dev/google.golang.org/protobuf/cmd/protoc-gen-go

- **protoc-gen-go-grpc** – gRPC Go code generator  
  https://pkg.go.dev/google.golang.org/grpc/cmd/protoc-gen-go-grpc

