package main

import (
	"../primes"
	"fmt"
)

func init() {
	primes.Load("../primes.gob")
}

// digitSum returns the sum of the digits in the number.
func digitSum(c int) int {
	sum := 0

	for c > 0 {
		sum += c % 10
		c /= 10
	}

	return sum
}

// harshad returns true if c is divisible by the sum of its digits.
func harshad(c int) bool {
	return c != 0 && c%digitSum(c) == 0
}

// rightTruncatableHarshad returns true if c is prime and is right truncatable.
func rightTruncatableHarshad(c int) bool {
	if c < 10 {
		return false
	}

	if !harshad(c) {
		return false
	}

	if c < 100 {
		return true
	}

	return rightTruncatableHarshad(c / 10)
}

// strong returns true if c divided by the sum of its digits is prime.
func strong(c int) bool {
	return c != 0 && primes.Prime(c/digitSum(c))
}

// strongRightTruncatableHarshad returns true if c is strong and is right truncatable.
func strongRightTruncatableHarshad(c int) bool {
	return strong(c) && rightTruncatableHarshad(c)
}

// strongRightTruncatableHarshadPrime returns true if p is prime and the first truncation is a strong right truncatable Harshad.
func strongRightTruncatableHarshadPrime(p int) bool {
	return primes.Prime(p) && strongRightTruncatableHarshad(p/10)
}

func sumSRTHP(max int) int {
	sum := 0

	for _, p := range primes.PackedPrimes {
		if p > max {
			break
		}
		if strongRightTruncatableHarshad(p / 10) {
			fmt.Println(p)
			sum += p
		}
	}

	fmt.Println("---Switching to slow primes---")

	for p := primes.PackedPrimes[primes.PackedPrimesEnd] + 2; p <= max; p += 2 {
		if strongRightTruncatableHarshadPrime(p) {
			fmt.Println(p)
			sum += p
		}
	}

	return sum
}

func main() {
	fmt.Printf("Welcome to 387\n\n")

	fmt.Printf("Initialized with primes up to: %d\n\n", primes.PackedPrimes[primes.PackedPrimesEnd])

	fmt.Println("Sum: ", sumSRTHP(100*1000*1000*1000*1000))
}
