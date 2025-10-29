package main

// go fmt ./... && go vet ./... && go test ./... && go build 027.go && time ./027

import (
	"fmt"

	"github.com/erikbryant/util-golang/primey"
)

// Euler discovered the remarkable quadratic formula:
//
// n^2 + n + 41
//
// It turns out that the formula will produce 40 primes for the consecutive integer values
// 0 <= n <= 39. However, when n = 40, 40^2 + 40 + 41 = 40(40 + 1) + 41 is divisible by 41,
// and certainly when n = 41, 41^2 + 41 + 41 is clearly divisible by 41.
//
// The incredible formula n^2 - 79n + 1601 was discovered, which produces 80 primes for the
// consecutive values 0 <= n <= 79. The product of the coefficients, -79 and 1601, is -126479.
// Considering quadratics of the form:
//
// n^2 + an + b, where |a| < 1000 and |b| <= 1000
//
// where |n| is the modulus/absolute value of n
// e.g. |11| = 11 and |-4| = 4
//
// Find the product of the coefficients, a and b, for the quadratic expression that produces
// the maximum number of primes for consecutive values of n, starting with n = 0.

func q(a, b int) int {
	primeCount := 0
	n := 0

	for {
		p := n*n + a*n + b
		if !primey.Prime(p) {
			break
		}
		primeCount++
		n++
	}

	return primeCount
}

// Considering quadratics of the form:
//
// n^2+an+b, where |a|<1000 and |b|≤1000
//
// Find the product of the coefficients, a and b,
// for the quadratic expression that produces the
// maximum number of primes for consecutive values
// of n, starting with n=0.
func findMax(bounds int) int {
	maxFound := 0
	product := 0

	for a := -bounds + 1; a < bounds; a++ {
		for b := -bounds; b <= bounds; b++ {
			primeCount := q(a, b)
			if primeCount > maxFound {
				maxFound = primeCount
				product = a * b
			}
		}
	}

	fmt.Println("maxFound:", maxFound, "product:", product)
	return product
}

func main() {
	fmt.Println("Welcome to 027")
	fmt.Println("answer:", findMax(1000))
}
