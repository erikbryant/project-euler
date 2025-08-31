package main

import (
	"testing"
)

func TestA(t *testing.T) {
	testCases := []struct {
		c        int
		expected int
	}{
		{1, 1},
		{2, 2},
		{3, -5},
		{4, 4},
		{5, 17},
		{6, -10},
		{7, -17},
		{8, 8},
		{9, -47},
		{10, 34},
	}

	for _, testCase := range testCases {
		answer := A(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}

func TestSumToPow2(t *testing.T) {
	testCases := []struct {
		c        int
		expected int
	}{
		{0, 1},
		{1, 3},
		{2, 2},
		{3, 0},
		{4, -4},
		{5, -12},
		{6, -28},
		{7, -60},
		{8, -124},
	}

	for _, testCase := range testCases {
		answer := sumToPow2(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}

func TestSumBrokenPairs(t *testing.T) {
	testCases := []struct {
		c        int
		expected int
	}{
		{1, 1},
		{2, 3},
		{3, -2},
		{4, 2},
		{5, 19},
		{6, 9},
		{7, -8},
		{8, 0},
		{9, -47},
		{10, -13},
		{11, 34},
		{12, 14},
		{13, 55},
		{14, 21},
		{15, -20},
		{16, -4},
		{17, 145},
		{18, 51},
		{19, -98},
		{20, -30},
		{21, -137},
		{22, -43},
		{23, 64},
		{24, 24},
		{25, -119},
		{26, -37},
		{27, 106},
		{28, 38},
		{29, 127},
		{30, 45},
		{31, -44},
		{32, -12},
		{33, -443},
		{34, -145},
		{35, 286},
		{36, 98},
		{37, 451},
		{38, 153},
		{39, -200},
	}

	for _, testCase := range testCases {
		answer := sumBrokenPairs(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}
