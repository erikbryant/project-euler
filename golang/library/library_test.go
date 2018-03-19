package library

import (
	"math/big"
	"testing"
)

func TestE(t *testing.T) {
	testCases := []struct {
		n        int
		expected int64
	}{
		{1, 2},
		{2, 1},
		{3, 2},
		{4, 1},
		{5, 1},
		{6, 4},
		{7, 1},
		{8, 1},
		{9, 6},
	}

	for _, testCase := range testCases {
		answer := E(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}

func TestConvergentE(t *testing.T) {
	testCases := []struct {
		n         int
		expectedN int64
		expectedD int64
	}{
		{1, 2, 1},
		{2, 3, 1},
		{3, 8, 3},
		{4, 11, 4},
		{5, 19, 7},
		{6, 87, 32},
		{7, 106, 39},
		{8, 193, 71},
		{9, 1264, 465},
		{10, 1457, 536},
	}

	for _, testCase := range testCases {
		expectedN := big.NewInt(testCase.expectedN)
		expectedD := big.NewInt(testCase.expectedD)
		answerN, answerD := Convergent(testCase.n, E)
		if answerN.Cmp(expectedN) != 0 || answerD.Cmp(expectedD) != 0 {
			t.Errorf("ERROR: For %d expected %d/%d, got %d/%d", testCase.n, testCase.expectedN, testCase.expectedD, answerN, answerD)
		}
	}
}

func TestConvergentSqrt2(t *testing.T) {
	testCases := []struct {
		n         int
		expectedN int64
		expectedD int64
	}{
		{1, 1, 1},
		{2, 3, 2},
		{3, 7, 5},
		{4, 17, 12},
		{5, 41, 29},
		{6, 99, 70},
		{7, 239, 169},
		{8, 577, 408},
	}

	for _, testCase := range testCases {
		expectedN := big.NewInt(testCase.expectedN)
		expectedD := big.NewInt(testCase.expectedD)
		answerN, answerD := Convergent(testCase.n, Sqrt2)
		if answerN.Cmp(expectedN) != 0 || answerD.Cmp(expectedD) != 0 {
			t.Errorf("ERROR: For %d expected %d/%d, got %d/%d", testCase.n, testCase.expectedN, testCase.expectedD, answerN, answerD)
		}
	}
}

func TestFactors(t *testing.T) {
	testCases := []struct {
		n        int
		expected []int
	}{
		{2, []int{}},
		{3, []int{}},
		{4, []int{2}},
		{5, []int{}},
		{6, []int{2, 3}},
		{7, []int{}},
		{8, []int{2}},
		{9, []int{3}},
		{10, []int{2, 5}},
		{11, []int{}},
		{12, []int{2, 3}},
		{210, []int{2, 3, 5, 7}},
		{2310, []int{2, 3, 5, 7, 11}},
	}

	for _, testCase := range testCases {
		answer := Factors(testCase.n)
		if len(answer) != len(testCase.expected) {
			t.Errorf("ERROR: For %d expected len=%d, got len=%d %v", testCase.n, len(testCase.expected), len(answer), answer)
		}
		for i := 0; i < len(answer); i++ {
			if answer[i] != testCase.expected[i] {
				t.Errorf("ERROR: For %d expected %v, got %v", testCase.n, testCase.expected, answer)
			}
		}
	}
}
