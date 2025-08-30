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
