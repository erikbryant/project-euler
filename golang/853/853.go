package main

import (
	"fmt"
	"math"

	"github.com/erikbryant/util-golang/primes"
)

var (
	piCache = map[int]int{1: 0}
)

// go fmt ./... && go vet ./... && go test && go build 853.go && time ./853

// For every positive integer n the Fibonacci sequence modulo n is periodic.
// The period depends on the value of n. This period is called the Pisano period
// for n, often shortened to π(n).
//
// There are three values of π(n) for which π(n) equals 18: 19, 38, and 76.
// The sum of those smaller than 50 is 57.
//
// Find the sum of the values of n smaller than 1,000,000,000 for which π(n)
// equals 120.

// Notes:
// * mod(n) of each value in the fibonacci sequence is identical to mod(n) of
//   each element as the sequence is being created.
// * pi(p^x) == pi(p)*p^(x-1) (where p is a prime)
//   This means there is no need to compute pi(p), pi(p^2), pi(p^3), ...
//   It is sufficient to compute pi(p) and then multiply it by p^x.
// * pi(n) is equal to LMC[ pi(p1), pi(p2), pi(p3), ... ] where p1, p2, ...
//   are the prime factors of n.

// fibonacciModPeriod returns the period of the {Fibonacci sequence mod n} sequence
func fibonacciModPeriod(n, limit int) int {
	pi, ok := piCache[n]
	if ok {
		return pi
	}

	a := 1
	b := 1

	i := 2
	for {
		next := (a + b) % n
		if b == 0 && next == 1 {
			piCache[n] = i
			return i
		}
		a = b
		b = next
		i++
		if i > limit {
			return -1
		}
	}
}

// setLimits returns a slice of highest power each base can be raised to
func setLimits(factors map[int]int, targetPi int) ([]int, []int, []int) {
	bases := []int{}
	counters := []int{}
	limits := []int{}

	for n, pi := range factors {
		bases = append(bases, n)
		counters = append(counters, 0)
		limit := 1
		tp := targetPi / pi
		for tp%n == 0 {
			limit++
			tp /= n
		}
		limits = append(limits, limit)
	}

	return bases, counters, limits
}

// increment returns true if it has incremented counters by one
func increment(counters, limits []int) bool {
	// Increment least significant digit
	if counters[0] < limits[0] {
		counters[0]++
		return true
	}

	counters[0] = 0

	// Propagate the carry
	for i := 1; i < len(counters); i++ {
		if counters[i] < limits[i] {
			counters[i]++
			return true
		}
		counters[i] = 0
	}

	return false
}

// product returns the product of each base raised to its corresponding power
func product(bases, counters []int) int {
	p := 1
	for i, b := range bases {
		p *= int(math.Pow(float64(b), float64(counters[i])))
	}
	return p
}

// multiples returns a slice of all multiples of the given factors
func multiples(factors map[int]int, targetPi int) []int {
	mul := []int{}

	bases, counters, limits := setLimits(factors, targetPi)

	for increment(counters, limits) {
		n := product(bases, counters)
		mul = append(mul, n)
	}

	return mul
}

func main() {
	fmt.Printf("Welcome to 853\n\n")

	targetPi := 120
	limit := 1000 * 1000 * 1000 // 1000:21966

	// Find all primes < limit that have a period that is
	// a divisor of the targetPi period. (If a prime has a
	// period that is not a divisor of the targetPi period
	// then neither it nor any powers of it can be divisors.)
	primeCandidates := map[int]int{}
	for _, prime := range primes.Primes {
		pi := fibonacciModPeriod(prime, targetPi)
		if pi == -1 {
			continue
		}
		if pi <= targetPi && targetPi%pi == 0 {
			primeCandidates[prime] = pi
		}
		if prime >= limit {
			break
		}
	}

	// Find all values of n that can be made from the
	// prime candidates.
	m := multiples(primeCandidates, targetPi)

	sum := 0

	// For each n < limit, determine whether pi(n) == targetPi
	for _, n := range m {
		if n >= limit {
			continue
		}
		pi := fibonacciModPeriod(n, targetPi)
		if pi == targetPi {
			sum += n
		}
	}

	fmt.Printf("Sum of n where [n < %d and pi(n) == %d]: %d\n\n", limit, targetPi, sum)
}
