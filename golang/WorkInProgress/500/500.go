package main

import (
	"../library"
	"../primes"
	"fmt"
)

func init() {
	primes.Load("../primes.gob")
}

func factorial(f int) int {
	product := 1
	for i := 2; i <= f; i++ {
		product *= i
	}

	return product
}

func nCk(n, k int) int {
	if k == 1 {
		return n
	}

	if k == n {
		return 1
	}

	return factorial(n) / (factorial(k) * factorial(n-k))
}

func main() {
	count := 0

	fmt.Println("Factors of 500500507:", library.FactorsCounted(500500507))

	for i := 1; i <= 120; i++ {
		if 120%i == 0 {
			fmt.Println(i)
			count++
		}
	}

	fmt.Println("Count:", count)

	// product := new(big.Int)
	// product.SetInt64(1)
	// mod := new(big.Int)
	// mod.SetInt64(500500507)
	// for i := 0; i < 100; i++ {
	// 	fmt.Println(i+1, product, product.Mod(product, mod))
	// 	pp := new(big.Int)
	// 	pp.SetInt64(int64(primes.PackedPrimes[i]))
	// 	product.Mul(product, pp)
	// }
}
