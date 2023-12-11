package main

// go fmt ./... && go vet ./... && go test && go run 587.go

//
// A square is drawn around a circle as shown in the diagram below on the left.
// We shall call the blue shaded region the L-section. A line is drawn from the
// bottom left of the square to the top right as shown in the diagram on the right.
// We shall call the orange shaded region a concave triangle.
//
// It should be clear that the concave triangle occupies exactly half of the
// L-section.
//
// Two circles are placed next to each other horizontally, a rectangle is drawn
// around both circles, and a line is drawn from the bottom left to the top right
// as shown in the diagram below.
//
// This time the concave triangle occupies approximately 36.46% of the L-section.
//
// If n circles are placed next to each other horizontally, a rectangle is drawn
// around the n circles, and a line is drawn from the bottom left to the top right,
// then it can be shown that the least value of n for which the concave triangle
// occupies less than 10% of the L-section is n = 15.
//
// What is the least value of n for which the concave triangle occupies less than 0.1% of the L-section?
//

import (
	"fmt"
	"math"
)

func lSectionArea(radius float64) float64 {
	squareArea := 2 * radius * 2 * radius
	circleArea := math.Pi * radius * radius
	return (squareArea - circleArea) / 4.0
}

// guessXY returns a solution to x^2+y^2=1 and y=mx+b
// I tried to do the math to intersect the line with the circle
// to directly calculate the crossing point, but the number
// of steps involved was too great. There would have been errors.
func guessXY(m, b, radius float64) (float64, float64) {
	// Start guessing from where y = x intersects the circle
	x := -1.0 * math.Sqrt(0.5)
	y := x

	// Slowly increment x through solutions to:
	//   y = mx + b
	// until (x, y) is a solution to:
	//   x^2 + y^2 = 1

	delta := 0.00000001
	lastEpsilon := 999.0
	for {
		x += delta
		y = m*x + b
		result := x*x + y*y
		epsilon := math.Abs(result - 1.0)
		if epsilon > lastEpsilon {
			break
		}
		lastEpsilon = epsilon
	}

	return x, y
}

func concaveTriangleArea(n int, radius float64) float64 {
	// (x, y) of the line/circle intercept
	m := radius / float64(n)
	b := -1.0*radius + m
	x, y := guessXY(m, b, radius)

	// length(c) = sqrt(x^2 + (y - -1)^2)
	lenC := math.Sqrt(x*x + (y+1)*(y+1))

	// θ = asin(length(c)/2)*2
	theta := math.Asin(lenC/2.0) * 2.0

	// area(segment c) = 1/2 * r^2 * (θ - sin(θ))
	areaSegment := 0.5 * radius * radius * (theta - math.Sin(theta))

	// area(t) = 1/2 * 1 * (y - -1)
	areaTriangle := 0.5 * 1.0 * (y + 1)

	// area(a) = area(t) - area(segment c)
	ctArea := areaTriangle - areaSegment

	return ctArea
}

func main() {
	fmt.Printf("Welcome to 587\n\n")

	radius := 1.0

	lArea := lSectionArea(radius)
	fmt.Println("lArea =", lArea)

	for n, ctAreaPct := 1, 100.0; ctAreaPct >= 0.0999; n++ {
		ctArea := concaveTriangleArea(n, radius)
		ctAreaPct = ctArea / lArea * 100.0
		fmt.Printf("For n=%3d, radius=%0.2f, lArea=%0.4f, ctArea=%0.4f, ct = %0.6f%%\n", n, radius, lArea, ctArea, ctAreaPct)
		if n > 15 && ctAreaPct >= .11 {
			n += 200
		}
	}
}
