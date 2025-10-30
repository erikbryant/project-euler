package main

// go fmt ./... && go vet ./... && go test ./... && go build 060.go && time ./060

import (
	"fmt"
	"math"

	"github.com/erikbryant/util-golang/primey"
)

// The primes 3, 7, 109, and 673, are quite remarkable. By taking any two primes and
// concatenating them in any order the result will always be prime. For example,
// taking 7 and 109, both 7109 and 1097 are prime. The sum of these four primes, 792,
// represents the lowest sum for a set of four primes with this property.
//
// Find the lowest sum for a set of five primes for which any two primes concatenate
// to produce another prime.

func shift(n int) int {
	digits := int(math.Log10(float64(n))) + 1
	return int(math.Pow10(digits))
}

func allCombosPrime(p []int) bool {
	for i := 0; i < len(p); i++ {
		p1 := p[i]
		p1Len := shift(p1)
		for j := i + 1; j < len(p); j++ {
			p2 := p[j]
			p2Len := shift(p2)
			c1 := p1*p2Len + p2
			c2 := p2*p1Len + p1
			if c1 > primey.PrimeMax() || c2 > primey.PrimeMax() {
				return false
			}
			if !primey.Prime(c1) || !primey.Prime(c2) {
				return false
			}
		}
	}
	return true
}

func printPrimes(p []int) int {
	sum := 0

	for i := 0; i < len(p); i++ {
		sum += p[i]
	}

	fmt.Printf("Sum: %6d  %v\n", sum, p)

	return sum
}

func findPrimes(maxP int) []int {
	for ia, a := range primey.Iter() {
		if a > maxP {
			break
		}
		for ib, b := range primey.Iterr(ia+1, primey.Len()-1) {
			if b > maxP {
				break
			}
			if allCombosPrime([]int{a, b}) {
				for ic, c := range primey.Iterr(ia+ib+1, primey.Len()-1) {
					if c > maxP {
						break
					}
					if allCombosPrime([]int{a, b, c}) {
						for id, d := range primey.Iterr(ia+ib+ic+1, primey.Len()-1) {
							if d > maxP {
								break
							}
							if allCombosPrime([]int{a, b, c, d}) {
								for _, e := range primey.Iterr(ia+ib+ic+id+1, primey.Len()-1) {
									if e > maxP {
										break
									}
									p := []int{a, b, c, d, e}
									if allCombosPrime(p) {
										return p
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return nil
}

func main() {
	fmt.Printf("Welcome to 060\n\n")

	p := findPrimes(10000)
	printPrimes(p)
}
