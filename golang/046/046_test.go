package main

import (
	"testing"
)

func TestGoldbach(t *testing.T) {
	testCases := []struct {
		n      int
		prime  int
		square int
	}{
		// It turns out there are multiple solutions for some numbers
		{9, 7, 1},
		// {15, 7, 2},
		{15, 13, 1},
		// {21, 3, 3},
		{21, 19, 1},
		// {25, 7, 3},
		{25, 23, 1},
		{27, 19, 2},
		{33, 31, 1},
	}

	for _, testCase := range testCases {
		answerPrime, answerSquare := goldbach(testCase.n)
		if answerPrime != testCase.prime || answerSquare != testCase.square {
			t.Errorf("ERROR: For %d expected %d/%d, got %d/%d", testCase.n, testCase.prime, testCase.square, answerPrime, answerSquare)
		}
	}
}
