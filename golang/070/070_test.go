package main

import (
	"testing"
)

func TestIsPermutation(t *testing.T) {
	testCases := []struct {
		n1       int
		n2       int
		expected bool
	}{
		{2, 1, false},
		{2, 2, true},
		{212, 122, true},
		{21222, 122, false},
		{212, 122222, false},
		{2000, 200, false},
		{87109, 79180, true},
	}

	for _, testCase := range testCases {
		answer := isPermutation(testCase.n1, testCase.n2)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d, %d expected %t, got %t", testCase.n1, testCase.n2, testCase.expected, answer)
		}
	}
}

func TestLooper(t *testing.T) {
	testCases := []struct {
		n         int
		expected1 int
		expected2 int
		expected3 float64
	}{
		{1000, 291, 192, 1.515625},
		{10 * 1000, 4435, 3544, float64(4435) / float64(3544)},
		{100 * 1000, 75841, 75184, 1.0087385613960418},
	}

	for _, testCase := range testCases {
		answer1, answer2, answer3 := looper(testCase.n)
		if answer1 != testCase.expected1 || answer2 != testCase.expected2 || answer3 != testCase.expected3 {
			t.Errorf("ERROR: For %d expected %d %d %f, got %d %d %f", testCase.n, testCase.expected1, testCase.expected2, testCase.expected3, answer1, answer2, answer3)
		}
	}
}
