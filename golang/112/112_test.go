package main

import (
	"testing"
)

func TestIsBouncy(t *testing.T) {
	testCases := []struct {
		n        int
		expected bool
	}{
		{0, false},
		{1, false},
		{25, false},
		{100, false},
		{101, true},
		{235, false},
		{538, true},
		{1000, false},
		{1001, true},
		{1234, false},
		{155349, true},
	}

	for _, testCase := range testCases {
		answer := bouncy(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %t, got %t", testCase.n, testCase.expected, answer)
		}
	}
}

func TestCountBouncy(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{999, 525},
	}

	for _, testCase := range testCases {
		answer := countBouncy(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}
