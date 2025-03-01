package main

import "testing"

func TestF(t *testing.T) {
	testCases := []struct {
		n        int64
		k        int64
		expected int64
	}{
		{0, 0, 0},
		{1, 1, 1},
		{6, 1, 12},
		{7, 2, 15},
		{8, 4, 12},
		{9, 5, 14},
		{10, 3, 29},
		{10, 5, 24},
		{11, 6, 27},
		{12, 1, 33},
		{13, 3, 27},
		{14, 7, 45},
		{15, 7, 22},
		{16, 10, 40},
	}

	for _, testCase := range testCases {
		answer := f(testCase.n, testCase.k)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d, %d expected %d, got %d", testCase.n, testCase.k, testCase.expected, answer)
		}
	}
}
