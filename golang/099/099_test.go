package main

import (
	"testing"
)

func TestS(t *testing.T) {
	testCases := []struct {
		b1       float64
		e1       int
		b2       float64
		e2       int
		expected int
		bMax     float64
		eMax     int
	}{
		// b1 == b2
		{10, 5, 10, 5, 0, 10, 5},
		{10, 6, 10, 5, -1, 10, 6},
		{10, 5, 10, 7, 1, 10, 7},

		// e1 == e2
		{10, 5, 10, 5, 0, 10, 5},
		{11, 5, 10, 5, -1, 11, 5},
		{10, 5, 12, 5, 1, 12, 5},

		// b1 > b2 && e1 > e2
		{11, 6, 10, 5, -1, 11, 6},

		// b2 > b1 && e2 > e1
		{10, 5, 12, 7, 1, 12, 7},

		// All other cases...
		{3, 2, 2, 3, -1, 3, 2},
		{5, 3, 12, 2, 1, 12, 2},
	}

	for _, testCase := range testCases {
		answer, bMax, eMax := compare(testCase.b1, testCase.e1, testCase.b2, testCase.e2)
		if answer != testCase.expected || bMax != testCase.bMax || eMax != testCase.eMax {
			t.Errorf("ERROR: For %.2f^%d <=> %.2f^%d expected %d: %.2f^%d, got %d: %.2f^%d", testCase.b1, testCase.e1, testCase.b2, testCase.e2, testCase.expected, testCase.bMax, testCase.eMax, answer, bMax, eMax)
		}
	}
}

func TestMax(t *testing.T) {
	testCases := []struct {
		n        [][2]int
		expected int
	}{
		{[][2]int{
			{999, 999},
			{888, 888},
			{777, 777},
		}, 1},
		{[][2]int{
			{888, 888},
			{999, 999},
			{777, 777},
		}, 2},
		{[][2]int{
			{888, 888},
			{777, 777},
			{999, 999},
		}, 3},
	}

	for _, testCase := range testCases {
		answer := maxFound(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}
