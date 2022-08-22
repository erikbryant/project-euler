package main

import (
	"github.com/erikbryant/project-euler/golang/primes"
	"fmt"
)

func init() {
	primes.Load("../primes.gob")
}

type fraction struct {
	n int
	d int
}

// Reduce a fraction n/d such that the HCF(n, d) == 1.
func reduce(n, d int) (rN, rD int) {
	for i := 0; primes.PackedPrimes[i] <= n; i++ {
		p := primes.PackedPrimes[i]
		for n%p == 0 && d%p == 0 {
			n = n / p
			d = d / p
		}
	}

	return n, d
}

// For 2 <= d <= max, find all fractions such that: 1/3 < n/d < 1/2
func rpf(max int) int {
	if max <= 4 {
		return 0
	}

	foundFractions := make(map[fraction]int)

	for d := 5; d <= max; d++ {
		start := d/3 + 1
		end := d / 2
		if d%2 == 0 {
			end--
		}

		for n := start; n <= end; n++ {
			rN, rD := reduce(n, d)
			foundFractions[fraction{n: rN, d: rD}] = 1
		}
	}

	return len(foundFractions)
}

func main() {
	fmt.Println("# fractions:", rpf(12000))
}
