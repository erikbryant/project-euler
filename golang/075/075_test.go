package main

import (
	"testing"
)

func TestTriples(t *testing.T) {
	testCases := []struct {
		c        int
		expected int
	}{
		{9000, 1008},
		{10000, 1120},
		{15000, 1663},
	}

	for _, testCase := range testCases {
		answer := triples(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}
