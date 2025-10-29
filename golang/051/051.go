package main

// go fmt ./... && go vet ./... && go test ./... && go build 051.go && time ./051

import (
	"fmt"

	"github.com/erikbryant/util-golang/algebra"
	"github.com/erikbryant/util-golang/primey"
)

// By replacing the 1st digit of the 2-digit number *3, it turns out that six of the nine
// possible values: 13, 23, 43, 53, 73, and 83, are all prime.
//
// By replacing the 3rd and 4th digits of 56**3 with the same digit, this 5-digit number
// is the first example having seven primes among the ten generated numbers, yielding the
// family: 56003, 56113, 56333, 56443, 56663, 56773, and 56993. Consequently, 56003, being
// the first member of this family, is the smallest prime with this property.
//
// Find the smallest prime which, by replacing part of the number (not necessarily adjacent digits) with the same digit, is part of an eight prime value family.

// prime() checks to see whether the digits make a prime number.
func prime(digits []int8) bool {
	return primey.Prime(algebra.DigitsToInt(digits))
}

// replacements() tries each of the 0-9 variations for a single digits/common pair.
func replacements(digits []int8, common []int) int {
	familyLen := 0

	// The digit to try.
	for d := int8(0); d <= 9; d++ {
		tester := make([]int8, len(digits))
		copy(tester, digits)
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

	tmp := make([]int, len(result))
	copy(tmp, result)
	c <- tmp

	for {
		i := r - 1
		for ; i >= 0 && indices[i] == i+n-r; i-- {
		}

		if i < 0 {
			return
		}

		indices[i]++
		for j := i + 1; j < r; j++ {
			indices[j] = indices[j-1] + 1
		}

		for ; i < len(indices); i++ {
			result[i] = pool[indices[i]]
		}
		tmp := make([]int, len(result))
		copy(tmp, result)
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
		//tmp := copySlice(list)
		//c <- tmp
	}()

	return c
}

// findCommon() finds each set of matching digits and returns the positions in a list.
func findCommon(digits []int8) <-chan []int {
	c := make(chan []int)

	go func() {
		defer close(c)
		repeats := make(map[int][]int)

		for i := 0; i <= 9; i++ {
			repeats[i] = make([]int, 0)
		}

		for i := 0; i < len(digits); i++ {
			repeats[int(digits[i])] = append(repeats[int(digits[i])], i)
		}

		for _, repeat := range repeats {
			if len(repeat) > 0 {
				for r := range combinations(repeat) {
					tmp := make([]int, len(r))
					copy(tmp, r)
					c <- tmp
				}
			}
		}
	}()

	return c
}

func main() {
	for _, n := range primey.Iter() {
		if n > 999999 {
			break
		}
		digits := algebra.IntToDigits(n)
		for common := range findCommon(digits) {
			familyLen := replacements(digits, common)
			if familyLen >= 8 {
				fmt.Println(n, common, familyLen)
			}
		}
	}
}
