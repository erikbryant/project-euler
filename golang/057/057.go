package main

import (
	"../library"
	"fmt"
	"math/big"
)

func main() {
	numeratorLonger := 0
	for i := 1; i <= 1000; i++ {
		numerator, denominator := library.Convergent(i, library.Sqrt2)

		lenN := 0
		lenD := 0
		for numerator.Cmp(big.NewInt(0)) != 0 {
			numerator.Div(numerator, big.NewInt(10))
			lenN++
		}
		for denominator.Cmp(big.NewInt(0)) != 0 {
			denominator.Div(denominator, big.NewInt(10))
			lenD++
		}
		if lenN > lenD {
			numeratorLonger++
		}
	}
	fmt.Println("Longer numerators:", numeratorLonger)
}
