package main

import (
	"testing"
)

func TestLongDivision(t *testing.T) {
	testCases := []struct {
		divisor  int
		answer   string
		expected int
	}{
		{2, "5", 0},
		{3, "3", 1},
		{4, "25", 0},
		{5, "2", 0},
		{6, "6", 1},
		{7, "142857", 6},
		{8, "125", 0},
		{9, "1", 1},
		{10, "1", 0},
		{11, "09", 2},
		{12, "3", 1},
		{13, "076923", 6},
		{14, "714285", 6},
		{15, "6", 1},
		{16, "0625", 0},
		{17, "0588235294117647", 16},
		{18, "5", 1},
		{19, "052631578947368421", 18},
		{20, "05", 0},
		{21, "047619", 6},
		{22, "45", 2},
		{23, "0434782608695652173913", 22},
	}

	for _, testCase := range testCases {
		answer, length := longDivision(1, testCase.divisor)
		if answer != testCase.answer {
			t.Errorf("ERROR: For %d expected '%s', got '%s'", testCase.divisor, testCase.answer, answer)
		}
		if length != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.divisor, testCase.expected, length)
		}
	}
}
