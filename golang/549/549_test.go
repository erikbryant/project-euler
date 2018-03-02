package main

import (
	"testing"
)

func TestS(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{10, 5},
		{25, 10},
	}

	for _, testCase := range testCases {
		answer := s(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}

func TestSums(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{100, 2012},
		{1000, 136817},
		// {1000 * 10, 10125843},
		// {1000 * 100, 793183093},
		// {1000 * 104, 854396111},
		// {1000 * 200, 2975450133},
		// {1000 * 300, 6452498638},
	}

	for _, testCase := range testCases {
		answer := sumS(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}
