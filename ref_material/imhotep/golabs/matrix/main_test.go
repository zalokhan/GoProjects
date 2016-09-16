package main

import "testing"

var useCases = map[operator]struct {
	m1 [3][3]int
	m2 [3][3]int
	m3 [3][3]int
}{
	add: {
		[3][3]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}},
		[3][3]int{{3, 3, 3}, {4, 4, 4}, {5, 5, 5}},
		[3][3]int{{4, 4, 4}, {6, 6, 6}, {8, 8, 8}},
	},
	substract: {
		[3][3]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}},
		[3][3]int{{3, 3, 3}, {4, 4, 4}, {5, 5, 5}},
		[3][3]int{{-2, -2, -2}, {-2, -2, -2}, {-2, -2, -2}},
	},
	multiply: {
		[3][3]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}},
		[3][3]int{{3, 3, 3}, {4, 4, 4}, {5, 5, 5}},
		[3][3]int{{12, 12, 12}, {24, 24, 24}, {36, 36, 36}},
	},
	divide: {
		[3][3]int{{3, 3, 3}, {4, 4, 4}, {5, 5, 5}},
		[3][3]int{{3, 3, 3}, {4, 4, 4}, {5, 5, 5}},
		[3][3]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}},
	},
}

func TestCompute(t *testing.T) {
	for op, tuple := range useCases {
		actual := compute(tuple.m1, tuple.m2, op)

		if actual != tuple.m3 {
			t.Fatalf("Operation %s, expected %#v GOT %#v", opToHuman(op), tuple.m3, actual)
		}
	}
}
