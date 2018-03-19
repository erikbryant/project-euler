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

func TestReduce(t *testing.T) {
	testCases := []struct {
		n         int
		d         int
		expectedN int
		expectedD int
	}{
		{1, 2, 1, 2},
		{2, 4, 1, 2},
		{3, 9, 1, 3},
		{4, 16, 1, 4},
		{5, 100, 1, 20},
		{4, 6, 2, 3},
	}

	for _, testCase := range testCases {
		answer := reduce([]int{testCase.n, testCase.d})
		if answer[0] != testCase.expectedN || answer[1] != testCase.expectedD {
			t.Errorf("ERROR: For %d/%d expected %d/%d, got %d/%d", testCase.n, testCase.d, testCase.expectedN, testCase.expectedD, answer[0], answer[1])
		}
	}
}
