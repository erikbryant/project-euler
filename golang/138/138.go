package main

// go fmt ./... && go vet ./... && go test && go build 138.go && time ./138

import (
	"fmt"
	"math/big"
)

// Consider the isosceles triangle with base length, b = 16, and legs, L = 17.
//
// https://projecteuler.net/resources/images/0138.png?1678992052
//
// By using the Pythagorean theorem it can be seen that the height of the triangle,
// h = sqrt(17^2 - 8^2) = 15, which is one less than the base length.
//
// With b = 272 and L = 305, we get h = 273, which is one more than the base length,
// and this is the second-smallest isosceles triangle with the property that h = b ± 1.
// Find Σ L for the twelve smallest isosceles triangles for which h = b ± 1 and
// b, L are positive integers.

// So... the brute force solution, even with lots of optimization and inlining is
// really slow. So slow that it never finished. But, it gave me enough terms to
// look up in the OEIS and get a high confidence level that the sequence is equivalent
// to 4x the sum of the squares of the first n Fibonacci numbers with index divisible by 3.
//
// https://oeis.org/A156084
//
// So, I coded that. Finishes in about 300 msecs.

var (
	one     = new(big.Int).SetInt64(1)
	two     = new(big.Int).SetInt64(2)
	four    = new(big.Int).SetInt64(4)
	sixteen = new(big.Int).SetInt64(16)
)

// findTrianglesSlow does an exhaustive search of all possible triangles
func findTrianglesSlow() {
	limit := 10

	fmt.Printf("First %d solutions:\n\n", limit)
	found := 0
	b := new(big.Int).SetInt64(16)
	sum := new(big.Int).SetInt64(0)
	h := new(big.Int).SetInt64(0)

	sumOfSquares := new(big.Int).SetInt64(0)
	squared := new(big.Int).SetInt64(0)
	L := new(big.Int).SetInt64(0)

	for {
		// The sequence alternates. First a -1 term then a +1 term. Repeat.
		for {
			h.Sub(b, one)

			// Is this a right triangle?
			sumOfSquares.Rsh(b, 1)
			sumOfSquares.Mul(sumOfSquares, sumOfSquares)
			squared.Mul(h, h)
			sumOfSquares.Add(sumOfSquares, squared)
			L.Sqrt(sumOfSquares)
			squared.Mul(L, L)
			if squared.Cmp(sumOfSquares) == 0 {
				sum.Add(sum, L)
				found++
				fmt.Printf("%2d  b: %12d-1  h: %12d  L: %12d  Σ L: %14d\n", found, b, h, L, sum)
				break
			}

			b.Add(b, sixteen)
		}

		if found == limit {
			break
		}

		// Now get a +1 term.
		for {
			h.Add(b, one)

			// Is this a right triangle?
			sumOfSquares.Rsh(b, 1)
			sumOfSquares.Mul(sumOfSquares, sumOfSquares)
			squared.Mul(h, h)
			sumOfSquares.Add(sumOfSquares, squared)
			L.Sqrt(sumOfSquares)
			squared.Mul(L, L)
			if squared.Cmp(sumOfSquares) == 0 {
				sum.Add(sum, L)
				found++
				fmt.Printf("%2d  b: %12d+1  h: %12d  L: %12d  Σ L: %14d\n", found, b, h, L, sum)
				break
			}

			b.Add(b, sixteen)
		}

		if found == limit {
			break
		}
	}
}

// fibonacci returns the first f terms in the sequence
func fibonacci(f int) []uint64 {
	if f <= 0 {
		return nil
	}

	fib := []uint64{0, 1, 1}

	l := len(fib)
	for i := 1; i <= f; i++ {
		next := fib[l-2] + fib[l-1]
		fib = append(fib, next)
		l++
	}

	return fib
}

// findTrianglesFast prints 4x the sum of the squares of every 3rd term
func findTrianglesFast() {
	b := new(big.Int).SetInt64(0)
	bSum := new(big.Int).SetInt64(0)
	h := new(big.Int).SetInt64(0)
	halfB := new(big.Int).SetInt64(0)
	squared := new(big.Int).SetInt64(0)
	sumOfSquares := new(big.Int).SetInt64(0)
	L := new(big.Int).SetInt64(0)
	sumL := new(big.Int).SetInt64(0)

	limit := 12
	f := fibonacci(limit * 3)

	for i := 1; i <= limit; i++ {
		fIndex := new(big.Int).SetUint64(f[i*3])
		b.Mul(fIndex, fIndex)
		b.Mul(b, four)
		bSum.Add(bSum, b)
		if i&0x01 == 0x01 {
			h.Sub(bSum, one)
		} else {
			h.Add(bSum, one)
		}
		halfB.Rsh(bSum, 1)
		sumOfSquares.Mul(halfB, halfB)
		squared.Mul(h, h)
		sumOfSquares.Add(sumOfSquares, squared)
		L.Sqrt(sumOfSquares)
		sumL.Add(sumL, L)
		fmt.Printf("%2d:  bSum: %18d  h: %18d  L: %20d  Σ L: %22d\n", i, bSum, h, L, sumL)
	}
}

func main() {
	fmt.Printf("Welcome to 138\n\n")

	//findTrianglesSlow()
	findTrianglesFast()
}
