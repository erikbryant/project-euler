package main

// go fmt ./... && go vet ./... && go test ./... && go build 060.go && time ./060

import (
	"fmt"
	"math"
	"strconv"

	"github.com/erikbryant/util-golang/primey"
)

// The primes 3, 7, 109, and 673, are quite remarkable. By taking any two primes and
// concatenating them in any order the result will always be prime. For example,
// taking 7 and 109, both 7109 and 1097 are prime. The sum of these four primes, 792,
// represents the lowest sum for a set of four primes with this property.
//
// Find the lowest sum for a set of five primes for which any two primes concatenate
// to produce another prime.

var (
	powPrimes []int
	maxPrime  int
)

func init() {
	for _, prime := range primey.Iter() {
		s := strconv.Itoa(prime)
		powPrimes = append(powPrimes, int(math.Pow10(len(s))))
	}
	maxPrime = primey.PrimeMax()
}

func allCombosPrime(p []int) bool {
	for i := 0; i < len(p); i++ {
		for j := i + 1; j < len(p); j++ {
			p1 := primey.Nth(p[i])
			p2 := primey.Nth(p[j])
			p1Len := powPrimes[p[i]]
			p2Len := powPrimes[p[j]]
			c1 := p1*p2Len + p2
			c2 := p2*p1Len + p1
			if c1 > maxPrime || c2 > maxPrime {
				return false
			}
			if !(primey.Prime(c1) && primey.Prime(c2)) {
				return false
			}
		}
	}
	return true
}

func printPrimes(p []int) int {
	sum := 0

	for i := 0; i < len(p); i++ {
		sum += primey.Nth(p[i])
	}

	if len(p) == 4 {
		fmt.Println("Sum:", sum, p, []int{primey.Nth(p[0]), primey.Nth(p[1]), primey.Nth(p[2]), primey.Nth(p[3])})
	}

	if len(p) == 5 {
		fmt.Println("Sum:", sum, p, []int{primey.Nth(p[0]), primey.Nth(p[1]), primey.Nth(p[2]), primey.Nth(p[3]), primey.Nth(p[4])})
	}

	return sum
}

func findPrimes(maxP int) {
	for a := 0; primey.Nth(a) <= maxP; a++ {
		for b := a + 1; primey.Nth(b) <= maxP; b++ {
			if allCombosPrime([]int{a, b}) {
				for c := b + 1; primey.Nth(c) <= maxP; c++ {
					if allCombosPrime([]int{a, b, c}) {
						for d := c + 1; primey.Nth(d) <= maxP; d++ {
							if allCombosPrime([]int{a, b, c, d}) {
								for e := 1; primey.Nth(e) <= maxP; e++ {
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
	fmt.Printf("Welcome to 060\n\n")

	findPrimes(10000)
}
