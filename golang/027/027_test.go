package main

import (
	"testing"
)

func TestQ(t *testing.T) {
	testCases := []struct {
		a        int
		b        int
		expected int
	}{
		{1, 41, 40},
		{-79, 1601, 80},
	}

	for _, testCase := range testCases {
		answer := q(testCase.a, testCase.b)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d, %d expected %d, got %d", testCase.a, testCase.b, testCase.expected, answer)
		}
	}
}
