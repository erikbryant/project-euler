package digits_413

import (
	"testing"
)

func TestFastForwardZero(t *testing.T) {
	testCases := []struct {
		digits   []int
		expected []int
	}{
		{[]int{9, 0, 3, 0, 2, 2}, []int{9, 0, 3, 1, 1, 1}},
		{[]int{1, 0, 0, 0}, []int{1, 0, 1, 1}},
	}

	for _, testCase := range testCases {
		fastForwardZero(testCase.digits, len(testCase.digits))
		if len(testCase.digits) != len(testCase.expected) {
			t.Errorf("ERROR: Expected %v, got %v", testCase.digits, testCase.expected)
		}
		for i := 0; i < len(testCase.expected); i++ {
			if testCase.digits[i] != testCase.expected[i] {
				t.Errorf("ERROR: For %v -> %v[%d] expected %d, got %d", testCase.digits, testCase.expected, i, testCase.expected[i], testCase.digits[i])
			}
		}
	}
}

func TestFastForwardIndex(t *testing.T) {
	testCases := []struct {
		digits   []int
		index    int
		expected []int
	}{
		{[]int{4, 4, 3, 5}, 1, []int{4, 5, 1, 1}},
		{[]int{1, 5, 5, 5, 3}, 2, []int{1, 5, 6, 1, 1}},
	}

	for _, testCase := range testCases {
		fastForwardIndex(testCase.digits, len(testCase.digits), testCase.index)
		if len(testCase.digits) != len(testCase.expected) {
			t.Errorf("ERROR: Expected %v, got %v", testCase.digits, testCase.expected)
		}
		for i := 0; i < len(testCase.expected); i++ {
			if testCase.digits[i] != testCase.expected[i] {
				t.Errorf("ERROR: For %v -> %v[%d] expected %d, got %d", testCase.digits, testCase.expected, i, testCase.expected[i], testCase.digits[i])
			}
		}
	}
}

func TestIncrement(t *testing.T) {
	testCases := []struct {
		digits   []int
		expected []int
		ok       bool
	}{
		{[]int{0}, []int{1}, true},
		{[]int{9}, []int{0}, true},
		{[]int{0, 9}, []int{1, 0}, true},
		{[]int{2, 9, 9, 9}, []int{3, 0, 1, 1}, true},
		{[]int{9, 9, 9, 9}, []int{0, 0, 0, 0}, false},
	}

	for _, testCase := range testCases {
		ok := increment(testCase.digits, len(testCase.digits))
		if ok == testCase.ok && ok == false {
			continue
		}
		if len(testCase.digits) != len(testCase.expected) {
			t.Errorf("ERROR: Expected %v, got %v", testCase.digits, testCase.expected)
		}
		for i := 0; i < len(testCase.expected); i++ {
			if testCase.digits[i] != testCase.expected[i] {
				t.Errorf("ERROR: For %v[%d] expected %d, got %d", testCase.expected, i, testCase.expected[i], testCase.digits[i])
			}
		}
	}
}

func TestCountChildren(t *testing.T) {
	testCases := []struct {
		digits        []int
		divisor       int
		end           int
		slen          int
		expectedCount int
	}{
		{[]int{1}, 1, 0, 1, 1},

		{[]int{1, 0, 4}, 3, 2, 3, 0},
		{[]int{1, 0, 4}, 3, 2, 2, 0},
		{[]int{1, 0, 4}, 3, 2, 1, 1},

		{[]int{5, 6, 7, 1}, 4, 3, 4, 0},
		{[]int{5, 6, 7, 1}, 4, 3, 3, 0},
		{[]int{5, 6, 7, 1}, 4, 3, 2, 1},
		{[]int{5, 6, 7, 1}, 4, 3, 1, 0},

		{[]int{1, 1, 3, 2, 4, 5, 1}, 7, 6, 7, 0},
		{[]int{1, 1, 3, 2, 4, 5, 1}, 7, 6, 6, 0},
		{[]int{1, 1, 3, 2, 4, 5, 1}, 7, 6, 5, 0},
		{[]int{1, 1, 3, 2, 4, 5, 1}, 7, 6, 4, 0},
		{[]int{1, 1, 3, 2, 4, 5, 1}, 7, 6, 3, 1},
		{[]int{1, 1, 3, 2, 4, 5, 1}, 7, 6, 2, 0},
		{[]int{1, 1, 3, 2, 4, 5, 1}, 7, 6, 1, 0},
	}

	for _, testCase := range testCases {
		count, _ := countChildrenOdd(testCase.digits, testCase.divisor, testCase.end, testCase.slen)
		if count != testCase.expectedCount {
			t.Errorf("ERROR: For %v/%d expected %d, got %d", testCase.digits, testCase.slen, testCase.expectedCount, count)
		}
	}
}

func TestOneChildManual(t *testing.T) {
	testCases := []struct {
		digits   int
		expected int
	}{
		{1, 9},
		{2, 20},
		{3, 360},
		{4, 2701},
		{5, 4096},
		// {6, 109466},
		// {7, 161022},
		// {8, 13068583},
		// {9, 2136960},
		// {10, 0},
		// {11, 71101800},
		// {12, 55121700430},
		// {13, 1057516028},
		// {14, },
		// {15, },
		// {16, },
		// {17, },
		// {18, },
		// {19, },
	}

	for _, testCase := range testCases {
		answer := router(testCase.digits, testCase.digits)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.digits, testCase.expected, answer)
		}
	}
}

func TestF(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{10, 9},
		{1000, 389},
		{1000 * 1000 * 10, 277674},
	}

	for _, testCase := range testCases {
		answer := F(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}
