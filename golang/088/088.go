package main

import (
	"fmt"

	"github.com/erikbryant/util-golang/algebra"
)

// go fmt ./... && go vet ./... && go test && go build 088.go && time ./088

// A natural number, N, that can be written as the sum and product of a given
// set of at least two natural numbers, {a_1, a_2, ..., a_k} is called
// a product-sum number: N = a_1 + a_2 + ... + a_k = a_1 x a_2 x ... x a_k.
//
// For example, 6 = 1 + 2 + 3 = 1 x 2 x 3.
//
// For a given set of size, k, we shall call the smallest N with this property
// a minimal product-sum number. The minimal product-sum numbers for sets of size,
// k = 2, 3, 4, 5, and 6 are as follows.
//
// k=2:  4 = 2 x 2 = 2 + 2
// k=3:  6 = 1 x 2 x 3 = 1 + 2 + 3
// k=4:  8 = 1 x 1 x 2 x 4 = 1 + 1 + 2 + 4
// k=5:  8 = 1 x 1 x 2 x 2 x 2 = 1 + 1 + 2 + 2 + 2
// k=6: 12 = 1 x 1 x 1 x 1 x 2 x 6 = 1 + 1 + 1 + 1 + 2 + 6
//
// Hence for 2 <= k <= 6, the sum of all the minimal product-sum numbers is
// 4+6+8+12 = 30; note that 8 is only counted once in the sum.
//
// In fact, as the complete set of minimal product-sum numbers for 2 <= k <= 12 is
// {4, 6, 8, 12, 15, 16}, the sum is 61.
//
// What is the sum of all the minimal product-sum numbers for 2 <= k <= 12000?

// justSum returns the sum of integers in a slice
func justSum(a []int) int {
	sum := 0

	for _, v := range a {
		sum += v
	}

	return sum
}

func addMoreDivisors(n, have, dListNext, divisorsNext int, dList, divisors []int, sums []int) []int {
	// We are trying to multiply up to 'n'
	// We have multiplied up to 'have'
	// 'dListNext' is the next element in the dList slice to be filled
	// 'divisorsNext' is the index of the next available divisor
	// dList is the list of divisors we have so far
	// divisors is a list of all divisors of n sorted ascending

	need := n / have

	if need == 1 {
		// The sum minus the number of terms. We'll add k to it later
		// which will put that difference back.
		sum := justSum(dList[0:dListNext]) - len(dList[0:dListNext])
		sums = append(sums, sum)
	}

	for i := divisorsNext; i >= 1; i-- {
		if dListNext > 0 && divisors[i] > dList[dListNext-1] {
			// Keep terms in descending order
			continue
		}
		if need%divisors[i] == 0 {
			dList[dListNext] = divisors[i]
			sums = addMoreDivisors(n, have*divisors[i], dListNext+1, i, dList, divisors, sums)
		}
	}

	return sums
}

var (
	permuteCache = map[int][]int{}
)

// permuteDivisors returns the minimum product-sum, or 2*k if none are found
func permuteDivisors(n, k int) int {
	sums, ok := permuteCache[n]
	if !ok {
		divisors := algebra.Divisors(n)
		dList := make([]int, k)
		sums = addMoreDivisors(n, 1, 0, len(divisors)-1, dList, divisors, []int{})
		permuteCache[n] = sums
	}

	minN := 2 * k
	for _, sum := range sums {
		if sum+k == n {
			minN = min(sum+k, minN)
		}
	}

	return minN
}

// findMinN returns the smallest N that is a product-sum of k terms
func findMinN(k int) int {
	minN := k * 2

	// N <= 2k   <- seen on OEIS
	// N > k+2   <- k-2 ones, and two terms > 1
	// Therefore: k+2 <= N <= 2k
	for n := k + 2; n <= 2*k; n++ {
		N := permuteDivisors(n, k)
		minN = min(minN, N)
	}

	return minN
}

func main() {
	fmt.Printf("Welcome to 088\n\n")

	found := map[int]bool{}

	upper := 12000
	for k := 2; k <= upper; k++ {
		N := findMinN(k)
		found[N] = true
	}

	sum := 0
	for val := range found {
		sum += val
	}

	fmt.Printf("For 2 <= k <= %d  unique product-sum sum = %d\n\n", upper, sum)
}
