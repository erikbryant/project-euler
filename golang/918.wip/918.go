package main

import (
	"fmt"
	"math"
)

// go fmt ./... && go vet ./... && go test && go build 918.go && time ./918

// The sequence an is defined by a(1) = 1, and then recursively for n >= 1:
// a(2n) = 2a(n)
// a(2n+1) = a(n) - 3a(n+1)
// The first ten terms are 1, 2, -5, 4, 17, -10, -17, 8, -47, 34.
// Define S(N) = Σn=1->N a(n). You are given S(10) = -13.
// Find S(10^12).

var (
	aMax = 1000 * 1000 * 1000 * 2
	a    = []int{}
)

func CacheFast(N int) {
	stop := len(a) - 1

	a[1] = 1

	n := 2
	for ; n+1 <= stop; n += 2 {
		// n is even
		a[n] = 2 * a[n/2]

		// n+1 is odd
		a[n+1] = a[n/2] - 3*a[n/2+1]
	}

	if n <= stop {
		// n is even
		a[n] = 2 * a[n/2]
	}
}

func A(n int) int {
	// n is even
	if n&0x01 == 0 {
		return 2 * A(n/2)
	}

	if n < len(a) {
		return a[n]
	}

	// n is odd
	return A((n-1)/2) - 3*A((n-1)/2+1)
}

// sumPow2 returns the sum of powers of 2 from 2^0 to 2^k
func sumPow2(k int) int {
	return (1 << (k + 1)) - 1
}

// sumFives returns the sum of a(1) to a(2^k) (except where n is a power of 2)
func sumFives(k int) int {
	sum := 0
	for i := 1; i < k; i++ {
		sum += -5 * (1 << (i - 1))
	}
	return sum
}

func sumToPow2(k int) int {
	sum := 0

	// Sum of powers of 2 from 2^0 to 2^k
	sum += sumPow2(k)

	// Sum of non-zero component of each 2^i to 2^(i+1) span
	sum += sumFives(k)

	return sum
}

func main() {
	fmt.Printf("Welcome to 918\n\n")

	N := 1000 * 1000 * 1000 * 10
	k := int(math.Log2(float64(N))) // sum up to 2^k

	sum := sumToPow2(k)
	fmt.Printf("sum from a(1) to a(%d) = %d\n", 1<<k, sum)

	a = make([]int, min(N, aMax)+1) // 1-based indexing
	CacheFast(N)

	for i := (1 << k) + 1; i <= N; i++ {
		sum += A(i)
	}

	fmt.Printf("sum = %d\n", sum)
}
