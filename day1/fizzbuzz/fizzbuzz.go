//Package fizzbuzz package which generates fizzes or buzzes based on parameter
package fizzbuzz

import "strconv"

func getAnswer(number int) string {
	if (number%3 == 0) && (number%5 == 0) {
		return " fizzbuzz"
	} else if number%5 == 0 {
		return " buzz"
	} else if number%3 == 0 {
		return " fizz"
	}
	res := " "
	res += strconv.Itoa(number)
	return res
}

//Fizzbuzz generates fizz buzz string
func Fizzbuzz(n int) string {
	res := ""
	for i := 1; i <= n; i++ {
		res += getAnswer(i)
	}
	return res
}
