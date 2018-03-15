package main

import (
	"testing"
)

func TestHighCard(t *testing.T) {
	testCases := []struct {
		hand     Hand
		expected []int
	}{
		{Hand{
			{9, 'C'},
			{8, 'C'},
			{7, 'C'},
			{6, 'C'},
			{5, 'C'},
		},
			[]int{9, 8, 7, 6, 5},
		},
		{Hand{
			{9, 'C'},
			{8, 'C'},
			{8, 'C'},
			{6, 'C'},
			{5, 'C'},
		},
			[]int{9, 8, 6, 5, 0},
		},
		{Hand{
			{8, 'C'},
			{8, 'C'},
			{8, 'C'},
			{8, 'C'},
			{8, 'C'},
		},
			[]int{8, 0, 0, 0, 0},
		},
	}

	for _, testCase := range testCases {
		for n := 1; n <= 5; n++ {
			answer := testCase.hand.highCard(n)
			if answer != testCase.expected[n-1] {
				t.Errorf("ERROR: For %v, %d expected %d, got %d", testCase.hand, n, testCase.expected[n-1], answer)
			}
		}
	}
}

func TestRoyalFlush(t *testing.T) {
	testCases := []struct {
		hand     Hand
		expected bool
	}{
		{Hand{
			{9, 'C'},
			{8, 'C'},
			{7, 'C'},
			{6, 'C'},
			{5, 'C'},
		},
			false,
		},
		{Hand{
			{14, 'C'},
			{13, 'C'},
			{12, 'C'},
			{11, 'C'},
			{10, 'D'},
		},
			false,
		},
		{Hand{
			{14, 'C'},
			{13, 'C'},
			{12, 'C'},
			{11, 'C'},
			{10, 'C'},
		},
			true,
		},
	}

	for _, testCase := range testCases {
		answer := testCase.hand.royalFlush()
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v, expected %t, got %t", testCase.hand, testCase.expected, answer)
		}
	}
}
