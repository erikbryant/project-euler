package main

import (
	"testing"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		c        int
		expected int
	}{
		{4, 3},
	}

	for _, testCase := range testCases {
		answer := solver(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}
