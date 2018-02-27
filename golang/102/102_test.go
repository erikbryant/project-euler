package main

import (
	"testing"
)

func TestContainsOrigin(t *testing.T) {
	testCases := []struct {
		triangle []vertex
		expected bool
	}{
		{[]vertex{{1, 1}, {-1, 1}, {0, -4}}, true},
		{[]vertex{{0, 0}, {1, 1}, {2, 2}}, false},
		{[]vertex{{1, 1}, {1, -1}, {-1, 0}}, true},
		{[]vertex{{-340, 495}, {-153, -910}, {835, -947}}, true},
		{[]vertex{{-175, 41}, {-421, -714}, {574, -645}}, false},
	}

	for _, testCase := range testCases {
		contains := containsOrigin(testCase.triangle)
		if contains != testCase.expected {
			t.Errorf("ERROR: For %v expected %t, got %t", testCase.triangle, testCase.expected, contains)
		}
	}
}
