package main

import (
	"fmt"
	"github.com/erikbryant/project-euler/golang/primes"
	"strconv"
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
	primes.Init(false)

	circularCount := 0
	for i, prime := range primes.Primes {
		if !prime {
			continue
		}
		if circular(i) {
			fmt.Println("Circular: ", i)
			circularCount++
		}
	}
	fmt.Println("Circular count: ", circularCount)
}
