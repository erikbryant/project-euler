package main

import (
	"fmt"
	"github.com/erikbryant/project-euler/golang/primes"
)

func init() {
	primes.Load("../primes.gob")
}

func doesItFactor(p, q, N int) bool {
	if N%p != 0 || N%q != 0 {
		return false
	}

	for N%p == 0 {
		N = N / p
	}

	for N%q == 0 {
		N = N / q
	}

	return N == 1
}

// M does...
// For two distinct primes p and q let M(p,q,N) be the
// largest positive integer ≤N only divisible by both
// p and q and M(p,q,N)=0 if such a positive integer
// does not exist.
func M(p, q, N int) int {
	max := p * q

	if max > N {
		return 0
	}

	pProduct := p
	for pProduct*q <= N {
		pqProduct := pProduct * q
		for pqProduct <= N {
			if pqProduct > max {
				max = pqProduct
			}
			pqProduct = pqProduct * q
		}
		pProduct = pProduct * p
	}

	return max
}

// S returns the sum of all distinct M(p,q,N).
func S(N int) (sum int) {
	for pIndex := 0; primes.PackedPrimes[pIndex] < N; pIndex++ {
		for qIndex := pIndex + 1; primes.PackedPrimes[qIndex]*primes.PackedPrimes[pIndex] <= N; qIndex++ {
			sum += M(primes.PackedPrimes[pIndex], primes.PackedPrimes[qIndex], N)
		}
	}

	return
}

func main() {
	N := 1000 * 1000 * 10
	sum := S(N)
	fmt.Printf("S(%d) = %d\n", N, sum)
}
