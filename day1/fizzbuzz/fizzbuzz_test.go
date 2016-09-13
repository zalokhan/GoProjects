package fizzbuzz_test

import (
	"fmt"
	"testing"

	"github.com/zalokhan/day1/fizzbuzz"
)

type tuple struct {
	number      int
	expectation string
}

var useCases = [...]tuple{
	{1, " 1"},
	{3, " 1 2 fizz"},
	{4, " 1 2 fizz 4"},
	{5, " 1 2 fizz 4 buzz"},
	{6, " 1 2 fizz 4 buzz fizz"},
	{15, " 1 2 fizz 4 buzz fizz 7 8 fizz buzz 11 fizz 13 14 fizzbuzz"},
}

func TestFizzbuzz(t *testing.T) {
	for _, useCase := range useCases {
		actual := fizzbuzz.Fizzbuzz(useCase.number)
		if useCase.expectation != actual {
			t.Fatalf("(%d) Expected %s but got %s", useCase.number, useCase.expectation, actual)
		}
	}
}

func BenchmarkFizzbuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, useCase := range useCases {
			actual := fizzbuzz.Fizzbuzz(useCase.number)
			if useCase.expectation != actual {
				b.Fatalf("(%d) Expected %s but got %s", useCase.number, useCase.expectation, actual)
			}
		}
	}
}

func ExampleFizzbuzz() {
	fmt.Println(fizzbuzz.Fizzbuzz(15))
	//Output:
	// 1 2 fizz 4 buzz fizz 7 8 fizz buzz 11 fizz 13 14 fizzbuzz
}
