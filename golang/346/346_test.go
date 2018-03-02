package main

import (
	"testing"
)

func TestBaseB(t *testing.T) {
	testCases := []struct {
		base     int
		max      int
		expected []int
	}{
		{2, 2, []int{1}},
		{6, 100, []int{1, 7, 43}},
		{10, 20000, []int{1, 11, 111, 1111, 11111}},
	}

	for _, testCase := range testCases {
		answer := baseB(testCase.base, testCase.max)
		if len(answer) != len(testCase.expected) {
			t.Errorf("ERROR: For baseb(%d, %d) expected %v, got %v", testCase.base, testCase.max, testCase.expected, answer)
		}
		for i := range testCase.expected {
			if answer[i] != testCase.expected[i] {
				t.Errorf("ERROR: For baseb(%d, %d) expected %v, got %v", testCase.base, testCase.max, testCase.expected, answer)
				break
			}
		}
	}
}

func TestIsRepunit(t *testing.T) {
	testCases := []struct {
		base     int
		r        int
		expected bool
	}{
		{5, 1, true},
		{6, 7, true},
		{6, 5, false},
		{6, 8, false},
	}

	for _, testCase := range testCases {
		answer := isRepunit(testCase.base, testCase.r)
		if answer != testCase.expected {
			t.Errorf("ERROR: For isRepunit(%d, %d) expected %t, got %t", testCase.base, testCase.r, testCase.expected, answer)
		}
	}
}

func TestSumRepunits(t *testing.T) {
	testCases := []struct {
		max      int
		expected int
	}{
		{50, 171},
		{1000, 15864},
	}

	for _, testCase := range testCases {
		answer := sumRepunits(testCase.max)
		if answer != testCase.expected {
			t.Errorf("ERROR: For sumRepunits(%d) expected %d, got %d", testCase.max, testCase.expected, answer)
		}
	}
}
