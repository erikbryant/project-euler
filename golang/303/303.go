package main

// go fmt ./... && go vet ./... && go test && go build 303.go && time ./303

import (
	"fmt"
	"log"
	"slices"
)

// For a positive integer n, define f(n) as the least positive multiple of n that, written in base 10,
// uses only digits <= 2.
//
// Thus f(2)=2, f(3)=12, f(7)=21, f(42)=210, f(89)=1121222.
//
//      100
// Also, Σ f(n)/n = 11363107.
//      n=1
//
//    10000
// Find Σ f(n)/n.
//     n=1

var (
	base3 = []uint{}
)

// AllRepeat returns all combinations with repetitions for a given slice,
// from 1 up to a maximum combination length of m.
func AllRepeat[T any](set []T, m int) (subsets [][]T) {
	if m < 1 {
		return nil
	}

	var generateCombos func([]T, int)
	generateCombos = func(current []T, depth int) {
		if depth == 0 {
			subset := make([]T, len(current))
			copy(subset, current)
			subsets = append(subsets, subset)
			return
		}

		for _, item := range set {
			generateCombos(append(current, item), depth-1)
		}
	}

	for length := 1; length <= m; length++ {
		generateCombos([]T{}, length)
	}

	return subsets
}

// digitsToUint returns the uint form of a slice of uint digits
func digitsToUint(digits []uint) uint {
	v := uint(0)
	for _, d := range digits {
		v *= 10
		v += d
	}
	return v
}

// init populates base3 with all 15-digit numbers with digits {0,1,2}
func init() {
	c := AllRepeat([]uint{0, 1, 2}, 15)
	for _, digits := range c {
		n := digitsToUint(digits)
		if n == 0 {
			continue
		}
		base3 = append(base3, n)
	}
	slices.Sort(base3)
}

// multiple returns k where k*n = a number with only digits {0,1,2}
func multiple(n uint) uint {
	// Rule of 9's: one more '12222' than 999, which is one more than 99 ...
	if n == 9999 {
		return 11112222222222222222 / 9999 // 1,111,333,355,557,778
	}

	for _, b3 := range base3 {
		if b3%n == 0 {
			return b3 / n
		}
	}

	log.Fatal("Ran off the end of base3 with n = ", n)
	return 0
}

// computeAll computes all multiples from 1 to upper and sums them
func computeAll(upper uint) uint {
	ks := make([]uint, upper+1)

	sum := uint(0)
	for n := uint(1); n <= upper; n++ {
		//flag := ""
		k := uint(0)
		if ks[n] != 0 {
			k = ks[n]
			//flag = "precomputed"
		} else {
			k = multiple(n)
			// f(5*n) = f(n)*10 = (k*n*10)/(n*5) = k*2
			if n*5 <= upper {
				if (n*5)%10 != 0 {
					ks[n*5] = k * 2
				}
			}
			// f(10*n) = f(n)*10 = (k*n*10)/(n*10) = k
			if n*10 <= upper {
				ks[n*10] = k
			}
		}
		sum += k
		//fmt.Printf("multiple(%5d) = %16d  %20d  sum: %20d  %s\n", n, k, k*n, sum, flag)
	}

	return sum
}

func main() {
	fmt.Printf("Welcome to 303\n\n")

	upper := uint(10 * 1000)
	s := computeAll(upper)
	fmt.Printf("\nsum(1..%d) = %d\n\n", upper, s)
}
