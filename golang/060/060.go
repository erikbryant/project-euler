package main

import (
	"fmt"
	"github.com/erikbryant/project-euler/golang/primes"
	"math"
	"strconv"
)

var (
	powPrimes []int
	maxPrime  int
)

func init() {
	primes.Load("../primes.gob")

	for i := 0; i < len(primes.PackedPrimes); i++ {
		s := strconv.Itoa(primes.PackedPrimes[i])
		powPrimes = append(powPrimes, int(math.Pow10(len(s))))
	}

	maxPrime = primes.PackedPrimes[len(primes.PackedPrimes)-1]
}

func allCombosPrime(p []int) bool {
	for i := 0; i < len(p); i++ {
		for j := i + 1; j < len(p); j++ {
			p1 := primes.PackedPrimes[p[i]]
			p2 := primes.PackedPrimes[p[j]]
			p1Len := powPrimes[p[i]]
			p2Len := powPrimes[p[j]]
			c1 := p1*p2Len + p2
			c2 := p2*p1Len + p1
			if c1 > maxPrime || c2 > maxPrime {
				return false
			}
			if !(primes.Prime(c1) && primes.Prime(c2)) {
				return false
			}
		}
	}
	return true
}

func printPrimes(p []int) int {
	sum := 0

	for i := 0; i < len(p); i++ {
		sum += primes.PackedPrimes[p[i]]
	}

	if len(p) == 4 {
		fmt.Println("Sum:", sum, p, []int{primes.PackedPrimes[p[0]], primes.PackedPrimes[p[1]], primes.PackedPrimes[p[2]], primes.PackedPrimes[p[3]]})
	}

	if len(p) == 5 {
		fmt.Println("Sum:", sum, p, []int{primes.PackedPrimes[p[0]], primes.PackedPrimes[p[1]], primes.PackedPrimes[p[2]], primes.PackedPrimes[p[3]], primes.PackedPrimes[p[4]]})
	}

	return sum
}

func findPrimes(max int) {
	for a := 0; primes.PackedPrimes[a] <= max; a++ {
		for b := a + 1; primes.PackedPrimes[b] <= max; b++ {
			if allCombosPrime([]int{a, b}) {
				for c := b + 1; primes.PackedPrimes[c] <= max; c++ {
					if allCombosPrime([]int{a, b, c}) {
						for d := c + 1; primes.PackedPrimes[d] <= max; d++ {
							if allCombosPrime([]int{a, b, c, d}) {
								for e := 1; primes.PackedPrimes[e] <= max; e++ {
									p := []int{a, b, c, d, e}
									if allCombosPrime(p) {
										printPrimes(p)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func main() {
	findPrimes(10000)
}
