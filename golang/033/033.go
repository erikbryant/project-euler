package main

// go fmt ./... && go vet ./... && go test ./... && go build 033.go && time ./033

import (
	"fmt"

	"github.com/erikbryant/util-golang/algebra"
)

// The fraction 49/98 is a curious fraction, as an inexperienced mathematician in attempting
// to simplify it may incorrectly believe that 49/98 = 4/8, which is correct, is obtained by
// cancelling the 9s.
//
// We shall consider fractions like, 30/50 = 3/5, to be trivial examples.
//
// There are exactly four non-trivial examples of this type of fraction, less than one in value,
// and containing two digits in the numerator and denominator.
//
// If the product of these four fractions is given in its lowest common terms, find the value
// of the denominator.

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

	n, d := algebra.ReduceFraction(nProduct, dProduct)

	fmt.Printf("\nProduct: %d/%d\n", nProduct, dProduct)
	fmt.Printf("Reduced fraction: %d/%d\n\n", n, d)
}
