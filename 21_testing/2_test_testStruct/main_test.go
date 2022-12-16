package main

import "testing"

func TestSum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}
	tests := []test{
		test{[]int{3, 1}, 4},
		test{[]int{5, 5}, 10},
		test{[]int{6, -3, 0}, 3},
		test{[]int{4, 5}, 9}, 
		test{[]int{5, 10, 1}, 16},
	}

	for _, test := range tests {
		total := sum(test.data...)
		if total != test.answer {
			t.Error("Expected :", test.answer, "Got", total)
		}
	}
}
