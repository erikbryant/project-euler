package main

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
		answer := e(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}

func TestConvergent(t *testing.T) {
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
		answerN, answerD := convergent(testCase.n)
		if answerN.Cmp(expectedN) != 0 || answerD.Cmp(expectedD) != 0 {
			t.Errorf("ERROR: For %d expected %d/%d, got %d/%d", testCase.n, testCase.expectedN, testCase.expectedD, answerN, answerD)
		}
	}
}
