package main

// go fmt ./... && go vet ./... && go test && go build 129.go && time ./129

import (
	"fmt"
	"log"

	"github.com/erikbryant/util-golang/algebra"
)

// A number consisting entirely of ones is called a repunit. We shall define R(k) to be a repunit
// of length k; for example, R(6) = 111111.
//
// Given that n is a positive integer and gcd(n, 10) = 1, it can be shown that there always exists
// a value, k, for which R(k) is divisible by n, and let A(n) be the least such value of k;
// for example, A(7) = 6 and A(41) = 5.
//
// The least value of n for which A(n) first exceeds ten is 17.
//
// Find the least value of n for which A(n) first exceeds one-million.

// A returns the lowest value of k for which R(k) is divisible by n where n is a positive integer and gcd(n, 10) = 1
func A(n int) (k int) {
	if n <= 0 || n%2 == 0 || n%5 == 0 {
		log.Fatal("n must be > 0 and not a multiple of {2,5}: ", n)
	}

	// The repunit generator is R(k)=(10^k-1)/9
	// If n is a divisor of R(k) then the remainder of 10^k/9*n has to be 1.
	// Find the k where n divides 10^k with remainder 1.

	for k = 1; algebra.PowerMod(10, k, 9*n) != 1; k++ {
	}

	return k
}

// findFirstKAbove returns the first n,k for which k > upper
func findFirstKAbove(upper int) (int, int) {
	var k int

	// For A(n)=k, k is strictly less than n for all values of n
	n := upper

	for ; ; n++ {
		if n%2 == 0 || n%5 == 0 {
			continue
		}
		k = A(n)
		if k > upper {
			break
		}
	}

	return n, k
}

func main() {
	fmt.Printf("Welcome to 129\n\n")

	upper := 1000 * 1000
	n, k := findFirstKAbove(upper)
	fmt.Printf("The first value of n where A(n) > %d is A(%d) = %d\n", upper, n, k)
}
