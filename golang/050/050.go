package main

// go fmt ./... && go vet ./... && go test ./... && go build 050.go && time ./050

import (
	"fmt"

	"github.com/erikbryant/util-golang/primes"
)

// The prime 41, can be written as the sum of six consecutive primes:
// 41 = 2 + 3 + 5 + 7 + 11 + 13.
//
// This is the longest sum of consecutive primes that adds to a prime below one-hundred.
//
// The longest sum of consecutive primes below one-thousand that adds to a prime,
// contains 21 terms, and is equal to 953.
//
// Which prime, below one-million, can be written as the sum of the most consecutive primes?

func main() {
	fmt.Printf("Welcome to 050\n\n")

	maxPrime := 0
	maxCount := 0

	for start := 0; start < 10000; start++ {
		sum := 0
		i := start
		for {
			if sum+int(primes.Primes[i]) >= 1000000 {
				break
			}
			sum += int(primes.Primes[i])

			if primes.Prime(sum) {
				if i-start > maxCount {
					maxCount = i - start
					maxPrime = sum
				}
			}

			i++
		}
	}

	fmt.Printf("\nLargest prime < 1,000,000 that is the sum of consecutive primes: %d  (%d terms in its sum)\n\n", maxPrime, maxCount)
}
