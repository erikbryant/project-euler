package main

// go fmt ./... && go vet ./... && go test ./... && go build 357.go && time ./357

import (
	"fmt"
	"math"

	"github.com/erikbryant/util-golang/primey"
)

// Consider the divisors of 30: 1,2,3,5,6,10,15,30.
// It can be seen that for every divisor d of 30, d + 30 / d is prime.
//
// Find the sum of all positive integers n not exceeding 100,000,000
// such that for every divisor d of n, d + n / d is prime.

func primeDivisors(n int) bool {
	//
	// For d=1 and d=n the equation has the same result:
	//   d=1 -> 1+n/1 = n+1
	//   d=n -> n+n/n = n+1
	// The same holds true for d=2 and d=(n/2). So, we
	// only need to examine the first half of the divisors,
	// since the second half has no new information. This
	// equates to the sqrt(n).
	//
	root := int(math.Sqrt(float64(n)))
	for d := 1; d <= root; d++ {
		if n%d == 0 {
			if !primey.Prime(d + n/d) {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Printf("Welcome to 357\n\n")

	sum := 0
	for n := 1; n <= 1000*1000*100; n++ {
		if primeDivisors(n) {
			sum += n
		}
	}

	fmt.Printf("Sum of n where n <= 100,000,000 and f(n) is prime: %d\n\n", sum)
}
