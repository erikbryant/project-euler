package library

import (
	"../primes"
	"fmt"
	"math"
	"math/big"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	primes.Load("../primes.gob")
}

// CtrlT() prints a debugging message when SIGUSR1 is sent to this process.
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

// E() returns the nth number (1-based) in the convergent series
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

// convergent() returns the nth convegence of whichever series you pass in a function for.
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

// Factors() returns a list of the unique prime factors of n.
func Factors(n int) []int {
	f := make([]int, 0)

	root := int(math.Sqrt(float64(n)))
	for i := 0; primes.PackedPrimes[i] <= root; i++ {
		if n%primes.PackedPrimes[i] == 0 {
			f = append(f, primes.PackedPrimes[i])
			// Since we are iterating only up to root (as opposed to n/2)
			// we need to also add the 'reciprocal' factors. For instance,
			// when n=10 we iterate up to 3, which would miss 5 as a factor.
			d := n / primes.PackedPrimes[i]
			if d != primes.PackedPrimes[i] && primes.Prime(d) {
				f = append(f, d)
			}
		}
	}

	return f
}

// FactorsCounted() returns a map of prime factors of n with counts
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
