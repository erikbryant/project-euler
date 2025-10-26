package main

// go fmt ./... && go vet ./... && go test ./... && go build 347.go && time ./347

import (
	"fmt"

	"github.com/erikbryant/util-golang/primes"
)

// The largest integer <= 100 that is only divisible by both the primes 2 and 3 is 96, as 96=32x 3=2^5 x 3.
// For two _distinct_ primes p and q let M(p,q,N) be the largest positive integer <= N only divisible by
// both p and q and M(p,q,N)=0 if such a positive integer does not exist.
//
// E.g. M(2,3,100)=96.
// M(3,5,100)=75 and not 90 because 90 is divisible by 2, 3 and 5.
// Also M(2,73,100)=0 because there does not exist a positive integer <= 100 that is divisible by both 2 and 73.
//
// Let S(N) be the sum of all distinct M(p,q,N).
// S(100)=2262.
//
// Find S(10,000,000).

// soloFactors returns true if p and q are the only factors of N
func soloFactors(p, q, N int) bool {
	for N%p == 0 {
		N = N / p
	}

	for N%q == 0 {
		N = N / q
	}

	return N == 1
}

// M returns ...
// For two distinct primes p and q let M(p,q,N) be the
// largest positive integer ≤ N only divisible by both
// p and q and zero if such an integer does not exist.
func M(p, q, N int) int {
	for N > p*q {
		// Shift N down to be a multiple of p*q
		N = N - (N % (p * q))
		if soloFactors(p, q, N) {
			return N
		}
		// Get below this multiple of p and q
		N--
	}
	return 0
}

// S returns the sum of all distinct M(p,q,N)
func S(N int) int {
	sum := 0

	for i, p := range primes.Iter() {
		if p >= N {
			break
		}
		for _, q := range primes.Iterr(i+1, -1) {
			if p*q > N {
				break
			}
			sum += M(p, q, N)
		}
	}

	return sum
}

func main() {
	fmt.Printf("Welcome to 347\n\n")

	N := 1000 * 1000 * 10
	sum := S(N)

	fmt.Printf("S(%d) = %d\n\n", N, sum)
}
