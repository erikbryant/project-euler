package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/erikbryant/util-golang/primes"
)

var (
	powPrimes []int
	maxPrime  int
)

func init() {
	for i := 0; i < len(primes.Primes); i++ {
		s := strconv.Itoa(primes.Primes[i])
		powPrimes = append(powPrimes, int(math.Pow10(len(s))))
	}

	maxPrime = primes.Primes[len(primes.Primes)-1]
}

func allCombosPrime(p []int) bool {
	for i := 0; i < len(p); i++ {
		for j := i + 1; j < len(p); j++ {
			p1 := primes.Primes[p[i]]
			p2 := primes.Primes[p[j]]
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
		sum += primes.Primes[p[i]]
	}

	if len(p) == 4 {
		fmt.Println("Sum:", sum, p, []int{primes.Primes[p[0]], primes.Primes[p[1]], primes.Primes[p[2]], primes.Primes[p[3]]})
	}

	if len(p) == 5 {
		fmt.Println("Sum:", sum, p, []int{primes.Primes[p[0]], primes.Primes[p[1]], primes.Primes[p[2]], primes.Primes[p[3]], primes.Primes[p[4]]})
	}

	return sum
}

func findPrimes(max int) {
	for a := 0; primes.Primes[a] <= max; a++ {
		for b := a + 1; primes.Primes[b] <= max; b++ {
			if allCombosPrime([]int{a, b}) {
				for c := b + 1; primes.Primes[c] <= max; c++ {
					if allCombosPrime([]int{a, b, c}) {
						for d := c + 1; primes.Primes[d] <= max; d++ {
							if allCombosPrime([]int{a, b, c, d}) {
								for e := 1; primes.Primes[e] <= max; e++ {
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
