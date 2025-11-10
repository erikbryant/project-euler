package main

import (
	"testing"
)

func TestFactor1(t *testing.T) {
	testCases := []struct {
		c        int
		expected int
	}{
		{0, 0},
		{23, 0},
		{127, 0},
		{128, 1},
		{129, 1},
		{1000, 1},
		{2186, 1},
		{2187, 2},
		{2188, 2},
		{10000, 2},
		{100000, 3},
		{1000000, 4},
		{10000000, 4},
		{1000 * 1000 * 1000 * 1000, 15},
	}

	for _, testCase := range testCases {
		answer := factor1(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}

func TestFactor2(t *testing.T) {
	testCases := []struct {
		c        int
		expected int
	}{
		{0, 0},
		{23, 0},
		{24, 1},
		{25, 1},
		{39, 1},
		{40, 2},
		{41, 2},
		{53, 2},
		{54, 3},
		{55, 3},
		{127, 6},
		{2186, 87},
		{1000, 44},
		{10000, 312},
		{100000, 2259},
		{1000000, 17459},
		{100 * 1000 * 1000, 1191658},
		//{1000 * 1000 * 1000 * 1000, 15},
	}

	for _, testCase := range testCases {
		answer := factor2(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}

func TestFactor3(t *testing.T) {
	testCases := []struct {
		c        int
		expected int
	}{
		{0, 0},
		{29, 0},
		{30, 1},
		{31, 1},
		{41, 1},
		{42, 2},
		{43, 2},
		{127, 9},
		{2186, 335},
		{1000, 135},
		{10000, 1800},
		{100000, 19919},
		{1000000, 206964},
		{100 * 1000 * 1000, 20710806},
		//{1000 * 1000 * 1000 * 1000, 15},
	}

	for _, testCase := range testCases {
		answer := factor3(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}

func TestDivisors8(t *testing.T) {
	testCases := []struct {
		c        int
		expected int
	}{
		{0, 0},
		{23, 0},
		{24, 1},
		{25, 1},
		{100, 10},
		{1000, 180},
		{1000 * 1000, 224427},
		{10 * 1000 * 1000, 2228418},
		{100 * 1000 * 1000, 21902470},
	}

	for _, testCase := range testCases {
		answer := divisors8(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}
