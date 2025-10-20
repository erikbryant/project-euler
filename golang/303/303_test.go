package main

import "testing"

func TestMultiple(t *testing.T) {
	testCases := []struct {
		c        uint
		expected uint
	}{
		{1, 1 / 1},
		{2, 2 / 2},
		{3, 12 / 3},
		{5, 10 / 5},
		{7, 21 / 7},
		{9, 1358},
		{10, 10 / 10},
		{25, 100 / 25},
		{42, 210 / 42},
		{89, 1121222 / 89},
		{99, 11335578},
		{100000, 100000 / 100000},
		{100001, 100001 / 100001},
	}

	for _, testCase := range testCases {
		answer := multiple(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}

func TestComputeAll(t *testing.T) {
	testCases := []struct {
		upper    uint
		expected uint
	}{
		{1, 1},
		{100, 11363107},
	}

	for _, testCase := range testCases {
		answer := computeAll(testCase.upper)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.upper, testCase.expected, answer)
		}
	}
}

// func TestXYZ(t *testing.T) {
// 	testCases := []struct {
// 		upper        int
// 		expected int
// 	}{
// 		{0, 0},
// 		{5, 5},
// 		{10, 1},
// 		{25, 7},
// 		{100000, 1},
// 		{100001, 2},
// 	}
//
// 	for _, testCase := range testCases {
// 		answer := XYZ(testCase.upper)
// 		if answer != testCase.expected {
// 			t.Errorf("ERROR: For %d expected %d, got %d", testCase.upper, testCase.expected, answer)
// 		}
// 	}
// }
