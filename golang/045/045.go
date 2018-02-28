package main

import (
	"fmt"
	"math"
)

func isInt(n float64) bool {
	return n > 0 && n == float64(int(n))
}

func pentagonal(n int) bool {
	root := math.Sqrt(float64(24*n + 1))
	return isInt((1+root)/6) || isInt((1-root)/6)
}

func hexagonal(n int) bool {
	root := math.Sqrt(float64(8*n + 1))
	return isInt((1+root)/4) || isInt((1-root)/4)
}

func main() {
	n := 1
	for {
		triangle := (n*n + n) / 2
		if pentagonal(triangle) && hexagonal(triangle) {
			fmt.Println(n, " (", triangle, ") is Tri + Pent + Hex")
			if triangle > 40755 {
				break
			}
		}
		n++
	}
}
