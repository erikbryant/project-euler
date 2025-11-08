package main

import (
	"testing"
)

func TestIsLeft(t *testing.T) {
	testCases := []struct {
		c        int
		expected bool
	}{
		{1, false},
		{2, false},
		{3, true},
		{4, false},
		{5, false},
		{6, true},
		{7, true},
		{8, false},
		{9, false},
		{10, false},
		{11, true},
		{12, true},
		{13, false},
		{14, true},
	}

	for _, testCase := range testCases {
		answer := isLeft(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %t, got %t", testCase.c, testCase.expected, answer)
		}
	}
}

func TestRunSteps(t *testing.T) {
	testCases := []struct {
		c         int
		expectedX int
		expectedY int
	}{
		{0, 0, 0},
		//{1, 0, 1},
		{2, 1, 1},
		{4, 2, 0},
		{500, 18, 16},
	}

	for _, testCase := range testCases {
		answerX, answerY := runSteps(testCase.c)
		if answerX != testCase.expectedX || answerY != testCase.expectedY {
			t.Errorf("ERROR: For %d expected (%d, %d), got (%d, %d)", testCase.c, testCase.expectedX, testCase.expectedY, answerX, answerY)
		}
	}
}
