package dnc

import (
	"math"
	"math/bits"
)

var (
	// Mod is the global digit mask. Don't change this. Unless you hate yourself.
	Mod = 10000000

	// MaxFives is a value greater than k where k is the largest 5^k factor we expect to encounter
	MaxFives = 16
)

// fix returns (f*2^(twos-fives))%Mod if twos > fives else (f*2^(fives-twos))%Mod (i.e., it puts back the excess 2's or 5's that multiply removed)
func fix(f, twos, fives int) int {
	if twos >= fives {
		twos -= fives
		fives = 0
	} else {
		fives -= twos
		twos = 0
	}

	for twos > 36 {
		f <<= 36
		twos -= 36
		f %= Mod
	}
	for twos > 0 {
		f <<= 1
		twos--
		f %= Mod
	}

	for fives > 0 {
		f *= 5
		fives--
		f %= Mod
	}

	return f
}

// oddsProduct returns f and j where f=(∏{start..end}[(x * f) % Mod])/5^j (i.e., the product of all odd numbers from start..n with all powers of 5 factored out)
func oddsProduct(start, end, f, fives int) (int, int) {
	start |= 0x01 // Round start up to the nearest odd number

	for i := start; i <= end; i += 2 {
		x := i
		for x%5 == 0 {
			fives++
			x /= 5
		}
		x %= Mod
		f *= x
		f %= Mod
	}

	return f, fives
}

// factorialOdd returns m and k where 2^?*m*5^k = n!
func factorialOdd(n int) (int, int) {
	//
	//  factorial(20) =
	//   i=4    16 *
	//     3    8 *
	//     2    4 * 12 * 20 *
	//     1    2 * 6 * 10 * 14 * 18 *
	//     0    1 * 3 * 5 * 7 * 9 * 11 * 13 * 15 * 17 * 19
	//
	//  Factoring out powers of 2 (factorialEven handles those) yields:
	//
	//  factorial(20) = 2^k *
	//   i=4    1 *
	//     3    1 *
	//     2    1 * 3 * 5 *
	//     1    1 * 3 * 5 * 7 * 9 *
	//     0    1 * 3 * 5 * 7 * 9 * 11 * 13 * 15 * 17 * 19
	//
	//  Each term can be computed from the next by multiplying by the extra odd
	//  numbers: e.g., to get from i=1 to i=0 multiply by (11 * 13 * 15 * 17 * 19).
	//
	//   log2(n)  n/2^i
	//      ∏       ∏  j
	//     i=0     j=1       [where j is odd]
	//

	f := 1
	fp := 1
	fives := 0
	five := 0

	// Highest power of two <= n
	i := int(math.Log2(float64(n)))
	two := int(math.Pow(float64(2), float64(i-1)))
	start := 1

	for ; two > 0; two /= 2 {
		end := n / two
		fp, five = oddsProduct(start, end, fp, five)
		f *= fp
		f %= Mod
		fives += five
		start = end + 1 + end&0x01 // force start to be odd
	}

	return f, fives
}

// factorialEven returns k where 2^k*m*5^? = n!
func factorialEven(n int) int {
	return n - bits.OnesCount(uint(n))
}

// Factorial returns the low-order log10(Mod) non-zero digits of n!
func Factorial(n int) int {
	// https://github.com/python/cpython/blob/5d2edf72d25c2616f0e13d10646460a8e69344fa/Modules/mathmodule.c#L1870

	twos := factorialEven(n)
	f, fives := factorialOdd(n)
	f = fix(f, twos, fives)

	return f
}
