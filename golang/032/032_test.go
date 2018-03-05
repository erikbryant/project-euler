package main

import (
	"testing"
)

func TestDigits(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{3, 6},
		{4, 24},
		{9, 362880},
	}

	for _, testCase := range testCases {
		c := make(chan []int, 1000)
		go makeDigits(testCase.n, c)
		answer := 0
		for {
			_, ok := <-c
			if !ok {
				break
			}
			answer++
		}
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}

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
