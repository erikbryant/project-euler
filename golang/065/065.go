package main

import (
	"../util"
	"fmt"
	"math/big"
)

func main() {
	numerator, denominator := util.Convergent(100, util.E)
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
