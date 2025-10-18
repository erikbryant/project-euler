package main

// go fmt ./... && go vet ./... && go test && go build 231.go && time ./231
// go fmt ./... && go vet ./... && go test && go build 231.go && ./231 && echo top | go tool pprof cpu.prof

import (
	"fmt"

	"github.com/erikbryant/util-golang/algebra"
)

// The binomial coefficient (10, 3) = 120.
// 120 = 2^3 x 3 x 5 = 2 x 2 x 2 x 3 x 5, and 2 + 2 + 2 + 3 + 5 = 14.
// So the sum of the terms in the prime factorisation of (10, 3) is 14.
//
// Find the sum of the terms in the prime factorisation of (20,000,000, 15,000,000).

// Equation for binomial coefficients:
//
//            n!
// ( n ) = ----------
// ( k )    k!(n-k)!
//

// findFactors returns the total set of factors in binomial(n, k)
func findFactors(n, k int) map[int]int {
	f := map[int]int{}

	// Accumulate factors
	for i := n; i > (n - k); i-- {
		for base, exp := range algebra.FactorsCounted(i) {
			f[base] += exp
		}
	}

	// Eliminate factors
	for i := k; i > 1; i-- {
		for base, exp := range algebra.FactorsCounted(i) {
			f[base] -= exp
		}
	}

	return f
}

// sumFactors returns the sum of all of the given factors
func sumFactors(f map[int]int) int {
	sum := 0
	for base, exp := range f {
		sum += base * exp
	}
	return sum
}

func main() {
	fmt.Printf("Welcome to 231\n\n")

	//c1 := 10
	//c2 := 3
	c1 := 20 * 1000 * 1000
	c2 := 15 * 1000 * 1000

	f := findFactors(c1, c2)
	sum := sumFactors(f)

	fmt.Printf("Sum of factors of binomial(%d, %d) = %d\n", c1, c2, sum)
}
