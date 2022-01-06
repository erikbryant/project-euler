package main

import (
	"testing"
)

func TestPeriod(t *testing.T) {
	testCases := []struct {
		c        int
		expected int
	}{
		{0, 0},
		{1, 0},
		{2, 1},
		{3, 2},
		{4, 0},
		{5, 1},
		{6, 2},
		{7, 4},
		{8, 2},
		{9, 0},
		{10, 1},
		{11, 2},
		{12, 2},
		{13, 5},
	}

	for _, testCase := range testCases {
		answer := period(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}
