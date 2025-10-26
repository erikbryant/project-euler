package main

import (
	"math"
	"testing"

	"github.com/erikbryant/util-golang/common"
	"github.com/erikbryant/util-golang/matrices"
)

func TestName(t *testing.T) {
	testCases := []struct {
		c        int
		expected string
	}{
		{0, "GO"},
		{2, "CC1"},
		{7, "CH1"},
		{10, "JL"},
		{17, "CC2"},
		{22, "CH2"},
		{24, "E3"},
		{39, "H2"},
	}

	expected := 40
	answer := len(SquareNames)
	if answer != expected {
		t.Errorf("ERROR: For len(SquareNames) expected %d, got %d", expected, answer)
	}

	for _, testCase := range testCases {
		answer := name(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %s, got %s", testCase.c, testCase.expected, answer)
		}
	}
}

func TestIndex(t *testing.T) {
	testCases := []struct {
		c        string
		expected int
	}{
		{"GO", 0},
		{"CC1", 2},
		{"CH1", 7},
		{"JL", 10},
		{"CC2", 17},
		{"CH2", 22},
		{"E3", 24},
		{"H2", 39},
	}

	expected := 40
	answer := len(SquareNames)
	if answer != expected {
		t.Errorf("ERROR: For len(SquareNames) expected %d, got %d", expected, answer)
	}

	for _, testCase := range testCases {
		answer := index(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %s expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}

// boardInitVanilla sets up a board where the only rule is "roll 2 dice and move that distance"
func boardInitVanilla[T common.Floats](A matrices.Matrix[T], dieSides int) {
	// This is used as a test case. The probabilities should
	// stabilize to 1/40 for each square regardless of die sides.
	// It verifies that the board is created properly and that
	// the movement probabilities (in the simplest case) are
	// calculated correctly.

	n := T(dieSides)

	// Initialize all rows with the default transition
	for row := 0; row < A.Cols(); row++ {
		// Single roll of 2 dice
		for d1 := 1; d1 <= dieSides; d1++ {
			for d2 := 1; d2 <= dieSides; d2++ {
				sum := d1 + d2
				p := 1.0 / (n * n)
				nextSquare := (row + sum) % A.Cols()
				A[row][nextSquare] += p
			}
		}
	}

	boardCheck(A, "Board after init")
}

func TestVanilla(t *testing.T) {
	expected := 1.0 / 40.0 // Each square equally likely
	epsilon := 0.000001

	boardV, stateV := boardNew[float64]()
	boardInitVanilla(boardV, 4)
	transition(boardV, stateV, 350)

	for i, p := range stateV[0] {
		if math.Abs(p-expected) > epsilon {
			t.Errorf("Expected boardV[%d] = %f got %f", i, expected, p)
		}
	}

	boardV, stateV = boardNew[float64]()
	boardInitVanilla(boardV, 6)
	transition(boardV, stateV, 150)

	for i, p := range stateV[0] {
		if math.Abs(p-expected) > epsilon {
			t.Errorf("Expected boardV[%d] = %f got %f", i, expected, p)
		}
	}
}
