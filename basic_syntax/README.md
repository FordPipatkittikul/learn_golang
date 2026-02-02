# Go Fundamentals – Notes & Examples

This repository contains **concise notes and runnable examples** covering core **Go (Golang) fundamentals**. The goal is to understand Go by reading and executing real code rather than memorizing theory.

---

## 📌 Topics Covered

### 1. Program Structure

* `package main`
* `import` statements
* `func main()` as the entry point

```go
package main

func main() {
    // program starts here
}
```

---

### 2. Variables & Control Flow

* Short variable declaration (`:=`)
* Blank identifier (`_`)
* `if / else` conditions

```go
point := 10
_ = point

if point > 5 {
    fmt.Println("point greater than 5")
}
```

---

### 3. Strings & UTF-8

* `len(string)` → byte length
* `utf8.RuneCountInString` → character length

```go
str := "Hello, 世界"
len(str)                       // bytes
utf8.RuneCountInString(str)    // runes (characters)
```

---

### 4. Arrays

* Fixed-size
* Same data type
* Length is part of the type

```go
x := [3]int{1, 2, 3}
z := [...]int{1, 2, 3, 4, 5}
```

---

### 5. Slices

* Dynamically sized
* Built on top of arrays
* Commonly used instead of arrays

```go
slice := []int{1, 2, 3}
slice = append(slice, 10)
```

---

### 6. Maps

* Key–value data structure
* Unordered
* Comma-ok idiom for safe lookup

```go
m := map[string]int{"apple": 5}
value, ok := m["apple"]
```

---

### 7. Loops

Go has only one loop keyword: `for`

```go
for i := 0; i < 5; i++ {}

for i, v := range values {}
```

---

### 8. Functions

* Functions are first-class citizens
* Anonymous functions
* Higher-order functions
* Variadic functions

```go
func sum(a, b int) int {
    return a + b
}

func greeting(msg ...string) {}
```

---

### 9. Packages

* Code is organized into packages
* Only exported identifiers (capitalized) are accessible

```go
import "basic_syntax/customer"
```

---

### 10. Pointers

* Go is pass-by-value
* Pointers allow modifying original values

```go
a := 10
b := &a
*b = 20
```

---

### 11. Structs

* Composite data type
* Groups related fields

```go
type Person struct {
    Name string
    Age  int
}
```
* Create an instance of the Person struct
```go
p := Person{Name: "Alice", Age: 30} 
```
---

### 12. Methods

* Functions with receivers
* Used to attach behavior to structs

```go
func (p Person) introduce() string {
    return "Hello " + p.Name
}
```

---

### 13. Encapsulation

Encapsulation in Go is achieved using **packages** and **exported / unexported identifiers**, not `private` or `protected` keywords.

### Rule

* **Capitalized** names → exported (public)
* **Lowercase** names → unexported (package-private)

---

### Example: Encapsulation with a Package

**Folder structure**

```
encapsulation/
├── main.go
└── user/
    └── user.go
```

---

### `user/user.go`

```go
package user

import "fmt"

// User has unexported fields
type User struct {
    name string
    age  int
}

// Constructor
func NewUser(name string, age int) *User {
    return &User{name: name, age: age}
}

// Getter
func (u *User) Name() string {
    return u.name
}

// Setter with validation
func (u *User) SetAge(age int) error {
    if age < 0 {
        return fmt.Errorf("age cannot be negative")
    }
    u.age = age
    return nil
}

// Getter
func (u *User) Age() int {
    return u.age
}
```

---

### `main.go`

```go
package main

import (
    "fmt"
    "encapsulation/user"
)

func main() {
    u := user.NewUser("Alice", 30)

    fmt.Println(u.Name())
    fmt.Println(u.Age())

    // u.age = 40 ❌ compile error

    _ = u.SetAge(35)
    fmt.Println("Updated age:", u.Age())
}
```

---

### Why this is idiomatic Go

* State is protected
* Validation is centralized
* Clean public API
* No unnecessary getters/setters

---

## 14. Interfaces

Interfaces define **behavior**, not data. A type implements an interface **implicitly** by implementing its methods.

```go
type Greeter interface {
    Greet() string
}

type User struct {
    Name string
}

func (u User) Greet() string {
    return "Hello, " + u.Name
}
```

Why it matters:

* Enables loose coupling
* Replaces inheritance
* Core to clean architecture

---

## 15. Error Handling

Go handles errors explicitly using the built-in `error` type.

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

result, err := divide(10, 0)
if err != nil {
    fmt.Println("error:", err)
    return
}
```

Key ideas:

* No exceptions
* Always check `if err != nil`
* Errors are values

---

## 16. Defer

`defer` schedules a function to run after the surrounding function returns.

```go
def fmt.Println("start")
defer fmt.Println("end")
```

Uses:

* Closing files
* Unlocking mutexes
* Cleaning up resources

---

## 17. Zero Values

Every type in Go has a **zero value**. Understanding zero values is critical for writing safe Go code.

### 🔹 Value Types (❌ cannot be nil)

These always have a concrete value (zero value if not set).

| Type           | Example                   | Zero Value                       |
| -------------- | ------------------------- | -------------------------------- |
| `bool`         | `true`                    | `false`                          |
| Integers       | `int`, `int64`, `uint`    | `0`                              |
| Floating point | `float32`, `float64`      | `0.0`                            |
| Complex        | `complex64`, `complex128` | `0+0i`                           |
| `string`       | `"abc"`                   | `""` (empty string, **not nil**) |
| `array`        | `[3]int`                  | `[0 0 0]`                        |
| `struct`       | `struct{}`                | fields’ zero values              |
| `uintptr`      |                           | `0`                              |

---

### 🔹 Reference / Nil-capable Types (✅ can be nil)

These types may hold `nil` and must be checked before use.

| Type                 | Can be `nil` | Notes                              |
| -------------------- | ------------ | ---------------------------------- |
| `pointer` (`*T`)     | ✅            | Most common for optional values    |
| `slice` (`[]T`)      | ✅            | `nil` slice ≠ empty slice          |
| `map` (`map[K]V`)    | ✅            | Writing to nil map panics          |
| `channel` (`chan T`) | ✅            | Send/receive blocks forever        |
| `function` (`func`)  | ✅            | Calling nil func panics            |
| `interface`          | ✅            | Subtle: typed nil vs nil interface |

---

### Important Notes

* `nil` is **not** a zero value for value types
* `nil` slices are safe to read, but not maps
* Interfaces can be tricky due to **typed nil**

| Data Type | Zero Value |
| :--- | :--- |
| `int` | `0` |
| `string` | `""` |
| `bool` | `false` |
| `pointer` | `nil` |
| `struct` | `Zero-value fields` |

---

## 18. Concurrency Basics

Go uses **goroutines** and **channels** for concurrency.

### Goroutines

```go
go func() {
    fmt.Println("running concurrently")
}()
```

### Channels

```go
ch := make(chan int)

go func() {
    ch <- 42
}()

value := <-ch
fmt.Println(value)
```

Key ideas:

* Goroutines are lightweight threads
* Channels communicate between goroutines
* "Do not communicate by sharing memory; share memory by communicating"

---



