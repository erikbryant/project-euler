package main

import (
	"testing"
)

func TestOddRoot(t *testing.T) {
	testCases := []struct {
		c         int
		expected1 int
		expected2 int
	}{
		{1, 1, 1},
		{2, 1, 2},
		{3, 3, 1},
		{4, 1, 4},
		{5, 5, 1},
		{6, 3, 2},
		{7, 7, 1},
		{8, 1, 8},
		{9, 9, 1},
		{10, 5, 2},
	}

	a = []int{0, 1}

	for _, testCase := range testCases {
		answer1, answer2 := oddRoot(testCase.c)
		if answer1 != testCase.expected1 {
			t.Errorf("ERROR: For %d expected %d / %d, got %d / %d", testCase.c, testCase.expected1, testCase.expected2, answer1, answer2)
		}
		if answer2 != testCase.expected2 {
			t.Errorf("ERROR: For %d expected %d / %d, got %d / %d", testCase.c, testCase.expected1, testCase.expected2, answer1, answer2)
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

func TestSumPow2(t *testing.T) {
	testCases := []struct {
		c        int
		expected int
	}{
		{0, 1},
		{1, 3},
		{2, 7},
		{3, 15},
		{4, 31},
	}

	a = []int{0, 1}

	for _, testCase := range testCases {
		answer := sumPow2(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}

func TestSumFives(t *testing.T) {
	testCases := []struct {
		c        int
		expected int
	}{
		{0, 0},
		{1, 0},
		{2, -5},
		{3, -5 + -10},
		{4, -5 + -10 + -20},
	}

	a = []int{0, 1}

	for _, testCase := range testCases {
		answer := sumFives(testCase.c)
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
	}

	a = []int{0, 1}

	for _, testCase := range testCases {
		answer := sumToPow2(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}
