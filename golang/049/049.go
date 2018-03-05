package main

import (
	"../primes"
	"fmt"
)

func init() {
	primes.Load("../primes.gob")
}

func triPrimes(c chan []int) {
	defer close(c)

	start := 0

	// Find the first 4-digit prime
	for i := 0; primes.PackedPrimes[i] <= 999; i++ {
		start = i
	}
	start++

	// Look at each 4-digit prime
	for i := start; primes.PackedPrimes[i] <= 9999; i++ {
		for j := i + 1; primes.PackedPrimes[j] <= 9999; j++ {
			distance := primes.PackedPrimes[j] - primes.PackedPrimes[i]
			next := primes.PackedPrimes[j] + distance
			if primes.Prime(next) && next <= 9999 {
				c <- []int{primes.PackedPrimes[i], primes.PackedPrimes[j], next}
			}
		}
	}
}

func mask(m int) int {
	mask := 0

	for m > 0 {
		shift := m % 10
		mask |= 1 << uint(shift)
		m = m / 10
	}

	return mask
}

func permutations(tri []int) bool {
	return mask(tri[0]) == mask(tri[1]) && mask(tri[0]) == mask(tri[2])
}

func main() {
	c := make(chan []int)
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
