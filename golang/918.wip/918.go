package main

import (
	"fmt"
)

// go fmt ./... && go vet ./... && go test && time go run 918.go

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

func SumCache(N int) (int, int) {
	sum := 0
	stop := len(a) - 1

	a[1] = 1
	sum += a[1]

	n := 2
	for ; n+1 <= stop; n += 2 {
		// n is even
		a[n] = 2 * a[n/2]
		sum += a[n]

		// n+1 is odd
		a[n+1] = a[n/2] - 3*a[n/2+1]
		sum += a[n+1]
	}

	if n <= stop {
		// n is even
		a[n] = 2 * a[n/2]
		sum += a[n]
	}

	return sum, stop
}

func A(n int) int {
	if n < len(a) {
		return a[n]
	}

	// n is even
	if n&0x01 == 0 {
		return 2 * A(n/2)
	}

	// n is odd
	return A((n-1)/2) - 3*A((n-1)/2+1)
}

func SumCalc(start, N int) int {
	sum := 0
	n := start

	if n&0x01 == 1 {
		// first index is odd, manually handle this
		sum += A(n)
		n++
	}

	for ; n+1 <= N; n += 2 {
		an2 := A(n / 2)
		// n is even
		sum += 2 * an2
		// n+1 is odd
		sum += an2 - 3*A(n/2+1)
	}

	if n <= N {
		// n is even
		sum += 2 * A(n/2)
	}

	return sum
}

func main() {
	fmt.Printf("Welcome to 918\n\n")

	N := 1000 * 1000 * 1000 * 10
	a = make([]int, min(N, aMax)+1) // 1-based indexing
	sum, last := SumCache(N)
	fmt.Printf("Sum(%d) = %d [%d/%d]\n", aMax, sum, last, N)
	if last < N {
		sum += SumCalc(last+1, N)
		fmt.Printf("Sum(%d) = %d\n", N, sum)
	}
}
