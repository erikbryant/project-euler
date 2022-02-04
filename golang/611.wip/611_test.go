package main

import (
	"testing"
)

func TestWalk(t *testing.T) {
	hallways := []struct {
		length   int
		expected int
	}{
		{5, 1},
		{100, 27},
		{1000, 233},
		{1000 * 10, 1812},
		{1000 * 100, 14023},
		{1000 * 1000, 112168},
		{1000 * 1000 * 10, 927208},
		{1000 * 1000 * 100, 7880154},
		// {1000 * 1000 * 1000, 68496000},
		// {1000 * 1000 * 1000 * 10, 0},
		// {1000 * 1000 * 1000 * 100, 0},
		// {1000 * 1000 * 1000 * 1000, 0},
	}

	for _, hallway := range hallways {
		open := walk(hallway.length)
		if open != hallway.expected {
			t.Errorf("ERROR: For length %d expected %d, got %d", hallway.length, hallway.expected, open)
		}
	}
}
