package main

import "testing"

func TestModSub(t *testing.T) {
	testCases := []struct {
		c        int64
		d        int64
		m        int64
		expected int64
	}{
		{5, 1, 1000000007, 4},
		{5, 6, 1000000007, 1000000006},
	}

	for _, testCase := range testCases {
		answer := modSub(testCase.c, testCase.d, testCase.m)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d - %d mod %d expected %d, got %d", testCase.c, testCase.d, testCase.m, testCase.expected, answer)
		}
	}
}

func TestExp10Mod(t *testing.T) {
	testCases := []struct {
		c        int64
		m        int64
		expected int64
	}{
		{0, 1000000007, 1},
		{1, 1000000007, 10},
		{2, 1000000007, 100},
		{3, 1000000007, 1000},
		{9, 1000000007, 1000000000},
		{10, 1000000007, 999999937},
		{18, 1000000007, 49},
		{36, 1000000007, 2401},
		{999, 1000000007, 522173006},
		{1000000005, 1000000007, 700000005},
		{1000000006, 1000000007, 1},
		{1000000007, 1000000007, 10},
		{1000000012, 1000000007, 1000000},
	}

	for _, testCase := range testCases {
		answer := exp10Mod(testCase.c, testCase.m)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d mod %d expected %d, got %d", testCase.c, testCase.m, testCase.expected, answer)
		}
	}
}

func TestSMod(t *testing.T) {
	testCases := []struct {
		c        int64
		m        int64
		expected int64
	}{
		// Verify summation
		{0, 1000000007, 0},
		{1, 1000000007, 1},
		{2, 1000000007, 3},
		{3, 1000000007, 6},
		{4, 1000000007, 10},
		{5, 1000000007, 15},
		{6, 1000000007, 21},
		{7, 1000000007, 28},
		{8, 1000000007, 36},
		{9, 1000000007, 45},
		{10, 1000000007, 64},
		{11, 1000000007, 93},
		{12, 1000000007, 132},
		{13, 1000000007, 181},
		{14, 1000000007, 240},
		{15, 1000000007, 309},
		{16, 1000000007, 388},
		{17, 1000000007, 477},
		{18, 1000000007, 576},
		{19, 1000000007, 775},
		{20, 1000000007, 1074},
		{21, 1000000007, 1473},
		{34, 1000000007, 40960},
		{55, 1000000007, 7999939},

		// Verify mod
		{0, 10, 0},
		{1, 10, 1},
		{2, 10, 3},
		{3, 10, 6},
		{4, 10, 0},
		{5, 10, 5},
		{6, 10, 1},
		{7, 10, 8},
		{8, 10, 6},
		{9, 30, 15},
		{10, 30, 4},
		{11, 30, 3},
		{12, 30, 12},
		{13, 30, 1},
		{14, 30, 0},
		{15, 101, 6},
		{16, 101, 85},
		{17, 101, 73},
		{18, 101, 71},
		{19, 101, 68},
		{20, 101, 64},
		{21, 101, 59},
		{34, 101, 55},
		{55, 101, 32},
	}

	for _, testCase := range testCases {
		answer := SMod(testCase.c, testCase.m)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d mod %d expected %d, got %d", testCase.c, testCase.m, testCase.expected, answer)
		}
	}
}

func TestFib(t *testing.T) {
	testCases := []struct {
		c        int64
		expected int64
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
		{12, 144},
	}

	for _, testCase := range testCases {
		answer := fib(testCase.c)
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
