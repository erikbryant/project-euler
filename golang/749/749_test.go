package main

import (
	"testing"
)

func TestSolvable(t *testing.T) {
	testCases := []struct {
		digits   []int
		expected bool
	}{
		{[]int{0}, false},
		{[]int{1}, false},
		{[]int{9}, true},
		{[]int{3, 5}, true},
		{[]int{1, 1}, false},
		{[]int{1, 1, 0}, false},
		{[]int{0, 1, 1, 0}, false},
		{[]int{0, 1, 1, 0, 3}, true},
		{[]int{4, 0, 1, 1, 0}, true},
	}

	generatePowers()

	for _, testCase := range testCases {
		answer := solvable(testCase.digits)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v expected %t, got %t", testCase.digits, testCase.expected, answer)
		}
	}
}

func TestCanMakeSumFromDigits(t *testing.T) {
	testCases := []struct {
		digits   []int
		sum      int
		expected bool
	}{
		{[]int{0}, 0, true},
		{[]int{1}, 1, true},
		{[]int{9}, 9, true},
		{[]int{3, 5}, 34, false},
		{[]int{3, 5}, 36, false},
		{[]int{3, 5}, 35, true},
	}

	generatePowers()

	for _, testCase := range testCases {
		answer := canMakeSumFromDigits(testCase.digits, testCase.sum)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v, %d expected %t, got %t", testCase.digits, testCase.sum, testCase.expected, answer)
		}
	}
}

func TestPowerSums(t *testing.T) {
	testCases := []struct {
		digits   []int
		expected int
	}{
		{[]int{3, 5}, 35},
		{[]int{5, 3}, 35},
		{[]int{1, 4, 5, 6, 7, 9}, 715469},
		{[]int{2, 2, 6, 7, 8, 8}, 688722},
		{[]int{2, 3, 4, 6, 6, 9}, 629643},
		{[]int{2, 5, 5, 7, 7, 8}, 528757},
		{[]int{1, 1}, 0},
		{[]int{1, 1, 0}, 0},
		{[]int{0, 1, 1, 0}, 0},
		{[]int{0, 1, 1, 0, 3}, 0},
		{[]int{4, 0, 1, 1, 0}, 0},
	}

	generatePowers()

	for _, testCase := range testCases {
		answer := powerSums(testCase.digits)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v expected %d, got %d", testCase.digits, testCase.expected, answer)
		}
	}
}
