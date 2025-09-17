package main

import "testing"

func TestModSub(t *testing.T) {
	testCases := []struct {
		c        int
		d        int
		m        int
		expected int
	}{
		{5, 1, 1000000007, 4},
		{5, 6, 1000000007, 1000000006},
	}

	for _, testCase := range testCases {
		answer := modSub(testCase.c, testCase.d, testCase.m)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d - %d mod %d expected %d, got %d", testCase.c, testCase.d, testCase.m, testCase.expected, answer)
		}
	}
}
