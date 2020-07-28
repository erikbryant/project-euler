package main

import (
	"testing"
)

func TestPrefixMatch(t *testing.T) {
	testCases := []struct {
		prefix int64
		target int64
		answer bool
	}{
		{1, 1, true},
		{2, 1, false},
		{2, 3, false},
		{12, 12, true},
		{12, 121, true},
		{99, 989, false},
	}

	for _, testCase := range testCases {
		answer := prefixMatch(testCase.prefix, testCase.target)
		if answer != testCase.answer {
			t.Errorf("For %d, %d expected %t, got %t", testCase.prefix, testCase.target, testCase.answer, answer)
		}
	}
}

func TestP(t *testing.T) {
	testCases := []struct {
		L      int64
		n      int64
		answer int64
	}{
		{1, 1, 0},
		{2, 1, 1},
		{12, 1, 7},
		{12, 2, 80},
		{123, 45, 12710},
	}

	for _, testCase := range testCases {
		answer := p(testCase.L, testCase.n)
		if answer != testCase.answer {
			t.Errorf("For L=%d n=%d expected %d, got %d", testCase.L, testCase.n, testCase.answer, answer)
		}
	}
}
