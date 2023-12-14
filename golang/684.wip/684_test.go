package main

import "testing"

func TestLittleS(t *testing.T) {
	testCases := []struct {
		c        uint64
		expected uint64
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
		{6, 6},
		{7, 7},
		{8, 8},
		{9, 9},
		{10, 19},
		{11, 29},
		{20, 299},
		{55, 1999999},
		{83, 2999999999},
		{100, 199999999999},
	}

	for _, testCase := range testCases {
		answer := s(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}

func TestBigS(t *testing.T) {
	testCases := []struct {
		c        uint64
		expected uint64
	}{
		{0, 0},
		{1, 1},
		{2, 3},
		{3, 6},
		{4, 10},
		{5, 15},
		{6, 21},
		{7, 28},
		{8, 36},
		{9, 45},
		{10, 64},
		{11, 93},
		{20, 1074},
		{55, 7999939},
		{83, 10999999911},
		{100, 799999999894},
	}

	for _, testCase := range testCases {
		answer := S(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}

func TestF(t *testing.T) {
	testCases := []struct {
		c        uint64
		expected uint64
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
		{11, 89},
		{20, 6765},
		{55, 139583862445},
		{90, 2880067194370816120},
	}

	for _, testCase := range testCases {
		answer := f(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}

// func TestXYZ(t *testing.T) {
// 	testCases := []struct {
// 		c        int
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
// 		answer := XYZ(testCase.c)
// 		if answer != testCase.expected {
// 			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
// 		}
// 	}
// }
