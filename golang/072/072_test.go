package main

import (
	"testing"
)

func TestCountNumerators(t *testing.T) {
	testCases := []struct {
		d        int
		expected int
	}{
		{2, 1},
		{3, 2},
		{4, 2},
		{5, 4},
		{6, 2},
		{7, 6},
		{8, 4},
	}

	for _, testCase := range testCases {
		answer := countNumerators(testCase.d)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.d, testCase.expected, answer)
		}
	}
}

func TestLooper(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{8, 21},
		{1000, 304191},
		{10000, 30397485},
	}

	for _, testCase := range testCases {
		answer := looper(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}
