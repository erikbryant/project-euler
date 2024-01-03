package main

import (
	"fmt"

	"github.com/erikbryant/util-golang/primes"
	"github.com/erikbryant/util-golang/util"
)

// seive() Implements the seive of Eranthoses using an array of counters. It identifies
// which numbers are divisible by the prime factors that make up product and which are
// not.
//
//	1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29
//
// -2:   1 2 3   5   7   9    11    13    15    17    19    21    23    25    27    29
// -3:   1 2 3   5   7        11    13          17    19          23    25          29
// -5:   1 2 3       7        11    13          17    19          23                29
func seive(product int) int {
	f := util.Factors(product)
	counters := make([]int, len(f))
	target := 15499.0 / 94744.0

	saved := 0
	i := 1
	for {
		keep := true
		// Increment each counter one tick.
		for c := 0; c < len(counters); c++ {
			counters[c]++
			if counters[c] == f[c] {
				// If any counter is zero, delete this number.
				counters[c] = 0
				keep = false
			}
		}
		if keep {
			saved++
		}
		ratio := float64(saved) / float64(i-1)
		if ratio < target {
			// Only values of i that are multiples of all of the
			// factors can count. Otherwise, we would have to
			// refactor and start over with counting what gets
			// saved.
			candidate := true
			for j := 0; j < len(f); j++ {
				if i%f[j] != 0 {
					candidate = false
					break
				}
			}
			if candidate {
				fmt.Println("i:", i, "saved:", saved, "ratio:", ratio, "target:", target)
				return saved
			}
		}
		i++
	}
}

func main() {
	fmt.Println("Welcome to 243. Home of the", 15499.0/94744.0)

	// It turns out that 9 is a good low point and 10
	// is too far out. Start at 9 and work upwards
	// from there.
	n := 9

	// Find the product of the first n primes.
	product := 1
	i := 0
	for i < n {
		product *= primes.PackedPrimes[i]
		i++
	}
	seive(product)
}
