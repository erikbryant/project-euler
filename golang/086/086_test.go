package main

import (
	// "testing"
	"testing"
)

func TestCuboidPaths(t *testing.T) {
	testCases := []struct {
		c        int64
		expected int64
	}{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 2},
		{4, 1},
		{5, 0},
		{6, 3},
		{7, 0},
		{8, 4},
		{9, 4},
		{10, 0},
		{11, 0},
		{12, 11},
		{13, 0},
		{14, 0},
		{99, 60},
		{100, 85},
	}

	for _, testCase := range testCases {
		answer := cuboidPaths(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}
