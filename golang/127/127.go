package main

// go fmt ./... && go vet ./... && go test && go build 127.go && time ./127

import (
	"fmt"

	"github.com/erikbryant/util-golang/algebra"
)

// The radical of n, rad(n), is the product of distinct prime factors of n.
// For example, 504 = 2^3 x 3^2 x 7, so rad(504) = 2 x 3 x 7 = 42.
//
// We shall define the triplet of positive integers (a, b, c) to be an abc-hit if:
//   gcd(a, b) = gcd(a, c) = gcd(b, c) = 1
//   a < b
//   a + b = c
//   rad(abc) < c
//
// For example, (5, 27, 32) is an abc-hit, because:
//   gcd(5, 27) = gcd(5, 32) = gcd(27, 32) = 1
//   5 < 27
//   5 + 27 = 32
//   rad(4320) = 30 < 32
//
// It turns out that abc-hits are quite rare and there are only thirty-one abc-hits
// for c < 1000, with Σ c = 12523.
//
// Find Σ c for c < 120,000.

// initFactors returns the prime factors of all integers <= k and the products of those factors
func initFactors(k int) ([]map[int]bool, []int) {
	factors := []map[int]bool{}
	products := []int{}

	for i := 0; i <= k; i++ {
		m := map[int]bool{}
		product := 1
		for _, factor := range algebra.Factors(i) {
			m[factor] = true
			product *= factor
		}
		factors = append(factors, m)
		products = append(products, product)
	}

	return factors, products
}

// coprime returns true if a and b are coprime (have no common prime factors)
func coprime(a, b int, factors []map[int]bool) bool {
	for factor, _ := range factors[a] {
		if factors[b][factor] {
			return false
		}
	}
	return true
}

// abcHit returns true if a, b, and c are an abc-hit
func abcHit(a, b, c int, factors []map[int]bool, products []int) bool {
	//   rad(abc) < c
	if products[a]*products[b]*products[c] >= c {
		return false
	}

	//   gcd(a, b) = gcd(a, c) = gcd(b, c) = 1
	if !coprime(a, b, factors) {
		return false
	}
	if !coprime(a, c, factors) {
		return false
	}
	if !coprime(b, c, factors) {
		return false
	}

	return true
}

// countAbcHits returns the count and sum of the abc-hits less than upper
func countAbcHits(upper int) (int, int) {
	count := 0
	sum := 0

	factors, products := initFactors(upper)

	for c := 1; c < upper; c++ {
		if factors[c][c] {
			// If c is prime then rad(a*b*c) will be > c
			continue
		}

		a := 1
		b := c - a
		for a < b {
			if abcHit(a, b, c, factors, products) {
				count++
				sum += c
				//fmt.Printf("a: %6d  b: %6d  c: %6d\n  %v\n  %v\n  %v\n", a, b, c, factors[a], factors[b], factors[c])
			}
			a++
			b--
		}
	}

	return count, sum
}

func main() {
	fmt.Printf("Welcome to 127\n\n")

	upper := 120000

	count, sum := countAbcHits(upper)

	fmt.Printf("There are %d abc-hits < %d with a sum of %d\n", count, upper, sum)
}
