package main

import (
	"math"
	"testing"
)

func TestSummable(t *testing.T) {
	testCases := []struct {
		c        int
		expected bool
	}{
		{9, false},
		{81, true},
		{2025, true},
		{3025, true},
		{9801, false},
	}

	for _, testCase := range testCases {
		root := int(math.Sqrt(float64(testCase.c)))
		answer := summable(root, testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %t, got %t", testCase.c, testCase.expected, answer)
		}
	}
}
