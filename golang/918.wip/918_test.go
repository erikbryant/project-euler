package main

import (
	"testing"
)

func TestSumCache(t *testing.T) {
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
	}

	for _, testCase := range testCases {
		a = make([]int, testCase.c+1)
		answer, _ := SumCache(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}

func TestSumCalc(t *testing.T) {
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
	}

	a = []int{0, 1}

	for _, testCase := range testCases {
		answer := SumCalc(1, testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}

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

	a = []int{0, 1}

	for _, testCase := range testCases {
		answer := A(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}
