package main

import (
	"testing"
)

func TestTotient(t *testing.T) {
	testCases := []struct {
		c        int
		expected int
	}{
		{0, 0},
		{5, 4},
		{10, 4},
		{25, 20},
		{100000, 40000},
		{100001, 90900},
	}

	for _, testCase := range testCases {
		answer := totient(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}

func TestSumPrimes(t *testing.T) {
	testCases := []struct {
		c        int
		l        int
		expected int
	}{
		{0, 1, 0},
		{8, 4, 12},
		{19, 4, 12},
		{9548418, 25, 9548417},
	}

	for _, testCase := range testCases {
		_, answer := sumPrimes(testCase.c, testCase.l)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}

// func TestXYZ(t *testing.T) {
// 	testCases := []struct {
// 		c        int
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
// 		answer := XYZ(testCase.c)
// 		if answer != testCase.expected {
// 			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
// 		}
// 	}
// }
