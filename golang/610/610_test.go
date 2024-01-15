package main

import (
	"testing"
)

func EqualMap(a, b map[string]bool) bool {
	if len(a) != len(b) {
		return false
	}

	for key, val := range a {
		if b[key] != val {
			return false
		}
	}

	return true
}

func TestPossibles(t *testing.T) {
	testCases := []struct {
		c        string
		expected map[string]bool
	}{
		{"", map[string]bool{
			"I": true,
			"X": true,
			"C": true,
			"D": true,
		}},
		{"I", map[string]bool{
			"I": true,
			"V": true,
			"X": true,
		}},
	}

	romans := []string{
		"I",
		"X",
		"C",
		"D",
		"II",
		"IV",
		"IX",
	}

	for _, testCase := range testCases {
		answer := possibles(testCase.c, romans)
		if !EqualMap(answer, testCase.expected) {
			t.Errorf("ERROR: For '%s' expected %v, got %v", testCase.c, testCase.expected, answer)
		}
	}
}

// func TestProbability(t *testing.T) {
// 	testCases := []struct {
// 		c         string
// 		expected  int
// 		expected2 int
// 	}{
// 		{"#", 1, 15},
// 		{"I#", 7, 165},
// 		{"C#", 7, 270},
// 	}

// 	romans := []string{
// 		"#",
// 		"I#",
// 		"II#",
// 		"IV#",
// 		"IX#",
// 		"CI#",
// 		"CV#",
// 		"CX#",
// 		"CL#",
// 		"CC#",
// 	}

// 	for _, testCase := range testCases {
// 		answer, answer2 := probability(testCase.c, romans)
// 		if answer != testCase.expected || answer2 != testCase.expected2 {
// 			t.Errorf("ERROR: For '%s' expected %d/%d, got %d/%d", testCase.c, testCase.expected, testCase.expected2, answer, answer2)
// 		}
// 	}
// }
