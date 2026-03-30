# Go: defer, panic, recover (Complete Guide)

---

# defer — Run Later

`defer` schedules a function to run **after the current function finishes**, no matter how it exits (return or panic).

## Example

```go
func main() {
    fmt.Println("start")

    defer fmt.Println("cleanup")

    fmt.Println("end")
}
```

### Output
```
start
end
cleanup
```

---

## Rules of defer

### 1. LIFO (Last-In, First-Out)
```go
defer fmt.Println("1")
defer fmt.Println("2")
```

Output:
```
2
1
```

---

### 2. Arguments evaluated immediately
```go
x := 10
defer fmt.Println(x)
x = 20
```

Output:
```
10
```

---

## Common Use Cases

### Close resources
```go
file, _ := os.Open("file.txt")
defer file.Close()
```

### Unlock mutex
```go
mu.Lock()
defer mu.Unlock()
```

### Logging
```go
defer fmt.Println("function finished")
```

---

#  panic — Crash Immediately

`panic` stops normal execution and begins **stack unwinding**.

## Example

```go
func main() {
    fmt.Println("start")
    panic("something went wrong")
    fmt.Println("end") // never runs
}
```

Output:
```
start
panic: something went wrong
```

---
### When to use panic 

Use panic only for unexpected, unrecoverable errors:

Good use cases
- Programmer mistakes
```go
if db == nil {
    panic("db should not be nil")
}
```
- Impossible states
```go
default:
    panic("unknown type")
```    
- Initialization failures
```go
config := loadConfig()
if config == nil {
    panic("failed to load config")
}
```
Bad use cases

- Normal error handling 
```go
// BAD

if err != nil {
    panic(err)
}
```
Use return error instead:
```go
if err != nil {
    return err
}
```
---
# recover — Catch Panic

`recover` prevents a panic from crashing the program.  
Must be used inside a `defer`.

```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered:", r)
        }
    }()

    panic("boom")
}
```

Output:
```
Recovered: boom
```

---

#  Stack Unwinding (IMPORTANT)

When a panic occurs:
1. Current function stops immediately
2. All deferred functions in that function run
3. Then Go moves to the caller function
4. Repeat until:
   - recovered OR
   - program crashes

---

## Example: Multiple Function Calls

```go
func c() {
    defer fmt.Println("defer in c")
    panic("panic in c")
}

func b() {
    defer fmt.Println("defer in b")
    c()
}

func a() {
    defer fmt.Println("defer in a")
    b()
}

func main() {
    defer fmt.Println("defer in main")
    a()
}
```

---

##  Output

```
defer in c
defer in b
defer in a
defer in main
panic: panic in c
```

---

## Key Insight

Even though panic starts in `c()`:

- `c` defer runs
- then `b` defer
- then `a` defer
- then `main` defer

👉 This is **stack unwinding**

---

##  With recover

```go
func b() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in b:", r)
        }
    }()

    c()
}
```

### Output

```
defer in c
Recovered in b: panic in c
defer in a
defer in main
```

👉 Panic stops at `b()` — program continues normally

---

# 🌐 Real-World Usage (Fiber HTTP)

---

## Without Recovery Middleware (BAD)

```go
func handler(c *fiber.Ctx) error {
    panic("something broke")
}
```

👉 This crashes the server ❌

---

## With defer + recover (Middleware)

```go
func RecoveryMiddleware(c *fiber.Ctx) error {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    return c.Next()
}
```

---

## Register Middleware

```go
app := fiber.New()
app.Use(RecoveryMiddleware)
```

---

## Handler Example

```go
func handler(c *fiber.Ctx) error {
    panic("unexpected error")
}
```

Output:
```
Recovered from panic: unexpected error
```

---

## Better Version (Return HTTP Response)

```go
func RecoveryMiddleware(c *fiber.Ctx) error {
    defer func() {
        if r := recover(); r != nil {
            c.Status(fiber.StatusInternalServerError).
                JSON(fiber.Map{
                    "error": "internal server error",
                })
        }
    }()

    return c.Next()
}
```

---

# Best Practices

## Use defer for:
- closing files
- unlocking mutex
- logging
- cleanup

---

## Use panic only for:
- impossible states
- programmer mistakes
- startup failures

---


| Concept  | Meaning                |
|----------|------------------------|
| defer    | run later (finally)    |
| panic    | crash now              |
| recover  | catch panic (in defer) |



# TL;DR

- `defer` = cleanup always runs  
- `panic` = stop execution + unwind stack  
- `recover` = stop panic from crashing program  
- stack unwinding = defers run from inner → outer  
- in web apps → always use recovery middleware  