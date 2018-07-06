package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{36, 63},
		{409, 904},
		{123, 321},
		{10002, 20001},
	}

	for _, testCase := range testCases {
		answer := reverse(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}

func TestOdd(t *testing.T) {
	testCases := []struct {
		n        int
		expected bool
	}{
		{99, true},
		{1313, true},
		{36, false},
		{409, false},
		{123, false},
		{10002, false},
		{1, true},
		{2, false},
	}

	for _, testCase := range testCases {
		answer := odd(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %t, got %t", testCase.n, testCase.expected, answer)
		}
	}
}

func TestReversible(t *testing.T) {
	testCases := []struct {
		n        int
		expected bool
	}{
		{36, true},
		{63, true},
		{409, true},
		{904, true},
	}

	for _, testCase := range testCases {
		answer := reversible(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %t, got %t", testCase.n, testCase.expected, answer)
		}
	}
}

func TestCountReversible(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{10, 0},
		{100, 20},
		{1000, 120},
		{1000 * 10, 720},
		{1000 * 100, 720},
		{1000 * 1000, 18720},
	}

	for _, testCase := range testCases {
		answer := countReversible(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}
