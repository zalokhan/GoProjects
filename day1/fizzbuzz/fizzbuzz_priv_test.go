package fizzbuzz

import (
	"testing"
)

type tuple struct {
	number      int
	expectation string
}

var useCases = [...]tuple{
	{1, " 1"},
	{3, " fizz"},
	{4, " 4"},
	{5, " buzz"},
	{15, " fizzbuzz"},
}

func TestFizzbuzzGetAnswer(t *testing.T) {
	for _, useCase := range useCases {
		actual := getAnswer(useCase.number)
		if useCase.expectation != actual {
			t.Fatalf("(%d) Expected %s but got %s", useCase.number, useCase.expectation, actual)
		}
	}
}

func BenchmarkFizzbuzzGetAnswer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, useCase := range useCases {
			actual := getAnswer(useCase.number)
			if useCase.expectation != actual {
				b.Fatalf("(%d) Expected %s but got %s", useCase.number, useCase.expectation, actual)
			}
		}
	}
}
