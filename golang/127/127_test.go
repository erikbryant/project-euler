package main

import (
	// "testing"
	"testing"
)

func TestCountAbcHits(t *testing.T) {
	testCases := []struct {
		c             int
		expectedCount int
		expectedSum   int
	}{
		{c: 1000, expectedCount: 31, expectedSum: 12523},
		//{c: 30000, expectedCount: 234, expectedSum: 2540437},
	}

	for _, testCase := range testCases {
		answerCount, answerSum := countAbcHits(testCase.c)
		if answerCount != testCase.expectedCount || answerSum != testCase.expectedSum {
			t.Errorf("ERROR: For %d expected %d / %d, got %d / %d", testCase.c, testCase.expectedCount, testCase.expectedSum, answerCount, answerSum)
		}
	}
}

// func TestXYZ(t *testing.T) {
// 	testCases := []struct {
// 		c        int
// 		expected int
// 	}{
// 		{0, 0},
// 		{5, 5},
// 		{10, 1},
// 		{25, 7},
// 		{100000, 1},
// 		{100001, 2},
// 	}
//
// 	for _, testCase := range testCases {
// 		answer := XYZ(testCase.c)
// 		if answer != testCase.expected {
// 			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
// 		}
// 	}
// }
