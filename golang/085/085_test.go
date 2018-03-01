package main

import (
	"testing"
)

func TestS(t *testing.T) {
	testCases := []struct {
		width    int
		height   int
		expected int
	}{
		{1, 1, 1},
		{1, 2, 3},
		{1, 3, 6},
		{2, 2, 9},
		{2, 3, 18},
	}

	for _, testCase := range testCases {
		answer := rectangles(testCase.width, testCase.height)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %dx%d expected %d, got %d", testCase.width, testCase.height, testCase.expected, answer)
		}
		answer = rectangles(testCase.height, testCase.width)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %dx%d expected %d, got %d", testCase.height, testCase.width, testCase.expected, answer)
		}
	}
}
