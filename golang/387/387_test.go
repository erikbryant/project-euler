package main

import (
	"testing"
)

func TestDigitSum(t *testing.T) {
	testCases := []struct {
		c        int
		expected int
	}{
		{0, 0},
		{5, 5},
		{10, 1},
		{25, 7},
		{100000, 1},
		{100001, 2},
	}

	for _, testCase := range testCases {
		answer := digitSum(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}

func TestHarshad(t *testing.T) {
	testCases := []struct {
		c        int
		expected bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{5, true},
		{7, true},
		{201, true},
		{2011, false},
		{100000, true},
		{100001, false},
	}

	for _, testCase := range testCases {
		answer := harshad(testCase.c, digitSum(testCase.c))
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %t, got %t", testCase.c, testCase.expected, answer)
		}
	}
}

func TestRightTruncatableHarshad(t *testing.T) {
	testCases := []struct {
		c        int
		expected bool
	}{
		{10, true},
		{20, true},
		{201, true},
		{2011, false},
		{100000, true},
		{100001, false},
	}

	for _, testCase := range testCases {
		answer := rightTruncatableHarshad(testCase.c, digitSum(testCase.c))
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %t, got %t", testCase.c, testCase.expected, answer)
		}
	}
}

func TestStrong(t *testing.T) {
	testCases := []struct {
		c        int
		expected bool
	}{
		{100, false},
		{201, true},
		{2011, false},
	}

	for _, testCase := range testCases {
		answer := strong(testCase.c, digitSum(testCase.c))
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %t, got %t", testCase.c, testCase.expected, answer)
		}
	}
}

func TestSumSRTHP(t *testing.T) {
	testCases := []struct {
		max      int
		expected int
	}{
		{181, 181},
		{182, 181},
		{10000, 90619},
		{10001, 90619},
	}

	for _, testCase := range testCases {
		c := make(chan int, 10)
		go findRTH(testCase.max, c)
		answer := sumSRTHP(testCase.max, c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.max, testCase.expected, answer)
		}
	}
}
