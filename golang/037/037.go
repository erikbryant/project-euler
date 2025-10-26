package main

// go fmt ./... && go vet ./... && go test ./... && go build 037.go && time ./037

import (
	"fmt"
	"strconv"

	"github.com/erikbryant/util-golang/primes"
)

// The number 3797 has an interesting property. Being prime itself, it is
// possible to continuously remove digits from left to right, and remain
// prime at each stage: 3797, 797, 97, and 7. Similarly, we can work from
// right to left: 3797, 379, 37, and 3.
//
// Find the sum of the only eleven primes that are both truncatable from
// left to right and right to left.
//
// NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.

func truncate(p string) bool {
	if len(p) < 2 {
		return false
	}
	s := p
	for len(s) > 0 {
		n, _ := strconv.Atoi(s)
		if !primes.Prime(n) {
			return false
		}
		s = s[1:]
	}

	s = p
	for len(s) > 0 {
		n, _ := strconv.Atoi(s)
		if !primes.Prime(n) {
			return false
		}
		s = s[:len(s)-1]
	}

	return true
}

func main() {
	fmt.Printf("Welcome to 037\n\n")

	count := 0
	sum := 0

	for _, prime := range primes.Iter() {
		s := strconv.Itoa(prime)
		if truncate(s) {
			fmt.Println(s)
			count++
			n, _ := strconv.Atoi(s)
			sum += n
		}
	}

	fmt.Println()
	fmt.Println("Count: ", count)
	fmt.Println("Sum  : ", sum)
	fmt.Println()
}
