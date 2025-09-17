package main

import (
	"testing"
)

func TestPermuteDivisors(t *testing.T) {
	testCases := []struct {
		n        int
		k        int
		expected int
	}{
		{6, 3, 6},
		{8, 4, 8},
		{15, 9, 15},
		{90, 77, 90},
	}

	for _, testCase := range testCases {
		answer := permuteDivisors(testCase.n, testCase.k)
		if answer != testCase.expected {
			t.Errorf("ERROR: For n=%d, k=%d expected %d, got %d", testCase.n, testCase.k, testCase.expected, answer)
		}
	}
}

func TestFindMinN(t *testing.T) {
	testCases := []struct {
		k        int
		expected int
	}{
		{3, 6},
		{4, 8},
		{9, 15},
		{77, 90},
	}

	for _, testCase := range testCases {
		answer := findMinN(testCase.k)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.k, testCase.expected, answer)
		}
	}
}

// func TestXYZ(t *testing.T) {
// 	testCases := []struct {
// 		k        int
// 		expected int
// 	}{
// 		{0, 0},
// 		{5, 5},
// 		{10, 1},
// 		{25, 7},
// 		{100000, 1},
// 		{100001, 2},
// 	}
//
// 	for _, testCase := range testCases {
// 		answer := XYZ(testCase.k)
// 		if answer != testCase.expected {
// 			t.Errorf("ERROR: For %d expected %d, got %d", testCase.k, testCase.expected, answer)
// 		}
// 	}
// }
