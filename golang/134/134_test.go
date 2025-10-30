package main

import (
	"testing"
)

func TestLCM(t *testing.T) {
	testCases := []struct {
		a        int
		b        int
		expected int
	}{
		{5, 7, 35},
		{7, 11, 77},
		{11, 13, 611},
		{13, 17, 1513},
		{19, 23, 1219},
	}

	for _, testCase := range testCases {
		answer := LCM(testCase.a, testCase.b)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d %d expected %d, got %d", testCase.a, testCase.b, testCase.expected, answer)
		}
	}
}

func TestLooper(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{5, 35},
		{7, 112},
		{11, 723},
		{13, 2236},
		{17, 3053},
		{19, 4272},
	}

	for _, testCase := range testCases {
		answer := looper(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}
