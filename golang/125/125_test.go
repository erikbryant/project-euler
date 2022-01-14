package main

import (
	"testing"
)

func TestLooper(t *testing.T) {
	testCases := []struct {
		c        int
		expected int
	}{
		{1000, 4164},
	}

	for _, testCase := range testCases {
		answer := looper(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}
