package util

import (
  "testing"
)

func TestIsPalindromeString(t *testing.T) {
	testCases := []struct {
		c        string
		expected bool
	}{
		{"", true},
    {"w", true},
		{"aba", true},
		{"aab", false},
		{"-22-", true},
	}

	for _, testCase := range testCases {
		answer := IsPalindromeString(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v expected %t, got %t", testCase.c, testCase.expected, answer)
		}
	}
}

func TestIsPalindromeInt(t *testing.T) {
	testCases := []struct {
		c        []int
		expected bool
	}{
		{[]int{}, true},
		{[]int{1}, true},
		{[]int{1, 2}, false},
		{[]int{1, 2, 1}, true},
		{[]int{6, 4, 4, 6}, true},
	}

	for _, testCase := range testCases {
		answer := IsPalindromeInt(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v expected %t, got %t", testCase.c, testCase.expected, answer)
		}
	}
}
