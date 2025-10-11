package main

import (
	"math"
	"testing"

	"github.com/erikbryant/util-golang/algebra"
)

func TestAssumption(t *testing.T) {
	// The assumption is that for all n where n is a Hamming number
	// times any number of primes with exponents <= 1 where prime-1
	// is a Hamming number, the totient(n) is a Hamming number

	testCases := []struct {
		c int
	}{
		{2 * 7},
		{2 * 7 * 11},
		{3 * 13},
		{3 * 17 * 151},
		{5 * 19},
		{5 * 31 * 101 * 3889},
	}

	for _, testCase := range testCases {
		answer := algebra.Hamming(algebra.Totient(testCase.c))
		if answer != true {
			t.Errorf("ERROR: For %d expected %t, got %t", testCase.c, true, answer)
		}
	}
}

func TestS(t *testing.T) {
	testCases := []struct {
		c        int
		expected int
	}{
		{100, 3728},
		{1000, 203813},
		{10000, 9586559},
	}

	mod := int(math.Pow(2, 32))

	for _, testCase := range testCases {
		_, answer := S(testCase.c, mod)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}
