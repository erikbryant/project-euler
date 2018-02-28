package main

import (
	"testing"
)

var testCases = []struct {
	x       int
	count   int
	addends [][]int
}{
	{2, 1, [][]int{
		[]int{1, 1},
	}},
	{3, 2, [][]int{
		[]int{2, 1},
		[]int{1, 1, 1},
	}},
	{4, 4, [][]int{
		[]int{3, 1},
		[]int{2, 2},
		[]int{2, 1, 1},
		[]int{1, 1, 1, 1},
	}},
	{5, 6, [][]int{
		[]int{4, 1},
		[]int{3, 2},
		[]int{3, 1, 1},
		[]int{2, 2, 1},
		[]int{2, 1, 1, 1},
		[]int{1, 1, 1, 1, 1},
	}},
	{6, 10, [][]int{
		[]int{5, 1},
		[]int{4, 2},
		[]int{3, 3},
		[]int{4, 1, 1},
		[]int{3, 2, 1},
		[]int{2, 2, 2},
		[]int{3, 1, 1, 1},
		[]int{2, 2, 1, 1},
		[]int{2, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1},
	}},
	{7, 14, [][]int{
		[]int{6, 1},
		[]int{5, 2},
		[]int{4, 3},
		[]int{5, 1, 1},
		[]int{4, 2, 1},
		[]int{3, 3, 1},
		[]int{3, 2, 2},
		[]int{4, 1, 1, 1},
		[]int{3, 2, 1, 1},
		[]int{2, 2, 2, 1},
		[]int{3, 1, 1, 1, 1},
		[]int{2, 2, 1, 1, 1},
		[]int{2, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 1},
	}},
	{8, 21, [][]int{
		[]int{7, 1},
		[]int{6, 2},
		[]int{5, 3},
		[]int{4, 4},
		[]int{6, 1, 1},
		[]int{5, 2, 1},
		[]int{4, 3, 1},
		[]int{4, 2, 2},
		[]int{3, 3, 2},
		[]int{5, 1, 1, 1},
		[]int{4, 2, 1, 1},
		[]int{3, 3, 1, 1},
		[]int{3, 2, 2, 1},
		[]int{2, 2, 2, 2},
		[]int{4, 1, 1, 1, 1},
		[]int{3, 2, 1, 1, 1},
		[]int{2, 2, 2, 1, 1},
		[]int{3, 1, 1, 1, 1, 1},
		[]int{2, 2, 1, 1, 1, 1},
		[]int{2, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 1, 1},
	}},
}

// Validate that I typed in the test cases correctly :-)
func TestTestCases(t *testing.T) {
	for _, testCase := range testCases {
		if testCase.count != len(testCase.addends) {
			t.Errorf("ERROR: For %d expected len=%d got len=%d", testCase.x, testCase.count, len(testCase.addends))
		}
		for _, addends := range testCase.addends {
			sum := 0
			last := addends[0]
			for _, addend := range addends {
				if addend > last {
					t.Errorf("ERROR: For %d %v list is unordered at %d", testCase.x, addends, addend)
				}
				sum += addend
				last = addend
			}
			if sum != testCase.x {
				t.Errorf("ERROR: For %d %v got sum: %d", testCase.x, addends, sum)
			}
		}
	}
}

func TestCountSums(t *testing.T) {
	for _, testCase := range testCases {
		count := countSums(testCase.x)
		expected := uint64(len(testCase.addends))
		if count != expected {
			t.Errorf("ERROR: For countSums %d expected %d, got %d", testCase.x, len(testCase.addends), count)
		}
	}
}

func TestCountSumsFast(t *testing.T) {
	for _, testCase := range testCases {
		count := countSumsFast(testCase.x)
		expected := uint64(len(testCase.addends))
		if count != expected {
			t.Errorf("ERROR: For countSumsFast %d expected %d, got %d", testCase.x, len(testCase.addends), count)
		}
	}
}

func TestCountSumsFastVsCountSums(t *testing.T) {
	for x := 2; x <= 31; x++ {
		count := countSumsFast(x)
		expected := countSums(x)
		if count != expected {
			t.Errorf("ERROR: For countSumsFast %d expected %d, got %d", x, expected, count)
		}
	}
}
