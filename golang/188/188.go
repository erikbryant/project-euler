package main

// go fmt ./... && go vet ./... && go test && go build 188.go && time ./188

import (
	"fmt"
)

// The hyperexponentiation or tetration of a number a by a positive integer b,
// denoted by a^^b or $^b a$, is recursively defined by:
//
//   a^^1 = a,
//   a^^(k+1) = a^(a^^k).
//
// Thus we have e.g. 3^^2 = 3^3 = 27, hence 3^^3 = 3^27 = 7625597484987 and
// 3^^4 is roughly 10^3.6383346400240996x10^12.
//
// Find the last 8 digits of 1777^^1855.

//func tetration(a, exp int) int {
//	if exp == 1 {
//		return a
//	}
//
//	return tetration(a, exp-1)
//}

// initProducts returns one full period of a^1 .. a^n
func initProducts(a, mask int) []int {
	product := a
	products := []int{product}

	for {
		product *= a
		product %= mask
		if product == a {
			break
		}
		products = append(products, product)
	}

	return products
}

// exp returns the value a^exp by looking it up in an offset table
func exp(a, exp int, products []int) int {
	offset := exp%len(products) - 1
	return products[offset]
}

func main() {
	fmt.Printf("Welcome to 188\n\n")

	a := 1777
	exponent := 1855
	mask := 1000 * 1000 * 100
	products := initProducts(a, mask)

	// For each level in the tetration, look up the resulting exponent for the next level down
	e := a
	for i := 1; i < exponent; i++ {
		e = exp(a, e, products)
	}

	fmt.Printf("Last 8 digits of %d^^%d = %d\n", a, exponent, e)
}
