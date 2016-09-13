package main

import "fmt"

func matrixAdd(a, b *[3][3]int) [3][3]int {
	result := [3][3]int{}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			result[i][j] = a[i][j] + b[i][j]
		}
	}
	return result
}

func matrixMult(a, b *[3][3]int) [3][3]int {
	result := [3][3]int{}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			for k := 0; k < len(a); k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return result
}

func main() {
	a := [3][3]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}
	b := [3][3]int{{4, 4, 4}, {5, 5, 5}, {6, 6, 6}}
	fmt.Println(" a + b = ", matrixAdd(&a, &b))

	fmt.Println(" a * b = ", matrixMult(&a, &b))

}
