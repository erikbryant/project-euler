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
		answer := harshad(testCase.c)
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
		answer := rightTruncatableHarshad(testCase.c)
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
		answer := strong(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %t, got %t", testCase.c, testCase.expected, answer)
		}
	}
}

func TestStrongRightTruncatableHarshad(t *testing.T) {
	testCases := []struct {
		c        int
		expected bool
	}{
		{10, false},
		{21, true},
		{201, true},
		{2011, false},
	}

	for _, testCase := range testCases {
		answer := strongRightTruncatableHarshad(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %t, got %t", testCase.c, testCase.expected, answer)
		}
	}
}

func TestStrongRightTruncatableHarshadPrime(t *testing.T) {
	testCases := []struct {
		c        int
		expected bool
	}{
		{10, false},
		{11, false},
		{201, false},
		{2011, true},
		{100000, false},
		{100001, false},
	}

	for _, testCase := range testCases {
		answer := strongRightTruncatableHarshadPrime(testCase.c)
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
		answer := sumSRTHP(testCase.max)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.max, testCase.expected, answer)
		}
	}
}

// 211
// 271
// 277
// 421
// 457
// 631
// 2011
// 2017
// 2099
// 2473
// 2477
// 4021
// 4027
// 4073
// 4079
// 4231
// 4813
// 4817
// 6037
// 8011
// 8017
// 8039
// 8461
// 8467
// 20071
// 20431
// 40867
// 48091
// 84061
// 84067
// 400237
// 400277
// 4008271
// 4860013
// 40000021
// 80402071
// 200400073
// 200400077
// 240840013
// 400002073
// 480006031

// 2000000011
// 2400000073
// 2408400811
//
// 4000008697
// 4008200071
// 4020800071
//
// 8004000619
// 8004600031
//
// 20000000431
// 20040000031
//
// 40000000861
// 40020000037
// 40208040091
