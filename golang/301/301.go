package main

import (
	"fmt"
	"math"
)

func nimSum(a, b, c int) int {
	return a ^ b ^ c
}

// X returns the nim sum of n
func X(n int) int {
	return nimSum(n, 2*n, 3*n)
}

func main() {
	zero := 0
	maxFound := int(math.Pow(2.0, 30.0))
	for n := 1; n <= maxFound; n++ {
		// fmt.Printf("n: %3d %8b %4d\n", n, n, X(n))
		if X(n) == 0 {
			zero++
		}
	}
	fmt.Println(maxFound, zero)
}
