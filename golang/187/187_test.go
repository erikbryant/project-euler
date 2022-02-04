package main

import (
	"testing"
)

func TestSemiprimes(t *testing.T) {
	testCases := []struct {
		c        int
		expected int
	}{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 1},
		{5, 1},
		{6, 2},
		{7, 2},
		{8, 2},
		{9, 3},
		{10, 4},
		{11, 4},
		{12, 4},
		{13, 4},
		{14, 5},
		{15, 6},
		{16, 6},
		{17, 6},
		{18, 6},
		{19, 6},
		{20, 6},
		{21, 7},
		{22, 8},
		{23, 8},
		{24, 8},
		{25, 9},
		{26, 10},
		{27, 10},
		{28, 10},
		{29, 10},
		{30, 10},
		{100, 34},
	}

	for _, testCase := range testCases {
		answer := semiprimes(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}
