# Go Testing Example (Unit, Benchmark, Example)

This project demonstrates how to write **unit tests**, **benchmarks**, and **examples** in Go using a simple grading function.

---

### `CheckGrade`

```go
func CheckGrade(score int) string {
	switch {
	case score >= 80:
		return "A"
	case score >= 70:
		return "B"
	case score >= 60:
		return "C"
	case score >= 50:
		return "D"
	default:
		return "F"
	}
}
```

---

### 1. Unit Test

* Uses **table-driven test**
* Uses **subtests (`t.Run`)**

```go
func TestCheckGrade(t *testing.T) {
	type testCase struct {
		name     string
		score    int
		expected string
	}

	cases := []testCase{
		{name: "test grade A", score: 80, expected: "A"},
		{name: "test grade B", score: 70, expected: "B"},
		{name: "test grade C", score: 60, expected: "C"},
		{name: "test grade D", score: 50, expected: "D"},
		{name: "test grade F", score: 30, expected: "F"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {

			grade := services.CheckGrade(c.score)
			assert.Equal(t, c.expected, grade)
			
		})
	}
}
```

---

### 2. Benchmark Test

* Measures performance
* Uses `b.N` loop

```go
func BenchmarkCheckGrade(b *testing.B){
	for i:=0; i<b.N; i++{
		services.CheckGrade(80)
	}
}
```

---

### 3. Example Test

* Used for documentation
* Runs with `go test`
* Can be shown in `godoc`

```go
func ExampleCheckGrade() {
	grade := services.CheckGrade(80)
	fmt.Println(grade)
}
```

---

### Run Unit Tests

```bash
go test ./...

=== RUN   TestCheckGrade
=== RUN   TestCheckGrade/test_grade_A
=== RUN   TestCheckGrade/test_grade_B
...
--- PASS: TestCheckGrade (0.00s)
PASS
```

### Run Tests (Verbose)

```bash
go test -v ./...
```

---

### Run Test Coverage

```bash
go test -cover ./...
```

---

### Run Benchmark

```bash
go test -bench=.
```

### Run Benchmark with Memory Allocation

```bash
go test -bench=. -benchmem
```

---

### Run Specific Test

```bash
go test -run TestCheckGrade
```

---

### Run Specific Benchmark

```bash
go test -bench BenchmarkCheckGrade
```

---

### Run Example

```bash
go test -run Example
```

---

### VS Code Configuration (Optional)

Enable coverage highlighting in VS Code:

```json
"go.coverOnSave": true,
"go.coverOnSingleTest": true,
"go.coverageDecorator": {
    "type": "gutter",
    "coveredHighlightColor": "rgba(64,128,128,0.5)",
    "uncoveredHighlightColor": "rgba(128,64,64,0.25)",        
    "coveredGutterStyle": "blockgreen",
    "uncoveredGutterStyle": "blockred"
}
```

---

### Notes

* **Unit Test** → correctness
* **Benchmark** → performance
* **Example** → documentation + usage
* Always keep tests **simple, readable, and deterministic**

---

### Tips

* Use table-driven tests for scalability
* Use subtests (`t.Run`) for better debugging
* Keep benchmarks isolated (no I/O, no network)

---


# Integration test

### Build Tag (Important)
```
//go:build integration
```
This means the test will NOT run by default.

Run it with:
```
go test <filepath> -tags=integration
```
Good practice:

- Unit tests = fast, default
- Integration tests = optional, slower


# Mock

### Key Concepts (Creation)
- mock.Mock → gives you .On(), .Return(), .Called()
- You MUST implement the same interface method
- m.Called(...) = capture input + return fake output

```go

// promotion file

package repositories

type Promotion struct {
	ID				int
	PurchaseMin 	int
	DiscountPercent int
}

type PromotionRepository interface {
	GetPromotion() (Promotion, error)
}


// promotion_mock file
package repositories

import "github.com/stretchr/testify/mock"

type promotionRepositoryMock struct {
	mock.Mock
}

func NewPromotionRepositoryMock() *promotionRepositoryMock{
	return &promotionRepositoryMock{}
}

func (m *promotionRepositoryMock) GetPromotion() (Promotion, error) {
	args := m.Called()
	return args.Get(0).(Promotion), args.Error(1)
}
```

- .On() → define expectation
- .Return() → define output

```go
func TestPromotionCalculateDiscount(t *testing.T) {
	
	type testCase struct {
		name 				string
		PurchaseMin 		int
		DiscountPercent		int
		amount 				int
		expected			int
	}

	cases := []testCase{
		{name: "applied 100", PurchaseMin: 100, DiscountPercent: 20, amount: 100, expected: 80},
		{name: "applied 200", PurchaseMin: 100, DiscountPercent: 20, amount: 200,expected: 160},
		{name: "applied 300", PurchaseMin: 100, DiscountPercent: 20, amount: 300,expected: 240},
		{name: "not applied", PurchaseMin: 100, DiscountPercent: 20, amount: 50,expected: 50},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {

			// Arrange
			promoRepo := repositories.NewPromotionRepositoryMock()
			promoRepo.On("GetPromotion").Return(repositories.Promotion{
				ID:					1,
				PurchaseMin: 		c.PurchaseMin,
				DiscountPercent: 	c.DiscountPercent,
			}, nil)
			promoService := services.NewPromotionService(promoRepo)

			// Act
			discount, _ := promoService.CalculateDiscount(c.amount)
	
			// Assert
			assert.Equal(t, c.expected, discount)
		})
	}
}
```