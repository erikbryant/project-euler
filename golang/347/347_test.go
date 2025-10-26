package main

import (
	"testing"
)

func TestDoesItFActor(t *testing.T) {
	testCases := []struct {
		p        int
		q        int
		N        int
		expected bool
	}{
		{2, 3, 96, true},
		{3, 5, 75, true},
		{3, 5, 90, false},
	}

	for _, testCase := range testCases {
		answer := soloFactors(testCase.p, testCase.q, testCase.N)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d, %d, %d expected %t, got %t", testCase.p, testCase.q, testCase.N, testCase.expected, answer)
		}
	}
}

func TestM(t *testing.T) {
	testCases := []struct {
		p        int
		q        int
		N        int
		expected int
	}{
		{2, 3, 100, 96},
		{3, 5, 100, 75},
		{2, 73, 100, 0},
		{5, 73, 100, 0},
	}

	for _, testCase := range testCases {
		answer := M(testCase.p, testCase.q, testCase.N)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d, %d, %d expected %d, got %d", testCase.p, testCase.q, testCase.N, testCase.expected, answer)
		}
	}
}

func TestS(t *testing.T) {
	testCases := []struct {
		N        int
		expected int
	}{
		{5, 0},
		{6, 6},
		{10, 16},
		{100, 2262},
		{1000, 193408},
	}

	for _, testCase := range testCases {
		answer := S(testCase.N)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.N, testCase.expected, answer)
		}
	}
}
