# Golang

## GO Command

| Phase             | Commands                     |
| ----------------- | ---------------------------- |
| Development | `go run`, `go fmt`, `go vet` |
| Testing        | `go test`                    |
| Deployment     | `go build`                   |
| Environment    | `go env`                     |

1 go run

```
go run main.go
```
What it does:

Compiles + runs your code instantly (no binary saved)

Example:
```
package main

import "fmt"

func main() {
    fmt.Println("Hello Dev")
}
```
Run:
```
go run main.go
```
Output:
```
Hello Dev
```
Use when:
Quickly testing code

Iterating fast

2 go fmt
```
go fmt ./...
```
What it does:

Automatically formats your code (indentation, spacing)

Example (before):
```
func main(){fmt.Println("hi")}
```
After go fmt:
```
func main() {
    fmt.Println("hi")
}
```
 Use when:

After writing code

Before committing

3 go vet
```
go vet ./...
```
What it does:

Finds potential bugs (not syntax errors)

Example:
```
fmt.Printf("%d", "hello") // WRONG
```
go vet warns:

Printf format %d has arg "hello" of wrong type string

Use when:

Before committing

Catch hidden issues


4 go test
```
go test ./...
```
What it does:

Runs all test files (*_test.go)

Example:
```
// math.go
func Add(a, b int) int {
    return a + b
}
```
```
// math_test.go
package main

import "testing"

func TestAdd(t *testing.T) {
    if Add(2, 3) != 5 {
        t.Error("Expected 5")
    }
}
```
Run:
```
go test
```
Output:
```
PASS
```
Use when:

Validating logic

Before pushing code

5 go build
```
go build -o app
```
What it does:

Compiles code into executable

Example:
```
./app
```
Output:
```
Hello Dev
```
Use when:

Creating production binary

Docker / deployment

Unlike go run, this creates a real file


6 go env
```
go env
```
What it does:

Shows Go environment config

Example output:
```
GOOS="linux"
GOARCH="amd64"
GOPATH="/home/user/go"
GOROOT="/usr/local/go"
```

6.5 go env GOPATH
```
go env GOPATH
```
Output:
```
/home/user/go
```
Use when:

Debugging environment issues

Checking paths

## REAL-WORLD FLOW (IMPORTANT)

- During development:
```
go run main.go
go fmt ./...
go vet ./...
```
- Before pushing code:
```
go test ./...
```
- For deployment:
```
go build -o app
```
## Go Modules Commands

1 Initialize a module
```
go mod init <module-name>
```
Example:
```
go mod init github.com/yourname/myapp
```
 Creates go.mod

2 Add dependencies
```
go get <package>
```
Example:
```
go get github.com/gin-gonic/gin
```
Adds dependency to go.mod and downloads it

3 Update dependencies
```
go get -u ./...
```
Updates all dependencies to latest versions

Or specific:
```
go get -u github.com/gin-gonic/gin
```

4 Clean up dependencies (VERY IMPORTANT)
```
go mod tidy
```
Removes unused deps + adds missing ones

Keeps go.mod and go.sum clean

Interview tip: Always run this before committing

5 Download dependencies
```
go mod download
```
Downloads modules without building

6 Check dependencies
```
go list -m all
```
Shows all modules used

7 Why is this dependency here?
```
go mod why <module>
```

8 Verify dependencies
```
go mod verify
```
Ensures downloaded modules are not corrupted

9 Vendor dependencies
```
go mod vendor
```
Copies dependencies into /vendor folder

Used in enterprise / offline builds

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

