package main

import (
	"testing"
)

func TestFactors(t *testing.T) {
	testCases := []struct {
		n        int
		expected []int
	}{
		{2, []int{}},
		{3, []int{}},
		{4, []int{2}},
		{5, []int{}},
		{6, []int{2, 3}},
		{7, []int{}},
		{8, []int{2}},
		{9, []int{3}},
		{10, []int{2, 5}},
		{11, []int{}},
		{12, []int{2, 3}},
		{210, []int{2, 3, 5, 7}},
		{2310, []int{2, 3, 5, 7, 11}},
	}

	for _, testCase := range testCases {
		answer := factors(testCase.n)
		if len(answer) != len(testCase.expected) {
			t.Errorf("ERROR: For %d expected len=%d, got len=%d %v", testCase.n, len(testCase.expected), len(answer), answer)
		}
		for i := 0; i < len(answer); i++ {
			if answer[i] != testCase.expected[i] {
				t.Errorf("ERROR: For %d expected %v, got %v", testCase.n, testCase.expected, answer)
			}
		}
	}
}

// func TestSeive(t *testing.T) {
// 	testCases := []struct {
// 		n        int
// 		expected int
// 	}{
// 		{30, 8},            // product of first 3 primes
// 		{210, 48},          // product of first 4 primes
// 		{2310, 480},        // product of first 5 primes
// 		{30030, 5760},      // product of first 6 primes
// 		{510510, 92160},    // product of first 7 primes
// 		{9699690, 1658880}, // product of first 8 primes
// 	}

// 	for _, testCase := range testCases {
// 		answer := seive(testCase.n)
// 		if answer != testCase.expected {
// 			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
// 		}
// 	}
// }
