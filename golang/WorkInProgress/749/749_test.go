package main

import (
	"testing"
)

func TestS(t *testing.T) {
	testCases := []struct {
		d        int
		expected int
	}{
		{1, 0},
		{2, 110},
		{6, 2562701},
		{8, 381784152},
	}

	generatePowers()

	for _, testCase := range testCases {
		answer := s(testCase.d)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.d, testCase.expected, answer)
		}
	}
}
