package services_test

import (
	"fmt"
	"gotest/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckGrade(t *testing.T) {

	// table-driven test with subtests

	type testCase struct {
		name 		string
		score 		int
		expected	string
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

func BenchmarkCheckGrade(b *testing.B){
	for i:=0; i<b.N; i++{
		services.CheckGrade(80)
	}
}

func ExampleCheckGrade() {
	grade := services.CheckGrade(80)
	fmt.Println(grade)
}