package main

import (
	"github.com/erikbryant/project-euler/golang/primes"
	"fmt"
)

func init() {
	primes.Load("../primes.gob")
}

func makeFractions(c chan []int) {
	defer close(c)

	for denominator := 11; denominator <= 99; denominator++ {
		if denominator%10 == 0 {
			continue
		}
		for numerator := 11; numerator < denominator; numerator++ {
			if numerator%10 == 0 {
				continue
			}
			c <- []int{numerator, denominator}
		}
	}
}

func cancels(fraction []int) bool {
	numerator := fraction[0]
	denominator := fraction[1]
	div := float64(numerator) / float64(denominator)

	if numerator%10 == 0 || denominator%10 == 0 {
		return false
	}

	nDigit1 := numerator / 10
	dDigit1 := denominator / 10
	nDigit2 := numerator % 10
	dDigit2 := denominator % 10

	if nDigit1 == dDigit1 {
		return float64(nDigit2)/float64(dDigit2) == div
	}
	if nDigit1 == dDigit2 {
		return float64(nDigit2)/float64(dDigit1) == div
	}
	if nDigit2 == dDigit1 {
		return float64(nDigit1)/float64(dDigit2) == div
	}
	if nDigit2 == dDigit2 {
		return float64(nDigit1)/float64(dDigit1) == div
	}

	return false
}

func reduce(fraction []int) []int {
	numerator := fraction[0]
	denominator := fraction[1]

	for i := 0; primes.PackedPrimes[i] <= numerator; i++ {
		p := primes.PackedPrimes[i]
		for numerator%p == 0 && denominator%p == 0 {
			numerator = numerator / p
			denominator = denominator / p
		}
	}

	return []int{numerator, denominator}
}

func main() {
	fmt.Println("Welcome to 033")
	c := make(chan []int, 100)
	nProduct := 1
	dProduct := 1

	go makeFractions(c)

	for {
		fraction, ok := <-c
		if !ok {
			break
		}
		if cancels(fraction) {
			nProduct *= fraction[0]
			dProduct *= fraction[1]
			fmt.Println(fraction)
		}
	}
	fmt.Println("Products:", nProduct, dProduct)
}
