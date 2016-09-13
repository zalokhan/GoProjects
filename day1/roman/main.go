package main

import "fmt"

var romanGlyph = [...]struct {
	limit int
	glyph string
} {
	{10:	"X"},
	{9:		"IX"},
	{5:		"V"},
	{4:		"IV"},
	{1:		"I"},
}

func arabicToRoman(n int) string {

	result := ""

	for _, tuple := range romanGlyph {
		for n >= tuple.limit {
			result += tuple.glyph
			n -= tuple.limit
		}
	}

	// for n >= 10 {
	// 	result += "X"
	// 	n -= 10
	// }

	// if n >= 9 {
	// 	result += "IX"
	// 	n -= 9
	// }

	// if n >= 5 {
	// 	result += "V"
	// 	n -= 5
	// }

	// for i := 0; i < n; i++ {
	// 	result += "I"
	// }

	return result
}

func main() {
	fmt.Println(arabicToRoman(1))
}