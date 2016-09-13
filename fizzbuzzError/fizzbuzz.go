/*Fizzbuzz game reloaded. Rework the fizzbuzz game by now returning 3 separate sentinel errors
when the number is divisible by 3, 5 and both 3 and 5.
Ensure the caller is checking these 3 errors and emits the actual "Fizz"/"Buzz"/"FizzBuzz in each
of the cases. Thus from the outside the output is still the same as the original game.

Expectations
FizzBuzz 1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz 16 17 Fizz 19
*/
package main

import (
	"errors"
	"fmt"
	"strconv"
)

var errDivBy3 = errors.New("35") // Define div by 3 error here!
var errDivBy5 = errors.New("5")  // Define div by 5 error here!
var errDivBy35 = errors.New("3") // Define div by 3 and 5 error here!

func _fizzbuzz(n int) (res int, err error) {
	// Enter your code here to return the correct error based on divisibility
	if (n%5 == 0) && (n%3 == 0) {
		return n, errDivBy35
	} else if n%5 == 0 {
		return n, errDivBy5
	} else if n%3 == 0 {
		return n, errDivBy3
	}
	return n, nil
}

func fizzbuzz(n int) string {
	// Check error codes here and return the correct value
	res, err := _fizzbuzz(n)
	if err == errDivBy35 {
		return " fizzbuzz"
	} else if err == errDivBy3 {
		return " fizz"
	} else if err == errDivBy5 {
		return " buzz"
	}
	return strconv.Itoa(res)
}

func main() {
	for i := 0; i < 20; i++ {
		fmt.Printf("%s ", fizzbuzz(i))
	}
}
