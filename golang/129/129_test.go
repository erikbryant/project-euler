package main

import "testing"

func TestA(t *testing.T) {
	testCases := []struct {
		c        int
		expected int
	}{
		{7, 6},
		{41, 5},
		{17, 16},
	}

	for _, testCase := range testCases {
		answer := A(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}
