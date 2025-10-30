package main

// go fmt ./... && go vet ./... && go test ./... && go build 757.go && time ./757

import (
	"fmt"
)

// A positive integer N is stealthy, if there exist positive
// integers a, b, c, d such that ab = cd = N and a+b = c+d+1.
// For example, 36 = 4x9 = 6x6 is stealthy.
//
// You are also given that there are 2851 stealthy numbers not
// exceeding 10^6.
//
// How many stealthy numbers are there that don't exceed 10^14?

// Odd numbers cannot be stealthy
// a + b = c + d + 1
// If n is odd then all of its divisors are odd
// ODD + ODD -> ODD
// ODD + ODD + 1 -> EVEN

// Pronic numbers are of the form   x(x+1)
// Bipronic numbers are of the form x(x+1) * y(y+1)
// That is, the product of two pronic numbers.
// All bipronic numbers > 0 are stealthy.

func main() {
	fmt.Printf("Welcome to 757\n\n")

	n := 1000 * 1000 * 1000 * 1000 * 100

	// Find all pronic numbers <= n
	pronics := []int{}
	i := 1
	for {
		p := i * (i + 1)
		if p > n {
			break
		}
		pronics = append(pronics, p)
		i++
	}

	fmt.Printf("Pronics found: %d\n", len(pronics))

	// Find all products of pronic numbers where the
	// product is <= n.
	bipronics := map[int]bool{}
	for i := 0; i < len(pronics); i++ {
		maxP := n / pronics[i]
		for j := i; pronics[j] <= maxP; j++ {
			p := pronics[i] * pronics[j]
			bipronics[p] = true
		}
	}

	fmt.Printf("\nStealthy numbers < 10^14 = %d\n\n", len(bipronics))
}
