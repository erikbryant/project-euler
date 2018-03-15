package main

import (
	"testing"
)

func TestDiagonals(t *testing.T) {
	testCases := []struct {
		n        int
		expected [4]int
	}{
		{3, [4]int{3, 5, 7, 9}},
		{5, [4]int{13, 17, 21, 25}},
		{7, [4]int{31, 37, 43, 49}},
	}

	for _, testCase := range testCases {
		answer := diagonals(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}
