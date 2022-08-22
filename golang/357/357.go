package main

import (
	"github.com/erikbryant/project-euler/golang/primes"
	"fmt"
	"math"
)

func primeDivisors(n int) bool {
	//
	// For d=1 and d=n the equation has the same result:
	//   d=1 -> 1+n/1 = n+1
	//   d=n -> n+n/n = n+1
	// The same holds true for d=2 and d=(n/2). So, we
	// only need to examine the first half of the divisors,
	// since the second half has no new information. This
	// equates to the sqrt(n).
	//
	root := int(math.Sqrt(float64(n)))
	for d := 1; d <= root; d++ {
		if n%d == 0 {
			if !primes.Primes[d+n/d] {
				return false
			}
		}
	}

	return true
}

func main() {
	primes.Load("../primes.gob")

	sum := 0
	for n := 1; n <= 1000*1000*100; n++ {
		if primeDivisors(n) {
			sum += n
		}
	}

	fmt.Println("Sum: ", sum)
}
