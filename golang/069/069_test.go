package main

import (
	"testing"
)

func TestTotient(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{2, 1},
		{3, 2},
		{4, 2},
		{5, 4},
		{6, 2},
		{7, 6},
		{8, 4},
		{9, 6},
		{10, 4},
		{11, 10},
		{12, 4},
		{13, 12},
		{14, 6},
		{15, 8},
		{16, 8},
		{17, 16},
		{18, 6},
		{19, 18},
		{20, 9},
		{21, 12},
	}

	for _, testCase := range testCases {
		answer := totient(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}
