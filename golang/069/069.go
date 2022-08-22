package main

import (
	"github.com/erikbryant/project-euler/golang/primes"
	"github.com/erikbryant/project-euler/golang/util"
	"fmt"
)

func init() {
	primes.Load("../primes.gob")
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
		phi := util.Totient(n)
		ratio := float64(n) / float64(phi)
		fmt.Println("n:", n, "ratio:", ratio)
	}
}
