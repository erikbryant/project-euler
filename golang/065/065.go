package main

import (
	"../library"
	"fmt"
	"math/big"
)

func main() {
	numerator, denominator := library.Convergent(100, library.E)
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
