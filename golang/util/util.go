package util

import (
	"../primes"
	"fmt"
	"math"
	"math/big"
	"os"
	"os/signal"
	"sort"
	"syscall"
)

func init() {
	primes.Load("../primes.gob")
}

// CtrlT prints a debugging message when SIGUSR1 is sent to this process.
func CtrlT(str string, val *int, digits []int) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGUSR1)

	fmt.Println("$ kill -SIGUSR1", os.Getpid())

	go func() {
		for {
			_ = <-c
			fmt.Println("^T] ", str, *val, digits)
		}
	}()
}

type convergentSeries func(int) int64

// E returns the nth number (1-based) in the convergent series
// of the number e [2; 1,2,1, 1,4,1, 1,6,1, ... ,1,2k,1, ...]
func E(n int) int64 {
	if n == 1 {
		return int64(2)
	}
	if n%3 == 0 {
		return int64(2 * n / 3)
	}
	return int64(1)
}

// Sqrt2 returns the nth number (1-based) in the convergent series
// of the square root of 2: [2;(2)]
func Sqrt2(n int) int64 {
	if n == 1 {
		return int64(1)
	}
	return int64(2)
}

// Convergent returns the nth convegence of whichever series you pass in a function for.
func Convergent(n int, fn convergentSeries) (*big.Int, *big.Int) {
	numerator := big.NewInt(fn(n))
	denominator := big.NewInt(1)

	for n > 1 {
		// Invert
		denominator, numerator = numerator, denominator

		// Add e(n-1)
		product := big.NewInt(fn(n - 1))
		product.Mul(product, denominator)
		numerator.Add(numerator, product)

		n--
	}

	return numerator, denominator
}

// Factors returns a sorted list of the unique prime factors of n.
func Factors(n int) []int {
	if n == 2 {
		return []int{}
	}

	m := make(map[int]int)

	root := int(math.Sqrt(float64(n))) + 1
	for i := 0; primes.PackedPrimes[i] <= root; i++ {
		if n%primes.PackedPrimes[i] == 0 {
			m[primes.PackedPrimes[i]] = 1
			// Since we are iterating only up to root (as opposed to n/2)
			// we need to also add the 'reciprocal' factors. For instance,
			// when n=10 we iterate up to 3, which would miss 5 as a factor.
			d := n / primes.PackedPrimes[i]
			if primes.Prime(d) {
				m[d] = 1
			}
		}
	}

	f := []int{}

	for v := range m {
		f = append(f, v)
	}

	sort.Ints(f)

	return f
}

// FactorsCounted returns a map of prime factors of n with counts
// of how many times each factor divides into n.
func FactorsCounted(n int) map[int]int {
	factors := make(map[int]int)

	// Find all of the 2 factors, since they are quick
	for (n & 0x01) == 0 {
		factors[2]++
		n = n >> 1
		if n == 1 {
			return factors
		}
	}

	root := int(math.Sqrt(float64(n)))
	for i := 1; primes.PackedPrimes[i] <= root; i++ {
		p := primes.PackedPrimes[i]
		for n%p == 0 {
			factors[p]++
			n = n / p
			if n == 1 {
				return factors
			}
		}
	}

	// We did not find any factors for 'n',
	// so it must be prime.
	factors[n]++
	return factors
}

// IsSquare returns true if f is a square
func IsSquare(n int) bool {
	root := math.Sqrt(float64(n))
	return root == math.Trunc(root)
}

// heapPermutation generates a permutation using Heap Algorithm
// https://www.geeksforgeeks.org/heaps-algorithm-for-generating-permutations/
func heapPermutation(digits []int, size int, c chan []int) {
	if size == 1 {
		var temp []int
		for i := 0; i < len(digits); i++ {
			temp = append(temp, digits[i])
		}
		c <- temp
		return
	}

	for i := 0; i < size; i++ {
		heapPermutation(digits, size-1, c)

		// if size is odd, swap first and last element
		// If size is even, swap ith and last element
		swap := 0
		if size%2 == 0 {
			swap = i
		}
		digits[swap], digits[size-1] = digits[size-1], digits[swap]
	}
}

// MakeDigits generates all permutations of the first n digits.
// For example:
//   n=2 [1 2] [2 1]
//   n=3 [1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]
func MakeDigits(n int, c chan []int) {
	defer close(c)

	var digits []int
	for i := 1; i <= n; i++ {
		digits = append(digits, i)
	}

	heapPermutation(digits, len(digits), c)
}

// IsPalindromeString returns true if the string is a palindrome
func IsPalindromeString(p string) bool {
	head := 0
	tail := len(p) - 1

	for head < tail {
		if p[head] != p[tail] {
			return false
		}
		head++
		tail--
	}

	return true
}

// IsPalindromeInt returns true if the digits of p are a palindrome
func IsPalindromeInt(p []int) bool {
	head := 0
	tail := len(p) - 1

	for head < tail {
		if p[head] != p[tail] {
			return false
		}
		head++
		tail--
	}

	return true
}

// DigitSum returns the sum of the digits in the number.
func DigitSum(n int) (sum int) {
	for n > 0 {
		sum += n % 10
		n /= 10
	}

	return
}

// Harshad returns true if n is divisible by the sum of its digits.
func Harshad(n, sum int) bool {
	return n%sum == 0
}

// Triangular returns true if n is a trianglar number
func Triangular(n int) bool {
	// n is triangular if 8*n+1 is a square
	root := math.Sqrt(float64(n<<3 + 1))
	return root == math.Trunc(root)
}

// Totient returns how many numbers k are relatively prime to n where
// 1 <= k < n. Relatively prime means that they have no common divisors (other
// than 1). 1 is considered relatively prime to all other numbers.
func Totient(n int) int {
	factors := Factors(n)

	// 1 is Totient prime to every number, but is not in factors.
	count := n - 1

	for f := 0; f < len(factors); f++ {
		count -= (n - 1) / factors[f]
		for f2 := f + 1; f2 < len(factors); f2++ {
			// We subtracted too many. Account for 'shadowed' factors.
			count += (n - 1) / (factors[f] * factors[f2])
		}
	}

	return count
}
