package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	testCases := []struct {
		s1       int
		s2       int
		expected int
	}{
		{5, 4, 14},
		{5, 6, 16},
		{6, 5, 17},
		{6, 7, 19},
	}

	for _, testCase := range testCases {
		answer := perimeter(testCase.s1, testCase.s2)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d %d expected %d, got %d", testCase.s1, testCase.s2, testCase.expected, answer)
		}
	}
}

func TestAreaIntHero(t *testing.T) {
	testCases := []struct {
		s1        int
		s2        int
		expected1 int
		expected2 bool
	}{
		{5, 4, 0, false},
		{5, 6, 12, true},
		{6, 5, 0, false},
		{6, 7, 0, false},
	}

	for _, testCase := range testCases {
		answer, ok := areaIntHero(testCase.s1, testCase.s2)
		if answer != testCase.expected1 {
			t.Errorf("ERROR: For %d %d expected %d, got %d", testCase.s1, testCase.s2, testCase.expected1, answer)
		}
		if ok != testCase.expected2 {
			t.Errorf("ERROR: For %d %d expected %t, got %t", testCase.s1, testCase.s2, testCase.expected2, ok)
		}
	}
}
