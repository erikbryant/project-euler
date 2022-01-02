package main

import (
	"testing"
)

func TestFindTriangle(t *testing.T) {
	makeTriangles()

	testCases := []struct {
		t        int
		min      int
		max      int
		expected int
	}{
		{0, 0, len(triangles) - 1, 0},
		{1, 0, 99, 1},
		{2, 0, 99, 1},
		{3, 0, 99, 2},
		{4, 0, 99, 3},
		{5, 0, 4, 3},
		{6, 3, 99, 3},
		{7, 4, 99, 4},
		{8, 0, 99, 4},
		{9, 3, 99, 4},
		{15, 0, len(triangles) - 1, 5},
		{15, 33, 33, 33},
	}

	for _, testCase := range testCases {
		answer := findTriangle(testCase.t, testCase.min, testCase.max)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.t, testCase.expected, answer)
		}
	}
}

func TestTriangular(t *testing.T) {
	testCases := []struct {
		c        int
		expected bool
	}{
		{0, true},
		{1, true},
		{9, false},
		{10, true},
		{1000, false},
		{17526, false},
		{1000 * 1000, false},
	}

	for _, testCase := range testCases {
		answer := triangular(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %t, got %t", testCase.c, testCase.expected, answer)
		}
	}
}

func TestTSumCount(t *testing.T) {
	makeTriangles()

	testCases := []struct {
		c        int
		expected int
	}{
		{9, 7},
		// {10, 9},
		// {1000, 78},
		// {17526, 312},
		// {1000 * 1000, 2106},
	}

	for _, testCase := range testCases {
		answer := tSumCount(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}
