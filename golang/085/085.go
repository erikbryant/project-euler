package main

import (
	"fmt"
	"math"
)

func rectangles(width, height int) int {
	count := 0

	// Look at each of the different heights the
	// rectangles could have. From 1 to the same
	// as the height of the bounding area.
	for h := 1; h <= height; h++ {
		// Look at the different rectangles that would
		// fit width-wise.
		wCount := 0
		for w := 1; w <= width; w++ {
			// Count this rectangle.
			wCount++
			// Count the extra spaces that this rectangle
			// could be slid across over.
			wCount += width - w
		}
		// Now see how many extra spaces downward those
		// rectangles could be slid down over.
		count += wCount + wCount*(height-h)
	}

	return count
}

func main() {
	fmt.Println("Welcome to 085")

	target := 2000000
	closestW := 0
	closestH := 0
	closestD := 99999999

	// A 2x3 rectangle is identical to a 3x2 rectangle. Only
	// iterate over one of the two.
	for width := 1; width <= 100; width++ {
		for height := 1; height <= width; height++ {
			r := rectangles(width, height)
			dist := int(math.Abs(float64(target - r)))
			if dist < closestD {
				closestD = dist
				closestW = width
				closestH = height
				fmt.Println("Target:", target, "dist:", closestD, "width:", closestW, "height:", closestH, "area:", closestW*closestH)
			}
		}
	}

	fmt.Println("Target:", target, "dist:", closestD, "width:", closestW, "height:", closestH, "area:", closestW*closestH)
}
