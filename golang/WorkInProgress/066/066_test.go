package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{2, 3},
		{3, 2},
		{5, 9},
		{6, 5},
		{7, 8},
		{13, 649},
	}

	for _, testCase := range testCases {
		x := solve(testCase.n, 0)
		if x != testCase.expected {
			t.Errorf("ERROR: For %d expected x=%d, got x=%d", testCase.n, testCase.expected, x)
		}
	}
}

func TestMaxSolution(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{3, 3},
		{7, 9},
	}

	for _, testCase := range testCases {
		_, answer := maxSolution(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}
