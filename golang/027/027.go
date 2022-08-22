package main

import (
	"fmt"
	"github.com/erikbryant/project-euler/golang/primes"
)

func init() {
	primes.Load("../primes.gob")
}

// Considering quadratics of the form:
//
// n^2+an+b, where |a|<1000 and |b|≤1000
//
// Find the product of the coefficients, a and b,
// for the quadratic expression that produces the
// maximum number of primes for consecutive values
// of n, starting with n=0.
func q(a, b int) int {
	primeCount := 0
	isPrime := true
	n := 0

	for isPrime {
		p := n*n + a*n + b
		if !primes.Prime(p) {
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
	max := 0
	product := 0

	for a := -bounds + 1; a < bounds; a++ {
		for b := -bounds; b <= bounds; b++ {
			primeCount := q(a, b)
			if primeCount > max {
				max = primeCount
				product = a * b
			}
		}
	}

	fmt.Println("max:", max, "product:", product)
	return product
}

func main() {
	fmt.Println("Welcome to 027")
	fmt.Println("answer:", findMax(1000))
}
