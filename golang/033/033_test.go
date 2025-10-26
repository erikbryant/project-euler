package main

import (
	"testing"
)

func TestCancels(t *testing.T) {
	testCases := []struct {
		fraction []int
		expected bool
	}{
		{[]int{49, 98}, true},
		{[]int{30, 50}, false}, // trivial cases do not count
		{[]int{25, 44}, false},
	}

	for _, testCase := range testCases {
		answer := cancels(testCase.fraction)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v expected %t, got %t", testCase.fraction, testCase.expected, answer)
		}
	}
}
