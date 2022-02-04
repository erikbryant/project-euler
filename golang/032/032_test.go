package main

import (
	"testing"
)

func TestMulDigits(t *testing.T) {
	testCases := []struct {
		digits   []int
		m2       int
		product  int
		expected bool
	}{
		{[]int{3, 9, 1, 8, 6, 7, 2, 5, 4}, 2, 5, true},
	}

	for _, testCase := range testCases {
		answer := mulDigits(testCase.digits, testCase.m2, testCase.product)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v, m1=%d, product=%d expected %t, got %t", testCase.digits, testCase.m2, testCase.product, testCase.expected, answer)
		}
	}
}
