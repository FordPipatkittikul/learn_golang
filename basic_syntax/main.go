package main

import (
	"basic_syntax/customer"
	"fmt"
	"unicode/utf8"
)

// package scope
type Person struct {
    Name string
    Age  int
}

// method of Person struct
func (p Person) introduce() string {
	return fmt.Sprintf("Hello, %s! You are %d years old.", p.Name, p.Age)
}

func main() {

	// ------------------------------------------------------------
	point := 10
	_ = point // we need to use point to avoid compiler error

	if point > 5 {
		fmt.Println("point greater than 5")
	} else {
		fmt.Println("point 5 or less")
	}

	// if you want to count string using utf8.RuneCountInString
	str := "Hello, 世界"
	fmt.Println("String length in bytes:", len(str))
	fmt.Println("String length in runes:", utf8.RuneCountInString(str))
	fmt.Printf("-----\n")
	// ------------------------------------------------------------

	
	
	// ------------------------------------------------------------
	// In Go, an array is a fixed-length, 
	// contiguous sequence of elements all of the same data type
	x := [3]int{1, 2, 3}
	y := [3]int{}
	z := [...]int{1,2,3,4,5}
	g := z[1:]
	
	fmt.Println("Array length:", len(x))
	fmt.Printf("%#v\n", y) // \n foor new line
	fmt.Printf("%#v\n", z)
	fmt.Printf("%#v\n", g)
	fmt.Printf("-----\n")
	// ------------------------------------------------------------



	// ------------------------------------------------------------
	// In Go, a slice is a dynamically-sized, 
	// flexible "view" into the elements of an underlying array
	sliceA := []int{1, 2, 3}
	sliceA = append(sliceA, 10)
	sliceB := append(sliceA, 10)
	sliceC := sliceB[3:4]
	fmt.Printf("%#v\n, len=%v\n", sliceB, len(sliceB))
	fmt.Printf("%#v\n", sliceC)
	fmt.Printf("-----\n")
	// ------------------------------------------------------------



	// ------------------------------------------------------------
	// In Go, a map is an unordered collection of key-value pairs
	m := map[string]int{
		"apple":  5,
		"banana": 10,
	}
	m["orange"] = 15
	fmt.Printf("%#v\n", m)
	fruit, amount := m["grape"]
	if amount {
		fmt.Println("grape:", fruit)
	} else {
		fmt.Println("grape not found")
	}
	// ------------------------------------------------------------



	// ------------------------------------------------------------
	values := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(values); i++{
		fmt.Println("value:", values[i])
	}

	// same as while loop
	index := 0
	for index < len(values) {
		fmt.Println("while value:", values[index])
		index++
	}

	// equivalent to for each loop in other languages
	for indexing, v := range values {
		fmt.Println("index:", indexing, "value:", v)
	}
	// if you don't need the index
	for _, v := range values {
		fmt.Println("value:", v)
	}
	// ------------------------------------------------------------



	// ------------------------------------------------------------
	// function คือ type หนึ่งใน Go
	answer := sum(50,50)
	fmt.Printf("%#v\n", answer)

	// Anonymous function
	greet := func(name string) {
        fmt.Printf("Hello, %v!\n", name)
    }

    greet("Alice") // Call the anonymous function via the variable
	cal(sum)

	// (Variadic Functions)
	greeting("Hello", "World", "!")
	// ------------------------------------------------------------



	// ------------------------------------------------------------
	// การใช้งาน package อื่นที่ไม่ใช่ package main
	fmt.Printf("%v", customer.Name)
	// ------------------------------------------------------------



	// ------------------------------------------------------------
	// pointers
	a := 42
	b := &a // b holds the memory address of a
	*b = 21 // dereferencing b to set the value of a to 21
	fmt.Println(" ")
	fmt.Println("a:", a)
	fmt.Println("b (address of a):", b)
	fmt.Println("value at address b:", *b) // dereferencing b to get the value of a

	xx := 5
	minus1(&xx) // pass by reference
	fmt.Println("After minus1, xx:", xx) // xx is now 0
	yy := 5
	minus2(yy) // pass by value
	fmt.Println("After minus2, yy:", yy) // yy remains 5
	// ------------------------------------------------------------



	// ------------------------------------------------------------
	// structs and methods
	// In Go (Golang), a struct (short for structure) is a composite data type that allows you to group together related values of different types under a single name.
	// Create an instance of the Person struct
    p := Person{Name: "Alice", Age: 30} 
	
	// Access individual fields using the dot operator
    fmt.Println(p.Name) // Output: Alice
    fmt.Println(p.Age)  // Output: 30

    // Structs are mutable
    p.Age = 31
    fmt.Println(p.Age) // Output: 31

	// method call
	fmt.Println(p.introduce())


	Secret := customer.Secret{}
	Secret.SetName("John Doe")
	Secret.SetAge(28)
	fmt.Println("Secret Name:", Secret.GetName())
	fmt.Println("Secret Age:", Secret.GetAge())
	// ------------------------------------------------------------


	
}


func sum(a int, b int) int {
	return a + b
}

func cal(f func(int, int) int) {
	fmt.Printf("%#v\n", f(10,30))
}

// greeting accepts a varying number of string arguments
func greeting(messages ...string) {
    for i, msg := range messages {
        fmt.Println(i, msg)
    }
}

// pass by reference
func minus1(result * int) {
	*result = 20 - 20
}

// pass by value
func minus2(result int) {
	result = 20 - 20
}