package main

import (
	"fmt"
	"strconv"

	"github.com/erikbryant/util-golang/primes"
)

// S returns the result of: For a prime p let S(p) = (∑(p-k)!) mod(p) for 1 ≤ k ≤ 5.
func S(p int) int {
	sum := 0

	// (p-1)!%p == p-1
	sum += p - 1

	// (p-2)!%p == 1
	sum++

	// (p-3)!%p == (p-1)/2
	sum += (p - 1) / 2

	// (p-4)!%p ==
	four := (p + 1) / 6
	if (p+1)%6 != 0 {
		four = p - four
	}
	sum += four

	// (p-5)!%p ==
	var five int
	switch p % 24 {
	case 1:
		five = (p - 1) / 24
	case 5:
		five = p/4 - p/24
	case 7:
		five = p/3 - p/24
	case 11:
		five = p/2 - p/24
	case 13:
		five = (p+1)/2 + (p-1)/24
	case 17:
		five = p - p/3 + p/24
	case 19:
		five = p - p/4 + p/24
	case 23:
		five = (p - 1) - (p-1)/24
	default:
		panic("Yikes! " + strconv.Itoa(p%24))
	}
	sum += five

	return sum % p
}

func sumS(min, max int) int {
	sum := 0

	for i := min; i < max; i++ {
		if primes.Prime(i) {
			sum += S(i)
		}
	}

	return sum
}

// Find ∑S(p) for 5 ≤ p < 10^8.
func main() {
	fmt.Println("Welcome to 381")
	fmt.Println(sumS(5, 1000*1000*100))
}
