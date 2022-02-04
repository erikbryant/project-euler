package main

import (
	"testing"
)

func TestMaxSquareDivisor(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{18, 9},
		{19, 1},
		{10, 1},
		{100, 100},
		{100000, 10000},
		{100001, 1},
	}

	for _, testCase := range testCases {
		answer := maxSquareDivisor(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}

func TestSumSquares(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{10, 24},
		{100, 767},
		{1000, 22606},
		{1001, 22607},
	}

	for _, testCase := range testCases {
		answer := sumSquares(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}
