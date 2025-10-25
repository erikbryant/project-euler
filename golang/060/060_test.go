package main

import (
	"testing"

	"github.com/erikbryant/util-golang/primes"
)

func TestAllCombosPrime(t *testing.T) {
	testCases := []struct {
		primeIndexes []int
		primes       []int
		expected     bool
	}{
		{[]int{0, 1, 2, 3}, []int{2, 3, 5, 7}, false},
		{[]int{1, 3, 28, 121}, []int{3, 7, 109, 673}, true},
	}

	// allCombosPrime() takes the index of the primes, not the
	// primes themselves. So, we have to pass it the indexes.
	// Make sure we got the indexes right. :-)
	for _, testCase := range testCases {
		for i := 0; i < len(testCase.primeIndexes); i++ {
			if testCase.primes[i] != primes.Primes[testCase.primeIndexes[i]] {
				t.Errorf("ERROR: For %v, %v expected %d, got %d", testCase.primeIndexes, testCase.primes, testCase.primes[i], primes.Primes[testCase.primeIndexes[i]])
			}
		}
	}

	for _, testCase := range testCases {
		answer := allCombosPrime(testCase.primeIndexes)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v expected %t, got %t", testCase.primes, testCase.expected, answer)
		}
	}
}
