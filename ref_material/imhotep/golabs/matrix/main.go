/*
Write a program (with tests!) to compute 3x3 matrix addition and multiplication.
Extra credit:support substract and divide operators
*/
package main

import "fmt"

type operator int

const (
	add operator = iota
	multiply
	substract
	divide
	_
	equal
)

func compute(m1, m2 [3][3]int, op operator) (res [3][3]int) {
	switch op {
	case add:
		res = mxStraight(m1, m2, false)
	case multiply:
		res = mxDot(m1, m2, false)
	case divide:
		res = mxDot(m1, m2, true)
	default:
		res = mxStraight(m1, m2, true)
	}
	return
}

func mxStraight(m1, m2 [3][3]int, inverse bool) (res [3][3]int) {
	for i := 0; i < len(m1); i++ {
		for j := 0; j < len(m1[i]); j++ {
			if inverse {
				res[i][j] = m1[i][j] - m2[i][j]
			} else {
				res[i][j] = m1[i][j] + m2[i][j]
			}
		}
	}
	return
}

func mxDot(m1, m2 [3][3]int, inverse bool) (res [3][3]int) {
	for n := 0; n < len(m1); n++ {
		for m := 0; m < len(m1[n]); m++ {
			sum := 0
			for i := 0; i < len(m1); i++ {
				if inverse {
					sum += m1[n][1] * 1 / m2[i][n]
				} else {
					sum += m1[n][i] * m2[i][n]
				}
			}
			res[n][m] = sum
		}
	}
	return
}

func printRow(m [3]int) {
	for j := 0; j < len(m); j++ {
		fmt.Printf("% 3d ", m[j])
	}
}

func printOp(i int, op operator) {
	if i == 1 {
		fmt.Printf(" %s ", opToHuman(op))
	} else {
		fmt.Printf("%3s", " ")
	}
}

func display(m1, m2, m3 [3][3]int, op operator) {
	for i := 0; i < len(m1); i++ {
		printRow(m1[i])
		printOp(i, op)
		printRow(m2[i])
		printOp(i, equal)
		printRow(m3[i])
		fmt.Println()
	}
}

func opToHuman(op operator) (s string) {
	switch op {
	case add:
		s = "+"
	case substract:
		s = "-"
	case multiply:
		s = "*"
	case divide:
		s = "/"
	case equal:
		s = "="
	}
	return
}

func main() {
	m1 := [3][3]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}
	m2 := [3][3]int{{3, 3, 3}, {4, 4, 4}, {5, 5, 5}}

	for _, op := range []operator{add, substract, multiply, divide} {
		display(m1, m2, compute(m1, m2, op), op)
		fmt.Println()
	}
}
