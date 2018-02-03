package main

import (
	"fmt"
	"math"
)

func main() {
	// a^2 + b^2 = c^2
	// Brute force all combinations of a and b
	// checking to see which yield an integer c

	// Store the number of times a given perimeter is found
	solutions := make(map[int]int)

	for a := 1; a < 1000; a++ {
		for b := 1; b < 1000-a; b++ {
			c2 := a*a + b*b
			c := math.Sqrt(float64(c2))
			if c != float64(int(c)) {
				continue
			}
			p := a + b + int(c)
			if p > 1000 {
				break
			}
			solutions[p] += 1
		}
	}
	maxVal := 0
	maxKey := 0
	for key, val := range solutions {
		if val > maxVal {
			maxVal = val
			maxKey = key
		}
	}
	fmt.Println("Max number of solutions is ", maxVal, " for perimeter: ", maxKey)
}
