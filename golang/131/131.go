package main

// go fmt ./... && go vet ./... && go test ./... && go build 131.go && time ./131

import (
	"fmt"

	"github.com/erikbryant/util-golang/primey"
)

// There are some prime values, p, for which there exists a positive integer, n,
// such that the expression n^3 + n^2p is a perfect cube.
//
// For example, when p = 19, 8^3 + 8^2×19 = 123.
//
// What is perhaps most surprising is that for each prime with this property the
// value of n is unique, and there are only four such primes below one-hundred.
//
// How many primes below one million have this remarkable property?

// This sequence is https://oeis.org/A002407
// Also can be calculated as the difference of two consecutive cubes

// looper counts the difference between consecutive cubes where the difference
// is prime
func looper(maxP int) int {
	i := 0
	cube := i * i * i
	count := 0

	for {
		prevCube := cube
		i++
		cube = i * i * i
		p := cube - prevCube
		if p >= maxP {
			break
		}
		if primey.Prime(p) {
			fmt.Println(p)
			count++
		}
	}

	return count
}

func main() {
	fmt.Printf("Welcome to 131\n\n")

	maxP := 1000 * 1000
	count := looper(maxP)
	fmt.Printf("\nPerfect cube solutions for n^3 + n^2*p where p < %d = %d\n\n", maxP, count)
}
