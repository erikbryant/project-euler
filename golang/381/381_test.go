package main

import (
	"testing"
)

// func TestFactorialMod(t *testing.T) {
// 	testCases := []struct {
// 		n        int
// 		expected int
// 	}{
// 		{0, 1},
// 		{1, 1},
// 		{2, 2},
// 		{3, 6},
// 		{4, 24},
// 	}

// 	for _, testCase := range testCases {
// 		answer := factorial(testCase.n)
// 		if answer != testCase.expected {
// 			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
// 		}
// 	}
// }

func TestS(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{5, 4},
		{7, 4},
		{11, 1},
		{13, 11},
		{17, 6},
		{19, 2},
	}

	for _, testCase := range testCases {
		answer := S(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}

func TestSumS(t *testing.T) {
	testCases := []struct {
		min      int
		max      int
		expected int
	}{
		{5, 100, 480},
	}

	for _, testCase := range testCases {
		answer := sumS(testCase.min, testCase.max)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d, %d expected %d, got %d", testCase.min, testCase.max, testCase.expected, answer)
		}
	}
}
