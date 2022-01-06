package main

import (
	"../util"
	"../primes"
	"fmt"
)

func init() {
	primes.Load("../primes.gob")
}

// totient() finds how many numbers k are relatively
// prime to n where 1 <= k < n. Relatively prime
// means that they have no common divisors (other
// than 1). Because of this rule, 1 is considered
// relatively prime to all other numbers.
func totient(n int) int {
	factors := util.Factors(n)

	// 1 is Totient prime to every number.
	count := n - 1

	for i := 2; i < n; i++ {
		if primes.Prime(i) {
			if n%i == 0 {
				count--
			}
			continue
		}
		for _, factor := range factors {
			if i%factor == 0 {
				count--
				break
			}
		}
	}

	return count
}

func main() {
	// Our goal is to find the highest ratio of n/phi.
	// This means that we are looking to minimize phi.
	// Numbers with a minimum phi are ones that have
	// a maximum number of factors. Find each of the
	// numbers that have a maximum set of factors
	// (that is, they are factorials of the first
	// prime numbers).

	n := 1
	for i := 0; n*primes.PackedPrimes[i] <= 1000*1000; i++ {
		n *= primes.PackedPrimes[i]
		phi := totient(n)
		ratio := float64(n) / float64(phi)
		fmt.Println("n:", n, "ratio:", ratio)
	}
}
