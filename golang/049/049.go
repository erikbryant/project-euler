package main

// go fmt ./... && go vet ./... && go test ./... && go build 049.go && time ./049

import (
	"fmt"

	"github.com/erikbryant/util-golang/primey"
)

// The arithmetic sequence, 1487, 4817, 8147, in which each of the terms increases by 3330,
// is unusual in two ways: (i) each of the three terms are prime, and, (ii) each of the
// 4-digit numbers are permutations of one another.
//
// There are no arithmetic sequences made up of three 1-, 2-, or 3-digit primes, exhibiting
// this property, but there is one other 4-digit increasing sequence.
//
// What 12-digit number do you form by concatenating the three terms in this sequence?

func triPrimes(c chan []int) {
	defer close(c)

	// Look at each 4-digit prime
	for i, pi := range primey.Iter() {
		// Only look at 4-digit primes
		if pi < 999 {
			continue
		}
		if pi > 9999 {
			break
		}
		for _, pj := range primey.Iterr(i+1, primey.Len()-1) {
			// Only look at 4-digit primes
			if pj > 9999 {
				break
			}
			distance := pj - pi
			next := pj + distance
			if primey.Prime(next) && next <= 9999 {
				c <- []int{pi, pj, next}
			}
		}
	}
}

func mask(m int) int {
	mask := int(0)

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
	fmt.Printf("Welcome to 049\n\n")

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
