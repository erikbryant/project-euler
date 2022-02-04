package main

import (
	"testing"
)

func TestTangentSlope(t *testing.T) {
	testCases := []struct {
		p        Point
		expected float64
	}{
		{Point{x: 1, y: 1}, -4},
		{Point{x: 1.4, y: -9.6}, -4.0 * 1.4 / -9.6},
	}

	for _, testCase := range testCases {
		answer := tangentSlope(testCase.p)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v expected %f, got %f", testCase.p, testCase.expected, answer)
		}
	}
}

func TestIntersect(t *testing.T) {
	testCases := []struct {
		p        Point
		m        float64
		b        float64
		expected Point
	}{
		{Point{x: 5, y: 0}, 0.0, 0.0, Point{x: -5, y: 0}},
		{Point{x: -5, y: 0}, 0.0, 0.0, Point{x: 5, y: 0}},
		{Point{x: 0, y: -10}, -2, -10, Point{x: -5, y: 0}},
		{Point{x: 0, y: -10}, 2, -10, Point{x: 5, y: 0}},
	}

	for _, testCase := range testCases {
		answer := intersect(testCase.p, testCase.m, testCase.b)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v, %.2f, %.2f expected %v, got %v", testCase.p, testCase.m, testCase.b, testCase.expected, answer)
		}
	}
}
