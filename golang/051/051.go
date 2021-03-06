package main

import (
	"../primes"
	"fmt"
	"math"
)

func init() {
	primes.Load("../primes.gob")
}

// prime() checks to see whether the digits make a prime number.
func prime(digits []int) bool {
	return primes.Prime(digitsToNumber(digits))
}

// reverse() reverses the order of the eements in a list.
func reverse(digits []int) []int {
	rev := make([]int, 0)

	for i := len(digits) - 1; i >= 0; i-- {
		rev = append(rev, digits[i])
	}

	return rev
}

// numberToDigits() converts an int into a list of its component digits.
func numberToDigits(n int) []int {
	digits := make([]int, 0)

	for n > 0 {
		d := n % 10
		digits = append(digits, d)
		n = n / 10
	}

	return reverse(digits)
}

// digitsToNumber() converts a list of digits to a single int.
func digitsToNumber(digits []int) int {
	number := 0

	for i := 0; i < len(digits); i++ {
		number += digits[i] * int(math.Pow(10.0, float64(len(digits)-1-i)))
	}

	return number
}

// copy() returns a copy of the list.
func copy(a []int) []int {
	b := make([]int, 0)
	for i := 0; i < len(a); i++ {
		b = append(b, a[i])
	}
	return b
}

// replacements() tries each of the 0-9 variations for a single digits/common pair.
func replacements(digits []int, common []int) int {
	familyLen := 0

	// The digit to try.
	for d := 0; d <= 9; d++ {
		tester := copy(digits)
		// The position(s) in which to try it.
		for i := 0; i < len(common); i++ {
			pos := common[i]
			tester[pos] = d
		}
		if tester[0] != 0 && prime(tester) {
			familyLen++
		}
	}

	return familyLen
}

// Written by Nuno Antunes, 2012-08-08
// https://play.golang.org/p/JEgfXR2zSH
// GitHub: https://github.com/ntns
func combinationsX(iterable []int, r int, c chan []int) {
	pool := iterable
	n := len(pool)

	if r > n {
		return
	}

	indices := make([]int, r)
	for i := range indices {
		indices[i] = i
	}

	result := make([]int, r)
	for i, el := range indices {
		result[i] = pool[el]
	}

	tmp := copy(result)
	c <- tmp

	for {
		i := r - 1
		for ; i >= 0 && indices[i] == i+n-r; i -= 1 {
		}

		if i < 0 {
			return
		}

		indices[i] += 1
		for j := i + 1; j < r; j += 1 {
			indices[j] = indices[j-1] + 1
		}

		for ; i < len(indices); i += 1 {
			result[i] = pool[indices[i]]
		}
		tmp := copy(result)
		c <- tmp
	}
}

func combinations(list []int) <-chan []int {
	c := make(chan []int)

	go func() {
		defer close(c)

		for i := 1; i <= len(list); i++ {
			combinationsX(list, i, c)
		}
		//tmp := copy(list)
		//c <- tmp
	}()

	return c
}

// findCommon() finds each set of matching digits and returns the positions in a list.
func findCommon(digits []int) <-chan []int {
	c := make(chan []int)

	go func() {
		defer close(c)
		repeats := make(map[int][]int)

		for i := 0; i <= 9; i++ {
			repeats[i] = make([]int, 0)
		}

		for i := 0; i < len(digits); i++ {
			repeats[digits[i]] = append(repeats[digits[i]], i)
		}

		for _, repeat := range repeats {
			if len(repeat) > 0 {
				for r := range combinations(repeat) {
					tmp := copy(r)
					c <- tmp
				}
			}
		}
	}()

	return c
}

func main() {
	for i := 0; i < len(primes.PackedPrimes); i++ {
		n := primes.PackedPrimes[i]
		if n > 999999 {
			break
		}
		digits := numberToDigits(n)
		for common := range findCommon(digits) {
			familyLen := replacements(digits, common)
			if familyLen >= 8 {
				fmt.Println(n, common, familyLen)
			}
		}
	}
}
