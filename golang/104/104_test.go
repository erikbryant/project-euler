package main

import "testing"

func TestPandigital(t *testing.T) {
	testCases := []struct {
		c        []byte
		expected bool
	}{
		{[]byte{}, false},
		{[]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, false},
		{[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}, true},
		{[]byte{9, 2, 3, 4, 5, 6, 7, 8, 1}, true},
		{[]byte{0, 2, 3, 4, 5, 6, 7, 8, 9}, false},
		{[]byte{1, 1, 3, 4, 5, 6, 7, 8, 9}, false},
	}

	for _, testCase := range testCases {
		answer := pandigital(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v expected %t, got %t", testCase.c, testCase.expected, answer)
		}
	}
}
