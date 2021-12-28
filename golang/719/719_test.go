package main

import (
	"testing"
)

func TestSSum(t *testing.T) {
	testCases := []struct {
		max      int
		expected int
	}{
		{90, 81},
		{100, 181},
		{10000, 41333},
	}

	for _, testCase := range testCases {
		answer := sSum(testCase.max)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.max, testCase.expected, answer)
		}
	}
}
