package main

import (
	"fmt"

	"github.com/erikbryant/util-golang/primes"
)

func main() {
	maxPrime := 0
	maxCount := 0
	for start := 0; start < 10000; start++ {
		sum := 0
		i := start
		for {
			if sum+primes.Primes[i] >= 1000000 {
				break
			}
			sum += primes.Primes[i]

			if primes.Prime(sum) {
				if i-start > maxCount {
					maxCount = i - start
					maxPrime = sum
				}
			}

			i++
		}
	}
	fmt.Println("Max Prime: ", maxPrime, " max count: ", maxCount)
}
