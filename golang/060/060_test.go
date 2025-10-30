package main

import (
	"testing"
)

func TestAllCombosPrime(t *testing.T) {
	testCases := []struct {
		primes   []int
		expected bool
	}{
		{[]int{2, 3, 5, 7}, false},
		{[]int{3, 7, 109, 673}, true},
	}

	for _, testCase := range testCases {
		answer := allCombosPrime(testCase.primes)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v expected %t, got %t", testCase.primes, testCase.expected, answer)
		}
	}
}
