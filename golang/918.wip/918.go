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
	a = []int{0, 1, 2, -5, 4, 17, -10, -17, 8, -47, 3}
)

func A(n int) int {
	// n is even
	if n&0x01 == 0 {
		power := n & -n
		root := n / power
		return power * A(root)
	}

	if n < len(a) {
		return a[n]
	}

	// n is odd
	return A((n-1)/2) - 3*A((n-1)/2+1)
}

func sumToPow2(k int) int {
	if k == 0 {
		return 1
	}
	if k == 1 {
		return 3
	}
	if k == 2 {
		return 2
	}

	sum := 0
	for i := 3; i < k; i++ {
		sum = sum*2 - 4
	}

	return sum
}

func sumBrokenPairs(N int) int {
	// Sum up to 2^k
	k := int(math.Log2(float64(N))) // sum up to 2^k
	sum := sumToPow2(k)
	twoK := 1 << k

	// Every* a(i)=x has a matching a(j)=-x for 4<i<j. These sum to zero.
	// Find any pairs where N falls between i and j. Add a(i) to the sum.
	skip := 4
	offset := 1
	width := 2
	for skip+offset <= N {
		for n := (skip * (N / skip)) + offset; n <= N; n += skip {
			if n+width > N {
				// Broken!
				sum += A(n)
				break
			}
		}
		skip *= 2
		offset *= 2
		width *= 2
	}

	// *There are singleton -5*2^(k-1) values that are not in pairs.
	// These are at n=3, 6, 12, 24, 48, ...
	// If there is one between n=2^k and n=N, add that to the sum.
	s := 3
	for {
		if s > N {
			break
		}
		if s >= twoK {
			sum += -5 * (twoK / 2)
			break
		}
		s *= 2
	}
	return sum
}

func main() {
	fmt.Printf("Welcome to 918\n\n")

	N := 1000 * 1000 * 1000 * 1000
	sum := sumBrokenPairs(N)
	fmt.Printf("sum = %d\n", sum)
}
