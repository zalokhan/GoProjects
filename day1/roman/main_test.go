package main

import (
	// "fmt"
	"testing"
)

type tuple struct {
	number int
	expectation string
}

var useCases = [...]tuple {
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{10, "X"},
	{20, "XX"},
}

func TestArabicToRoman(t *testing.T) {
	for _, useCase := range useCases {
		actual := arabicToRoman(useCase.number)

		if actual != useCase.expectation {
			t.Fatalf("(%d) Expected %s but got %s", useCase.number, useCase.expectation, actual)
		}
	}
}