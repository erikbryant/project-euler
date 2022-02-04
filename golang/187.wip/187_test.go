package main

import (
	"testing"
)

func TestComposite2(t *testing.T) {
	testCases := []struct {
		c        int
		expected bool
	}{
		{0, false},
		{1, false},
		{2, false},
		{3, false},
		{4, true},
		{5, false},
		{6, true},
		{7, false},
		{8, false},
		{9, true},
		{10, true},
		{11, false},
		{12, false},
		{13, false},
		{14, true},
		{15, true},
		{16, false},
		{17, false},
		{18, false},
		{19, false},
		{20, false},
		{21, true},
		{22, true},
		{23, false},
		{24, false},
		{25, true},
		{26, true},
		{27, false},
		{28, false},
		{29, false},
		{30, false},
	}

	for _, testCase := range testCases {
		answer := composite2(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %t, got %t", testCase.c, testCase.expected, answer)
		}
	}
}
