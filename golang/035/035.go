package main

// go fmt ./... && go vet ./... && go test ./... && go build 035.go && time ./035

import (
	"fmt"
	"strconv"

	"github.com/erikbryant/util-golang/primes"
)

// The number, 197, is called a circular prime because all rotations of the
// digits: 197, 971, and 719, are themselves prime.
//
// There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.
//
// How many circular primes are there below one million?

func rotate(s string) string {
	return s[1:] + string(s[0])
}

func circular(number int) bool {
	if !primes.Prime(number) {
		return false
	}
	digits := strconv.Itoa(number)

	for i := 0; i < len(digits); i++ {
		digits = rotate(digits)
		n, _ := strconv.Atoi(digits)
		if !primes.Prime(n) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Printf("Welcome to 035\n\n")

	circularCount := 0
	for _, prime := range primes.Iter() {
		if circular(prime) {
			fmt.Println("Circular: ", prime)
			circularCount++
		}
	}

	fmt.Printf("\nCircular prime count < 1,000,000: %d\n\n", circularCount)
}
