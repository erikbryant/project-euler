package main

import (
	"fmt"
	"strconv"

	"github.com/erikbryant/util-golang/primes"
)

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
	circularCount := 0
	for _, prime := range primes.Primes {
		if circular(prime) {
			fmt.Println("Circular: ", prime)
			circularCount++
		}
	}
	fmt.Println("Circular count: ", circularCount)
}
