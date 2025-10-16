package main

import (
	"slices"
	"testing"
)

func TestItoA(t *testing.T) {
	testCases := []struct {
		c        int
		expected []int8
	}{
		{0, []int8{0}},
		{5, []int8{5}},
		{10, []int8{1, 0}},
		{9998, []int8{9, 9, 9, 8}},
	}

	for _, testCase := range testCases {
		answer := ItoA(testCase.c)
		if !slices.Equal(answer, testCase.expected) {
			t.Errorf("ERROR: For %d expected %v, got %v", testCase.c, testCase.expected, answer)
		}
	}
}

func TestAtoI(t *testing.T) {
	testCases := []struct {
		c        []int8
		expected int
	}{
		{[]int8{0}, 0},
		{[]int8{5}, 5},
		{[]int8{1, 0}, 10},
		{[]int8{9, 9, 9, 8}, 9998},
	}

	for _, testCase := range testCases {
		answer := AtoI(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}

func TestCompare(t *testing.T) {
	testCases := []struct {
		a        []int8
		b        []int8
		expected int
	}{
		// Same length
		{[]int8{0}, []int8{1}, -1},
		{[]int8{0}, []int8{0}, 0},
		{[]int8{1}, []int8{0}, 1},
		{[]int8{1, 5}, []int8{1, 7}, -1},
		{[]int8{1, 5}, []int8{1, 5}, 0},
		{[]int8{1, 5}, []int8{1, 2}, 1},

		// a is longer
		{[]int8{1, 0}, []int8{9}, 1},
		{[]int8{4, 1, 0}, []int8{2, 9}, 1},

		// b is longer
		{[]int8{0}, []int8{9, 1}, -1},
		{[]int8{7, 7}, []int8{3, 9, 1}, -1},
	}

	for _, testCase := range testCases {
		answer := compare(testCase.a, testCase.b)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v > %v expected %d, got %d", testCase.a, testCase.b, testCase.expected, answer)
		}
	}
}

func TestSubtract(t *testing.T) {
	testCases := []struct {
		s         []int8
		i         int
		expectedS []int8
		expectedN bool
	}{
		{[]int8{4, 9}, 49, []int8{0}, false},
		{[]int8{2}, 2, []int8{0}, false},
		{[]int8{1, 0}, 8, []int8{2}, false},
		{[]int8{1}, 8, []int8{7}, true},
	}

	for _, testCase := range testCases {
		answerS, answerN := subtract(testCase.s, testCase.i)
		if !slices.Equal(answerS, testCase.expectedS) || answerN != testCase.expectedN {
			t.Errorf("ERROR: For %v-%d expected %v:%t, got %v:%t", testCase.s, testCase.i, testCase.expectedS, testCase.expectedN, answerS, answerN)
		}
	}
}

func TestPDRStr(t *testing.T) {
	testCases := []struct {
		n        int
		p        int
		expected bool
	}{
		// n == 0
		//{0, 7, true},

		// n < p
		{4, 3, false},

		// n == p
		{3, 3, true},

		// p == 2
		//{21, 2, false},
		//{22, 2, true},

		// p == 5
		//{0, 5, true},
		//{10, 5, true},
		//{35, 5, true},
		//{99, 5, false},

		// Actual cases
		{22, 3, false},
		{21, 3, true},
		{48, 7, false},
		{49, 7, true},
		{1000, 11, false},
		{1111, 11, true},

		// Various bugs
		{4334, 13, false},
		{24, 13, false},
		{33, 13, false},
		{34, 13, false},
		{7, 7, true},
		{4334, 11, true},
		{4334, 197, true},
		{8668, 11, true},
		{8668, 197, true},
		{13002, 3, true},
		{13002, 197, true},
		{17336, 197, true},
		{21670, 197, true},
		{26004, 3, true}, //
		{26004, 11, true},
		{26004, 197, true},
		{30338, 3, false},
		{30338, 7, true},
		{30338, 11, true},
		{30338, 197, true},
		{34672, 3, false},
		{34672, 11, true},
		{34672, 197, true},
		{39006, 11, true},
		{39006, 197, true},
		{43340, 11, true},
		{43340, 197, true},
		{47674, 11, true},
		{47674, 43, false},
		{47674, 197, true},
		{47674, 3407, false},
		{52008, 3, true},
		{52008, 11, true},
		{52008, 197, true},
		{56342, 11, true},
		{56342, 13, true},
		{56342, 197, true},
		{60676, 7, true},
		{60676, 11, true},
		{60676, 197, true},
		{60676, 3793, false},
		{65010, 3, true},
		{65010, 11, true},
		{65010, 197, true},
		{69344, 11, true},
		{69344, 197, true},
		{73678, 7, false},
		{73678, 11, true},
		{73678, 17, true},
		{73678, 197, true},
		{78012, 3, true},
		{78012, 7, false},
		{78012, 11, true},
		{78012, 197, true},

		// Prime factors of 1111111111111
		{1111111111111, 53, true},
		{1111111111111, 79, true},
		{1111111111111, 265371653, true},

		// Prime factors of 11111111111111
		{11111111111111, 11, true},
		{11111111111111, 239, true},
		{11111111111111, 4649, true},
		{11111111111111, 909091, true},
	}

	for _, testCase := range testCases {
		nStr := ItoA(testCase.n)
		answer := PDRStr(nStr, testCase.p)
		if answer != testCase.expected {
			t.Errorf("ERROR: For string(%d)/%d expected %t, got %t", testCase.n, testCase.p, testCase.expected, answer)
		}
	}
}
