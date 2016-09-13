package main

import "flag"

func mult(multiplier int, number int) int {
	return multiplier * number
}

func double(num int, f func(int, int) int) int {
	z := f(num, 2)
	return z
}

func main() {
	flag.Int("u", 0, "Pass multiplier")

}
