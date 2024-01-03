package main

import (
	"slices"
	"testing"

	"github.com/erikbryant/util-golang/util"
)

// func TestPermuteSum(t *testing.T) {
// }

func TestVerify(t *testing.T) {
	testCases := []struct {
		target   int
		p        [][]int
		expected bool
	}{
		{0, [][]int{}, true},
		{12, [][]int{
			{12},
			{1, 3, 4, 4},
			{1, 3, 4, 4, 0},
			{1, 3, 4, 4, 0, 0},
		}, true},
		{11, [][]int{
			{12},
			{1, 3, 4, 4},
			{1, 3, 4, 4, 0},
			{1, 3, 4, 4, 0, 0},
		}, false},
	}

	for _, testCase := range testCases {
		answer := verify(testCase.target, testCase.p)
		if (answer == nil) != testCase.expected {
			t.Errorf("ERROR: For %d %v expected %t, got %t", testCase.target, testCase.p, testCase.expected, answer == nil)
		}
	}
}

// func TestFixLengthsAndZeroes(t *testing.T) {
// }

// func TestRemoveDuplicatePermutations(t *testing.T) {
// }

// func TestPermutations(t *testing.T) {
// }

func TestOtherHalf(t *testing.T) {
	testCases := []struct {
		digits   []int
		pool     []int
		expected []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{4}, []int{4, 5}, []int{5}},
		{[]int{0, 1, 2, 3}, []int{4, 5, 7, 6, 0, 3, 1, 2}, []int{7, 6, 5, 4}},
		{[]int{}, []int{4, 5}, []int{5, 4}},
	}

	for _, testCase := range testCases {
		answer := otherHalf(testCase.digits, testCase.pool)
		// util.Equal() expects the slices to both be sorted
		// answer is already sorted
		slices.Sort(testCase.expected)
		if !util.Equal(answer, testCase.expected) {
			t.Errorf("ERROR: For %v %v expected %v, got %v", testCase.digits, testCase.pool, testCase.expected, answer)
		}
	}
}

func TestMakeInt(t *testing.T) {
	testCases := []struct {
		digits   []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{0}, 0},
		{[]int{0, 0}, 0},
		{[]int{0, 1}, 1},
		{[]int{9, 9, 8, 8, 7, 7}, 998877},
		{[]int{0, 9, 0}, 90},
		{[]int{9, 9, 8, 8, 7, 7, 6, 6, 5, 5}, 9988776655},
		{[]int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4}, 11223344},
	}

	for _, testCase := range testCases {
		answer := makeInt(testCase.digits)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.digits, testCase.expected, answer)
		}
	}
}

func TestPairCount(t *testing.T) {
	testCases := []struct {
		digits    []int
		expected  int
		expected2 int
	}{
		{[]int{}, 0, 0},
		{[]int{1, 1}, 1, 0},
		{[]int{0}, 0, 1},
		{[]int{0, 0}, 1, 2},
		{[]int{9, 8, 0, 7, 9}, 1, 1},
		{[]int{9, 8, 7, 6, 9, 8, 7, 6}, 4, 0},
	}

	for _, testCase := range testCases {
		answer, answer2 := pairCount(testCase.digits)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d %d, got %d %d", testCase.digits, testCase.expected, testCase.expected2, answer, answer2)
		}
	}
}

func TestFactorial(t *testing.T) {
	testCases := []struct {
		c        int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{9, 362880},
		{10, 3628800},
	}

	for _, testCase := range testCases {
		answer := factorial(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}

func TestComboCompute(t *testing.T) {
	testCases := []struct {
		c        []int
		ignore   bool
		expected int
	}{
		{[]int{}, false, 0},
		{[]int{9}, false, 1},
		{[]int{1, 2}, false, 2},
		{[]int{9, 9, 8, 8}, false, 6},
		{[]int{0, 9, 8}, false, 6},
		{[]int{0, 2, 3, 0}, false, 12},
		{[]int{0, 9, 8}, true, 4},
		{[]int{0, 2, 3, 0}, true, 6},
	}

	for _, testCase := range testCases {
		answer := comboCompute(testCase.c, testCase.ignore)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d %t expected %d, got %d", testCase.c, testCase.ignore, testCase.expected, answer)
		}
	}
}

// func TestCombinations(t *testing.T) {
// }

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

// 	for _, testCase := range testCases {
// 		answer := XYZ(testCase.c)
// 		if answer != testCase.expected {
// 			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
// 		}
// 	}
// }
