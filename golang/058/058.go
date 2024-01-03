package main

import (
	"fmt"

	"github.com/erikbryant/util-golang/primes"
)

func diagonals(sideLength int) [4]int {
	d := [4]int{0, 0, 0, 0}

	if sideLength < 3 {
		return d
	}

	// Bottom right
	d[3] = sideLength * sideLength

	// Bottom left
	d[2] = d[3] - sideLength + 1

	// Top left
	d[1] = d[2] - sideLength + 1

	// Top right
	d[0] = d[1] - sideLength + 1

	return d
}

func main() {
	fmt.Println("Welcome to 058")

	numbers := 1
	sideLength := 3
	prime := 0

	for {
		d := diagonals(sideLength)
		numbers += len(d)
		// d[3] is never prime; don't check it.
		for i := 0; i <= 2; i++ {
			if primes.Prime(d[i]) {
				prime++
			}
		}
		primePct := float64(prime) / float64(numbers) * 100.0
		fmt.Printf("Side length: %d  Numbers: %d  Primes: %d (%2f%%)\n", sideLength, numbers, prime, primePct)
		if primePct < 10.0 {
			break
		}
		sideLength += 2
	}
}
