package main

import (
	"testing"
)

func Test(t *testing.T) {
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
		answer1, answer2, answer3 := findNinRange(testCase.n)
		if answer1 != testCase.expected1 || answer2 != testCase.expected2 || answer3 != testCase.expected3 {
			t.Errorf("ERROR: For %d expected %d %d %f, got %d %d %f", testCase.n, testCase.expected1, testCase.expected2, testCase.expected3, answer1, answer2, answer3)
		}
	}
}
