package main

import (
	"testing"
)

func TestIsSquare(t *testing.T) {
	testCases := []struct {
		n        int
		expected bool
	}{
		{1, true},
		{2, false},
		{4, true},
		{6, false},
		{7, false},
		{1000, false},
		{10000, true},
	}

	for _, testCase := range testCases {
		answer := isSquare(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %t, got %t", testCase.n, testCase.expected, answer)
		}
	}
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{2, 3},
		{3, 2},
		{5, 9},
		{6, 5},
		{7, 8},
		{13, 649},
	}

	for _, testCase := range testCases {
		x := solution(testCase.n)
		if x != testCase.expected {
			t.Errorf("ERROR: For %d expected x=%d, got x=%d", testCase.n, testCase.expected, x)
		}
	}
}

func TestMaxSolution(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{3, 3},
		{7, 9},
	}

	for _, testCase := range testCases {
		_, answer := maxSolution(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}
