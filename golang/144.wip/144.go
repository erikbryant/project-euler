package main

import (
	"fmt"
	"math"
)

// Point represents a cartesian (x,y) point
type Point struct {
	x float64
	y float64
}

func tangentSlope(p Point) float64 {
	return (-4.0 * p.x) / p.y
}

func epsilonEqual(p1, p2 Point) bool {
	epsilon := 0.0000001
	diffX := p1.x - p2.x
	diffY := p1.y - p2.y
	return math.Abs(diffX) < epsilon && math.Abs(diffY) < epsilon
}

// further() returns a or b; whichever is further from x.
func further(x, a, b float64) float64 {
	xa := math.Abs(x - a)
	xb := math.Abs(x - b)
	if xa > xb {
		return a
	}
	return b
}

// intersects() returns the [other] point at which a line intersects the ellipse.
func intersect(p Point, m, b float64) Point {
	// Ellipse: 4x^2 + y^2 = 100
	// Line: y = mx + b
	// Solve for x:
	// 4x^2 + (mx + b)^2 = 100
	// 4x^2 + m^2x^2 + 2bmx + b^2 = 100
	// (4+m^2)x^2 + 2bmx + b^2 - 100 = 0
	// xa = (4+m*m)
	// xb = 2*b*m
	// xc = b^2 - 100

	// Quadratic equation:
	// x1 = (-b + Sqrt(b^2 - 4ac)) / 2a
	// x2 = (-b - Sqrt(b^2 - 4ac)) / 2a

	// Substituting:
	// x1 = (-(2bm) + Sqrt((2bm)^2 - 4(4+m^2)(b^2-100))) / 2(4+m^2)
	// x2 = (-(2bm) - Sqrt((2bm)^2 - 4(4+m^2)(b^2-100))) / 2(4+m^2)

	root := math.Sqrt((2.0*b*m)*(2.0*b*m) - 4.0*(4.0+m*m)*(b*b-100.0))
	x1 := (-2.0*b*m + root) / (2.0 * (4.0 + m*m))
	x2 := (-2.0*b*m - root) / (2.0 * (4.0 + m*m))

	// Given these x values, find the y values that satisfy the
	// ellipse equation. Then see which two satisfy the line equation.
	//
	// 4x^2 + y^2 = 100
	// y = math.Sqrt(100 - 4x^2)
	cross := Point{}
	foundY := math.Sqrt(100.0 - 4.0*x1*x1)
	if foundY == m*x1+b {
		candidate := Point{x: x1, y: foundY}
		if !epsilonEqual(candidate, p) {
			cross = candidate
		}
	}
	foundY = -foundY
	if foundY == m*x1+b {
		candidate := Point{x: x1, y: foundY}
		if !epsilonEqual(candidate, p) {
			cross = candidate
		}
	}
	foundY = math.Sqrt(100.0 - 4.0*x2*x2)
	if foundY == m*x2+b {
		candidate := Point{x: x2, y: foundY}
		if !epsilonEqual(candidate, p) {
			cross = candidate
		}
	}
	foundY = -foundY
	if foundY == m*x2+b {
		candidate := Point{x: x2, y: foundY}
		if !epsilonEqual(candidate, p) {
			cross = candidate
		}
	}

	return cross
}

func reflectionSlope() {
}

// bounce() takes a starting point, an ending point, and returns where the resulting vector will hit the cell.
func bounce(p1, p2 Point) (Point, Point) {

	// TODO: Figure out where the laser will next hit...

	return p2, Point{x: 0.0, y: 0.0}
}

func main() {
	fmt.Println("Welcome to 144. It has lasers! :-)")

	p1 := Point{x: 0.0, y: 10.1}
	p2 := Point{x: 1.4, y: -9.6}

	reflections := 0
	for {
		fmt.Println("Line", p1, "->", p2)
		p1, p2 = bounce(p1, p2)
		reflections++

		// The section corresponding to −0.01 ≤ x ≤ +0.01 at the top is
		// missing, allowing the light to enter and exit through the hole.
		if p2.x >= -0.01 && p2.x <= 0.01 && p2.y >= 0 {
			break
		}
	}

	fmt.Println("Reflections:", reflections, "Exit:", p2)
}
