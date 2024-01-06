package main

import (
	"fmt"
	"math"
	"math/big"

	"github.com/erikbryant/util-golang/primes"
)

// MaxCache is the max length of the factor cache
const MaxCache = 1000*1000*10 + 10

var queue [10]int
var head int = 0
var tail int = 0
var fCache [MaxCache + 1]int
var divisors = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199}

func push(n int) {
	queue[tail] = n
	tail++
	if tail >= len(queue) {
		tail = 0
	}
}

func pop() int {
	val := queue[head]
	head++
	if head >= len(queue) {
		head = 0
	}
	return val
}

func reset() {
	head = 0
	tail = 0
}

// Largest prime factor of n
func f(n int) int {
	// Reduce the number to make the prime hunt faster
	for _, d := range divisors {
		if d > n>>1 {
			break
		}
		for n%d == 0 {
			n = n / d
		}
		if n == 1 {
			return d
		}
		if primes.Prime(n) {
			return n
		}
	}

	if fCache[n] != -1 {
		return fCache[n]
	}

	for i := primes.PackedIndex(n >> 1); i >= 0; i-- {
		if n%primes.PackedPrimes[i] == 0 {
			fCache[n] = primes.PackedPrimes[i]
			return primes.PackedPrimes[i]
		}
	}

	if primes.Prime(n) {
		return n
	}
	fmt.Println("ERROR: did not find prime factor for ", n)
	return 0
}

func fnew(n *big.Int) *big.Int {
	return n
}

// Let h(n) be the maximum value of g(k) for 2 ≤ k ≤ n.
func h(n int) int {
	// We have already calculated quite a ways up;
	// start where we last left off
	last := 2
	max := 0
	// last := 100
	// max := 417
	// last := 1000 * 1000
	// max := 4475951
	// last := 1000 * 1000 * 10
	// max := 44925571
	val := 0

	reset()
	for i := 0; i <= 8; i++ {
		g := f(last + i)
		push(g)
		val += g
	}

	for k := last; k <= n; k++ {
		if val > max {
			max = val
		}
		g := f(k + 9)
		push(g)
		val = val - pop() + g
	}

	return max
}

func main() {
	for i := 0; i <= MaxCache; i++ {
		fCache[i] = -1
	}

	sum := f(100)
	if sum != 5 {
		fmt.Println("Expected 5, got ", sum)
	}
	sum = f(101)
	if sum != 101 {
		fmt.Println("Expected 101, got ", sum)
	}
	sum = f(100) + f(101) + f(102) + f(103) + f(104) + f(105) + f(106) + f(107) + f(108)
	if sum != 409 {
		fmt.Println("Expected 409, got ", sum)
	}

	val := int(math.Pow(10, 6))
	if val+10 > MaxCache {
		fmt.Println("ERROR: val is too large for MaxCache. ", val, " > ", MaxCache)
		return
	}
	if val+10 > primes.MaxPrime {
		fmt.Println("ERROR: val is too large for MaxPrime. ", val, " > ", primes.MaxPrime)
		return
	}
	sum = h(val)
	if sum != 4475951 {
		fmt.Println("Expected 4475951, got ", sum)
	}
	fmt.Println("For ", val, " sum = ", sum)
}
