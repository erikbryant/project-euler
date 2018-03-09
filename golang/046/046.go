package main

import (
	"../primes"
	"fmt"
)

func init() {
	primes.Load("../primes.gob")
}

// It was proposed by Christian Goldbach that every odd composite number
// can be written as the sum of a prime and twice a square.
//
//   9  = 7 + 2×1^2
//   15 = 7 + 2×2^2
//   21 = 3 + 2×3^2
//
// If the number can, return that prime and that square. If not, return 0, 0.
func goldbach(composite int) (prime, square int) {
	i := 1
	for {
		s := 2 * i * i
		if s >= composite {
			break
		}
		p := composite - s
		if primes.Prime(p) {
			return p, i
		}
		i++
	}

	return 0, 0
}

func oddComposite(c chan int) {
	defer close(c)

	// Return the odd, non-prime numbers
	for i := 3; i <= 1000*1000; i += 2 {
		if !primes.Prime(i) {
			c <- i
		}
	}
}

func main() {
	fmt.Println("Welcome to 046")
	c := make(chan int)
	go oddComposite(c)

	for {
		composite, ok := <-c
		if !ok {
			break
		}
		prime, square := goldbach(composite)
		if prime == 0 && square == 0 {
			fmt.Println("Composite is not Goldbachian:", composite)
			break
		}
	}
}
