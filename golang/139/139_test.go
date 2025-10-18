package main

import (
	"testing"
)

func TestCanTile(t *testing.T) {
	testCases := []struct {
		a        int
		b        int
		c        int
		expected bool
	}{
		{3, 4, 5, true},
		{5, 12, 13, false},
	}

	for _, testCase := range testCases {
		answer := CanTile(testCase.a, testCase.b, testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d, %d, %d expected %t, got %t", testCase.a, testCase.b, testCase.c, testCase.expected, answer)
		}
	}
}
