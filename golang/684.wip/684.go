package main

// go fmt ./... && go vet ./... && go test && go run 684.go

// Define s(n) to be the smallest number that has a digit sum of n.
// For example s(10) = 19.
//
// Let S(k) = n=1->k∑s(n). You are given S(20) = 1074.
//
// Further let fi be the Fibonacci sequence defined by
// f0 = 0, f1 = 1 and fi = fi-2 + fi-1 for all i ≧ 2.
//
// Find i=2->90∑S(fi). Give your answer modulo 1,000,000,007.

import (
	"fmt"
	"math"
)

var (
	fCache = map[uint64]uint64{}
	SCache = map[uint64]uint64{}
)

// s returns the smallest number that has a digit sum of n
func s(n uint64) uint64 {
	var firstDigit uint64
	var nineCount int
	var result uint64

	firstDigit = n % 9
	nineCount = int(n / 9)
	result = (firstDigit+1)*uint64(math.Pow10(nineCount)) - 1

	return result
}

// S returns n=1->k∑s(n)
func S(k uint64) uint64 {
	var sum uint64

	sum, ok := SCache[k]
	if ok {
		return sum
	}

	if k == 0 {
		return 0
	}
	if k == 1 {
		return s(1)
	}

	SCache[k] = s(k) + S(k-1)
	return SCache[k]
}

func f(i uint64) uint64 {
	var sum uint64

	sum, ok := fCache[i]
	if ok {
		return sum
	}

	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}

	fCache[i] = f(i-2) + f(i-1)
	return fCache[i]
}

func main() {
	fmt.Printf("Welcome to 684\n\n")

	sum := uint64(0)

	for i := uint64(2); i <= 90; i++ {
		sum += S(f(i)) % 1000000007
		sum %= 1000000007
		fmt.Println(i, f(i), sum)
	}

	fmt.Println("Sum %:", sum)
}
