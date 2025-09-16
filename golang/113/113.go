package main

// go fmt ./... && go vet ./... && go test && go build 113.go && time ./113

// Working from left-to-right if no digit is exceeded by the digit to its left
// it is called an decreasing number; for example, 134468.
//
// Similarly, if no digit is exceeded by the digit to its right it is called a
// decreasing number; for example, 66420.
//
// We shall call a positive integer that is neither decreasing nor decreasing
// a "bouncy" number; for example, 155349.
//
// As n increases, the proportion of bouncy numbers below n increases such that
// there are only 12951 numbers below one-million that are not bouncy and only
// 277032 non-bouncy numbers below 10^10.
//
// How many numbers below a googol (10^100) are not bouncy?

// Note: zero is not included in this counting.

import (
	"fmt"
	"math"
)

// eval returns flat: 0, decreasing: 1, increasing: 2, bouncy: 3
func eval(n int64) int {
	x := n

	if x < 10 {
		return 0
	}

	flat := true
	increasing := true
	decreasing := true

	prev := x % 10
	x /= 10
	for x > 0 {
		digit := x % 10
		if digit > prev {
			increasing = false
			flat = false
		}
		if digit < prev {
			decreasing = false
			flat = false
		}
		x /= 10
		prev = digit
	}

	if flat {
		return 0
	}
	if decreasing {
		return 1
	}
	if increasing {
		return 2
	}
	return 3 // bouncy
}

// flat returns the count of flat (non-increasing non-decreasing non-bouncy) numbers below 10^digits
func flat(digits int) int64 {
	return int64(9 * digits)
}

// decreasing returns the count of decreasing numbers below 10^digits
func decreasing(digits int) int64 {
	if digits == 1 {
		return 0
	}

	counts := [][10]int64{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, // Single digit numbers
	}

	for d := int64(1); d < int64(digits); d++ {
		this := [10]int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		for j := int64(0); j < 10; j++ {
			for k := int64(0); k <= j; k++ {
				this[j] += counts[d-1][k]
			}
		}
		counts = append(counts, this)
	}

	count := int64(0)
	for _, c := range counts {
		for _, val := range c {
			count += val - 1
		}
	}

	return count
}

// increasing returns the count of increasing numbers below 10^digits
func increasing(digits int) int64 {
	if digits == 1 {
		return 0
	}

	counts := [][10]int64{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, // Single digit numbers
	}

	for d := int64(1); d < int64(digits); d++ {
		this := [10]int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		for j := int64(0); j < 10; j++ {
			for k := int64(1); k <= j; k++ {
				this[j] += counts[d-1][k]
			}
		}
		counts = append(counts, this)
	}

	count := int64(0)
	for _, c := range counts {
		for _, val := range c {
			count += val - 1
		}
	}

	return count + (int64(digits) - 1)
}

func main() {
	fmt.Printf("Welcome to 113\n\n")

	Digits := 100

	if Digits < 10 {
		// Brute force check all numbers to confirm
		counters := [4]int64{}
		N := int64(math.Pow10(Digits))
		for i := int64(1); i < N; i++ {
			e := eval(i)
			counters[e]++
		}
		fmt.Printf("Digits: %d  flat: %d  decreasing: %d  increasing: %d  bouncy: %d\n", Digits, counters[0], counters[1], counters[2], counters[3])
	}

	d := decreasing(Digits)
	f := flat(Digits)
	i := increasing(Digits)
	fmt.Printf("Digits: %d  flat: %d  decreasing: %d  increasing: %d\n", Digits, f, d, i)
	fmt.Printf("\nDigits below 10^%d that are not bouncy: %d\n", Digits, f+d+i)
}
