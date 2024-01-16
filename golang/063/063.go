package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to 063")

	count := 0

	// 10^x is the cut over to where val grows faster than
	// the exponent, so no need to go that high on the base
	for i := 1; i < 10; i++ {
		val := 1.0
		for pow := 1; pow <= 25; pow++ {
			val *= float64(i)
			s := fmt.Sprintf("%0.0f", val)
			if len(s) == pow {
				count++
				fmt.Printf("%d^%d = %s   count = %d\n", i, pow, s, count)
			}
		}
	}
}
