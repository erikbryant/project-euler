package main

import (
	"testing"
)

func TestDivisorSums(t *testing.T) {
	testCases := []struct {
		c        int
		expected int
	}{
		{0, 0},
		{1, 0},
		{2, 1},
		{4, 3},
		{28, 28},
		{220, 284},
		{284, 220},
	}

	for _, testCase := range testCases {
		answer := divisorSums[testCase.c]
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}

func TestIn(t *testing.T) {
	testCases := []struct {
		c        int
		s        []int
		expected bool
	}{
		{0, []int{}, false},
		{0, []int{0}, true},
		{1, []int{9, 4, 10, 12}, false},
		{2, []int{2, 3, 4, 5}, true},
		{4, []int{1, 8, 4, 3, 6}, true},
		{6, []int{4, 8, 9, 3, 6}, true},
	}

	for _, testCase := range testCases {
		answer := in(testCase.s, testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d  %v expected %t, got %t", testCase.s, testCase.c, testCase.expected, answer)
		}
	}
}

func TestChainer(t *testing.T) {
	testCases := []struct {
		c         int
		expected1 int
		expected2 int
		expected3 bool
	}{
		{0, 0, 0, false},
		{1, 0, 0, false},
		{6, 6, 1, true},
		{25, 6, 1, true},
		{28, 28, 1, true},
		{220, 220, 2, true},
	}

	for _, testCase := range testCases {
		answer1, answer2, answer3 := chainer(testCase.c)
		if answer1 != testCase.expected1 || answer2 != testCase.expected2 || answer3 != testCase.expected3 {
			t.Errorf("ERROR: For %d expected %d %d %t, got %d %d %t", testCase.c, testCase.expected1, testCase.expected2, testCase.expected3, answer1, answer2, answer3)
		}
	}
}

func TestLooper(t *testing.T) {
	testCases := []struct {
		c         int
		expected1 int
		expected2 int
		expected3 int
	}{
		{0, Max, 0, 0},
		{1, Max, 0, 0},
		{25, 6, 1, 6},
		{28, 6, 1, 6},
		{220, 220, 2, 220},
	}

	for _, testCase := range testCases {
		answer1, answer2, answer3 := looper(testCase.c)
		if answer1 != testCase.expected1 || answer2 != testCase.expected2 || answer3 != testCase.expected3 {
			t.Errorf("ERROR: For %d expected %d %d %d, got %d %d %d", testCase.c, testCase.expected1, testCase.expected2, testCase.expected3, answer1, answer2, answer3)
		}
	}
}
