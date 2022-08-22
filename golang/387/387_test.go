package main

import (
	"github.com/erikbryant/project-euler/golang/util"
	"testing"
)

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
		answer := rightTruncatableHarshad(testCase.c, util.DigitSum(testCase.c))
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
		answer := strong(testCase.c, util.DigitSum(testCase.c))
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
