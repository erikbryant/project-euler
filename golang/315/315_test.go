package main

import (
	"testing"
)

func TestTransitionCosts(t *testing.T) {
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			answer1 := TransitionCosts[i][j]
			answer2 := TransitionCosts[j][i]
			if answer1 != answer2 {
				t.Errorf("ERROR: For %d <-> %d expected %d, got %d", i, j, answer1, answer2)
			}
		}
	}

	for i := 0; i <= 9; i++ {
		answer := TransitionCosts[i][i]
		if answer != 0 {
			t.Errorf("ERROR: For %d <-> %d expected %d, got %d", i, i, 0, answer)
		}
	}
}

func TestDigitalRoot(t *testing.T) {
	testCases := []struct {
		n        string
		expected string
	}{
		{"10", "1"},
		{"25", "7"},
		{"137", "11"},
		{"11", "2"},
	}

	for _, testCase := range testCases {
		answer := digitalRoot(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %s expected %s, got %s", testCase.n, testCase.expected, answer)
		}
	}
}

func TestSamDisplayCost(t *testing.T) {
	testCases := []struct {
		n        string
		expected int
	}{
		{"137", 22},
		{"11", 8},
		{"2", 10},
	}

	for _, testCase := range testCases {
		answer := SamDisplayCost(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %s expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}

func TestMaxDisplayCost(t *testing.T) {
	testCases := []struct {
		initial  string
		n        string
		expected int
	}{
		{"", "137", 11},
		{"137", "11", 7},
		{"11", "2", 7},
		{"2", "", 5},
	}

	for _, testCase := range testCases {
		answer := MaxDisplayCost(testCase.initial, testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %s expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}

func TestCostComparison(t *testing.T) {
	testCases := []struct {
		start       int
		end         int
		samExpected int
		maxExpected int
	}{
		{137, 137, 40, 30},
	}

	for _, testCase := range testCases {
		samAnswer, maxAnswer := costComparison(testCase.start, testCase.end)
		if samAnswer != testCase.samExpected || maxAnswer != testCase.maxExpected {
			t.Errorf("ERROR: For %d, %d expected %d/%d, got %d/%d", testCase.start, testCase.end, testCase.samExpected, testCase.maxExpected, samAnswer, maxAnswer)
		}
	}
}
