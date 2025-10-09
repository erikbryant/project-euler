package main

// go fmt ./... && go vet ./... && go test && go build 214.go && time ./214

import (
	"fmt"

	"github.com/erikbryant/util-golang/algebra"
	"github.com/erikbryant/util-golang/primes"
)

// Let ɸ be Euler's totientCached function, i.e. for a natural number n,
// ɸ(n) is the number of k, 1 <= k <=> n, for which gcd(k, n) = 1.
//
// By iterating ɸ, each positive integer generates a decreasing chain of numbers ending in 1.
// E.g., if we start with 5 the sequence 5,4,2,1 is generated.
// Here is a listing of all chains with length 4:
//
//    5,4,2,1
//    7,6,2,1
//    8,4,2,1
//    9,6,2,1
//   10,4,2,1
//   12,4,2,1
//   14,6,2,1
//   18,6,2,1
//
// Only two of these chains start with a prime, their sum is 12.
//
// What is the sum of all primes less than 40,000,000 which generate a chain of length 25?

// -----------------------------------------v

// My algorithm. Runs in about 74 seconds.
//
// For each prime, check its totient chain length.
// Use caching to avoid duplicated computation.

var (
	totientLenCache = map[int]int{0: 0, 1: 1}
)

// This function becomes quite slow if inlined; forbid inlining
//
//go:noinline
func totient(n int) int {
	factors := 1

	for _, prime := range primes.PackedPrimes {
		if prime > n {
			break
		}
		if n%prime == 0 {
			// This prime is a factor
			factors *= prime - 1
			n /= prime
		}
	}

	return n * factors
}

func totientLen(n int) int {
	l, ok := totientLenCache[n]
	if !ok {
		l = totientLen(totient(n)) + 1
		totientLenCache[n] = l
	}
	return l
}

func sumPrimesSlow(upper, runLen int) (int, int) {
	count := 0
	sum := 0

	for _, prime := range primes.PackedPrimes {
		if prime >= upper {
			break
		}
		l := totientLen(prime-1) + 1
		if l == runLen {
			count++
			sum += prime
		}
	}

	return count, sum
}

// -----------------------------------------^

// -----------------------------------------v

// Algorithm taken from:
// https://projecteuler.net/thread=214;page=2#26361
// Runs in < 1 second
//
// Create a slice of totients
// Sum the prime paths of length 25
//   - if totients[x] == x-1 then x is prime

func pathLen(x int, totients []int) int {
	i := 1
	for x != 1 {
		x = totients[x]
		i += 1
	}
	return i
}

func sumPrimes(upper, runLen int) (int, int) {
	totients := algebra.Totients(upper)

	count := 0
	sum := 0
	for x := 0; x < upper; x++ {
		if totients[x] == x-1 && pathLen(x, totients) == runLen {
			count++
			sum += x
		}
	}

	return count, sum
}

// -----------------------------------------^

func main() {
	fmt.Printf("Welcome to 214\n\n")

	upper := 40 * 1000 * 1000
	runLen := 25

	count, sum := sumPrimes(upper, runLen)
	//count, sum := sumPrimesFast(upper, runLen)

	fmt.Printf("For totient chains of prime numbers < %d with length %d, count = %d sum = %d\n", upper, runLen, count, sum)
}
