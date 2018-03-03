package main

import (
	"testing"
)

func TestReduce(t *testing.T) {
	testCases := []struct {
		n  int
		d  int
		rN int
		rD int
	}{
		{1, 2, 1, 2},
		{1, 3, 1, 3},
		{2, 3, 2, 3},
		{1, 4, 1, 4},
		{2, 4, 1, 2},
		{3, 4, 3, 4},
		{3, 4, 3, 4},
		{3, 9, 1, 3},
	}

	for _, testCase := range testCases {
		rN, rD := reduce(testCase.n, testCase.d)
		if rN != testCase.rN || rD != testCase.rD {
			t.Errorf("ERROR: For %d/%d expected %d/%d, got %d/%d", testCase.n, testCase.d, testCase.rN, testCase.rD, rN, rD)
		}
	}
}

func TestRpf(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 0},
		{5, 1},
		{6, 1},
		{7, 2},
		{8, 3},
		{9, 4},
	}

	for _, testCase := range testCases {
		answer := rpf(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}
