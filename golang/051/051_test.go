package main

import (
	"testing"
)

func TestPrime(t *testing.T) {
	testCases := []struct {
		n        []int
		expected bool
	}{
		{[]int{5, 6, 0, 0, 3}, true},
		{[]int{2}, true},
		{[]int{2, 3}, true},
		{[]int{3, 2}, false},
	}

	for _, testCase := range testCases {
		answer := prime(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v expected %t, got %t", testCase.n, testCase.expected, answer)
		}
	}
}

func TestCopy(t *testing.T) {
	testCases := []struct {
		expected []int
	}{
		{[]int{5, 6, 0, 0, 3}},
		{[]int{2}},
		{[]int{2, 3}},
		{[]int{1, 2, 3, 0}},
		{[]int{0}},
	}

	for _, testCase := range testCases {
		answer := copy(testCase.expected)
		if len(answer) != len(testCase.expected) {
			t.Errorf("ERROR: For %v expected %v, got %v", testCase.expected, testCase.expected, answer)
		}
		for i := 0; i < len(testCase.expected); i++ {
			if answer[i] != testCase.expected[i] {
				t.Errorf("ERROR: For %v expected %v, got %v", testCase.expected, testCase.expected, answer)
			}
		}
	}
}

func TestReplacements(t *testing.T) {
	testCases := []struct {
		digits   []int
		common   []int
		expected int
	}{
		{[]int{5, 6, 0, 0, 3}, []int{2, 3}, 7},
		{[]int{2}, []int{0}, 4},
		{[]int{2, 3}, []int{0}, 6},
		{[]int{1, 2, 3, 0}, []int{0}, 0},
	}

	for _, testCase := range testCases {
		answer := replacements(testCase.digits, testCase.common)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v %v expected %d, got %d", testCase.digits, testCase.common, testCase.expected, answer)
		}
	}
}

// func TestCombinationsX(t *testing.T) {
// 	testCases := []struct {
// 		digits   []int
// 		n        int
// 		expected [][]int
// 	}{
// 		{
// 			[]int{1, 2, 3},
// 			3,
// 			[][]int{
// 				[]int{1, 2, 3},
// 			},
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		c := make(chan []int, 10)

// 		i := 0
// 		combinationsX(testCase.digits, testCase.n, c)
// 		for {
// 			answer, ok := <-c
// 			if !ok {
// 				break
// 			}
// 			if len(answer) != len(testCase.expected[i]) {
// 				t.Errorf("ERROR: For %v %d expected %v, got %v", testCase.digits, testCase.n, testCase.expected, answer)
// 			}
// 			i++
// 		}
// 		close(c)
// 	}
// }
