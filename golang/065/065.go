package main

import (
	"../primes"
	"fmt"
	"math/big"
)

func init() {
	primes.Load("../primes.gob")
}

// e() returns the nth number (1-based) in the convergent series
// [2; 1,2,1, 1,4,1, 1,6,1, ... ,1,2k,1, ...]
func e(n int) int64 {
	if n == 1 {
		return int64(2)
	}
	if n%3 == 0 {
		return int64(2 * n / 3)
	}
	return int64(1)
}

// convergent() returns the nth convegence of e.
func convergent(n int) (*big.Int, *big.Int) {
	numerator := big.NewInt(e(n))
	denominator := big.NewInt(1)

	for n > 1 {
		// Invert
		denominator, numerator = numerator, denominator

		// Add e(n-1)
		product := big.NewInt(e(n - 1))
		product.Mul(product, denominator)
		numerator.Add(numerator, product)

		n--
	}

	return numerator, denominator
}

func main() {
	numerator, denominator := convergent(100)
	fmt.Println(numerator, denominator)

	// The answer is the sum of the digits in the numerator.
	sum := big.NewInt(0)
	ten := big.NewInt(10)
	temp := big.NewInt(0)
	for numerator.Cmp(big.NewInt(0)) != 0 {
		sum.Add(sum, temp.Mod(numerator, ten))
		numerator.Div(numerator, big.NewInt(10))
	}
	fmt.Println("Sum:", sum)
}
