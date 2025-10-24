package swing

// Prime-Swing factorial
//
// http://www.luschny.de/math/factorial/FastFactorialFunctions.htm
// http://www.luschny.de/math/factorial/SwingIntro.pdf

import (
	"160/dnc"
	"math"
	"sort"
)

import (
	primePkg "github.com/erikbryant/util-golang/primes"
)

var (
	// Mod is the global digit mask. Don't change this. Unless you hate yourself.
	Mod = 10000000
)

// product returns p and fives where p*5^fives =∏s
func product(s []int) (int, int) {
	p := 1
	fives := 0
	for _, v := range s {
		for v%5 == 0 {
			fives++
			v /= 5
		}
		p *= v
		p %= Mod
	}
	return p, fives
}

// swing returns n⎱
func swing(m int) (int, int) {
	if m < 4 {
		return []int{1, 1, 1, 3}[m], 0
	}

	primes := primePkg.PackedPrimes

	mSqrt := int(math.Sqrt(float64(m)))
	s := sort.SearchInts(primes, 1+mSqrt)
	d := sort.SearchInts(primes, 1+m/3)
	e := sort.SearchInts(primes, 1+m/2)
	g := sort.SearchInts(primes, 1+m)

	factors := append([]int{}, primes[e:g]...)

	for i := s; i < d; i++ {
		p := primes[i]
		if (m/p)&0x01 == 1 {
			factors = append(factors, p)
		}
	}

	for i := 1; i < s; i++ {
		prime := primes[i]
		p, q := 1, m
		for {
			q /= prime
			if q == 0 {
				break
			}
			if q&1 == 1 {
				p *= prime
			}
		}
		if p > 1 {
			factors = append(factors, p)
		}
	}

	return product(factors)
}

// factorialOdd returns m and k where 2^?*m*5^k = n!
func factorialOdd(n int) (int, int) {
	if n < 2 {
		return 1, 0
	}

	// f = oddFactorial(n/2, primes)^2 * swing(n, primes)

	f := 1
	fives := 0

	// Highest power of two <= n
	i := int(math.Log2(float64(n)))
	two := int(math.Pow(2, float64(i)))

	for ; two > 0; two /= 2 {
		f *= f
		fives *= 2
		f %= Mod
		fSwing, five := swing(n / two)
		f *= fSwing
		f %= Mod
		fives += five
	}

	return f, fives
}

func Factorial(n int) int {
	twos := dnc.FactorialEven(n)
	f, fives := factorialOdd(n)
	f = dnc.Fix(f, twos, fives)
	return f
}
