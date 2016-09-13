package main

import "fmt"

func matrixAdd(a, b [][]int) [][]int {
	result := make([][]int, 3, 3)

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			result[i] = append(result[i], a[i][j]+b[i][j])
		}
	}
	return result
}

func matrixMult(a, b [][]int) [][]int {
	result := make([][]int, 3, 3)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			var sum int
			for k := 0; k < len(a); k++ {
				sum += a[i][k] * b[k][j]
			}
			result[i] = append(result[i], sum)
		}
	}
	return result
}

func main() {
	a := [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}
	b := [][]int{{4, 4, 4}, {5, 5, 5}, {6, 6, 6}}
	fmt.Println(" a + b = ", matrixAdd(a, b))

	fmt.Println(" a * b = ", matrixMult(a, b))

}
