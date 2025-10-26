package main

// go fmt ./... && go vet ./... && go test ./... && go build 049.go && time ./049

import (
	"fmt"

	"github.com/erikbryant/util-golang/primes"
)

// The arithmetic sequence, 1487, 4817, 8147, in which each of the terms increases by 3330,
// is unusual in two ways: (i) each of the three terms are prime, and, (ii) each of the
// 4-digit numbers are permutations of one another.
//
// There are no arithmetic sequences made up of three 1-, 2-, or 3-digit primes, exhibiting
// this property, but there is one other 4-digit increasing sequence.
//
// What 12-digit number do you form by concatenating the three terms in this sequence?

func triPrimes(c chan []int32) {
	defer close(c)

	start := 0

	// Find the first 4-digit prime
	for i := 0; primes.Primes[i] <= 999; i++ {
		start = i
	}
	start++

	// Look at each 4-digit prime
	for i := start; primes.Primes[i] <= 9999; i++ {
		for j := i + 1; primes.Primes[j] <= 9999; j++ {
			distance := primes.Primes[j] - primes.Primes[i]
			next := primes.Primes[j] + distance
			if primes.Prime(int(next)) && next <= 9999 {
				c <- []int32{primes.Primes[i], primes.Primes[j], next}
			}
		}
	}
}

func mask(m int32) int32 {
	mask := int32(0)

	for m > 0 {
		shift := m % 10
		mask |= 1 << uint(shift)
		m = m / 10
	}

	return mask
}

func permutations(tri []int32) bool {
	return mask(tri[0]) == mask(tri[1]) && mask(tri[0]) == mask(tri[2])
}

func main() {
	fmt.Printf("Welcome to 049\n\n")

	c := make(chan []int32)
	go triPrimes(c)
	for {
		tri, ok := <-c
		if !ok {
			break
		}
		if permutations(tri) {
			fmt.Println(tri)
		}
	}
}
