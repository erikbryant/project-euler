package main

import (
	"testing"
)

func TestRadical(t *testing.T) {
	testCases := []struct {
		c        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 2},
		{5, 5},
		{6, 6},
		{7, 7},
		{8, 2},
		{9, 3},
		{10, 10},
	}

	for _, testCase := range testCases {
		answer := radical(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}
