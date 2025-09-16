package main

import (
	// "testing"
	"testing"
)

func TestEval(t *testing.T) {
	testCases := []struct {
		c        int64
		expected int
	}{
		{0, 0},
		{5, 0},
		{10, 1},
		{25, 2},
		{864210, 1},
		{123689, 2},
		{100001, 3},
		{129291, 3},
	}

	for _, testCase := range testCases {
		answer := eval(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}

func TestFlat(t *testing.T) {
	testCases := []struct {
		c        int
		expected int64
	}{
		{1, 9},
		{2, 9 + 9},
		{3, 9 + 9 + 9},
		{4, 9 + 9 + 9 + 9},
	}

	for _, testCase := range testCases {
		answer := flat(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}

func TestNonBouncy(t *testing.T) {
	testCases := []struct {
		c        int
		expected int64
	}{
		{1, 9},
		{2, 99},
		{6, 12951},
		{10, 277032},
	}

	for _, testCase := range testCases {
		answer := flat(testCase.c) + decreasing(testCase.c) + increasing(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}
