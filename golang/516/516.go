package main

// go fmt ./... && go vet ./... && go test && go build 516.go && ./516 && echo top | go tool pprof cpu.prof
// go fmt ./... && go vet ./... && go test && go build 516.go && time ./516

import (
	"fmt"
	"math"

	"github.com/erikbryant/util-golang/algebra"
	"github.com/erikbryant/util-golang/primes"
)

// 5-smooth numbers are numbers whose largest prime factor doesn't exceed 5.
// 5-smooth numbers are also called Hamming numbers.
//
// Let S(L) be the sum of the numbers n not exceeding L such that Euler's totient
// function ɸ(n) is a Hamming number. S(100)=3728.
//
// Find S(10^12). Give your answer modulo 2^32.

// The formula for totient(n) is:
//
//             ⌈ k - 1 ⌉
//          N  |  i    |
//   Φ(n) = ∏  | p     | x (p - 1)
//         i=0 ⌊  i    ⌋     i
//
// Where p  are the prime factors of n and k  are their exponents
//        i                                 i
//
// Therefore Φ(n) is a Hamming number if:
//
//   p ∈ {2,3,5}  OR  k  == 1
//    i                i
//
//   AND
//
//   Hamming(p  - 1)
//            i
//
// So, if Φ(n) is a Hamming number then n must be of the form:
//
//    a    b    c    1    1          1
//   2  * 3  * 5  * p  * p  * ... * p
//                   1    2          i
//
// For all values of a, b, c and for all p  where p  has a unary exponent and (p  - 1) is a Hamming number
//                                        i        i                            i
//
// This can be rewritten as:
//
//       ⌈                    ⌉
//   H X | p  * p  * ... * p  |
//       ⌊  1    2          i ⌋
//
// Where H is the set of all Hamming numbers crossed with all combinations from the set of all primes where Hamming(p - 1)
//

// HammingPrimes returns a sorted list of primes where (prime-1) is a Hamming number
func HammingPrimes(upper int, hammings []int) []int {
	hPrimes := []int{}

	for _, hamming := range hammings {
		if hamming+1 > upper {
			break
		}
		if algebra.Hamming(hamming + 1) {
			// No need to add a duplicate Hamming
			continue
		}
		if !primes.Prime(hamming + 1) {
			continue
		}
		hPrimes = append(hPrimes, hamming+1)
	}

	return hPrimes
}

// S returns the count and sum of values of n for which totient(n) is a Hamming number
func S(upper, mod int) (int, int) {
	hammings := algebra.Hammings(upper)
	hPrimes := HammingPrimes(upper, hammings)

	count := 0
	sum := 0

	// Compute hn times all combinations of p1 * p2 * ... pn
	for _, ham := range hammings {
		hamList := []int{ham}
		for _, prime := range hPrimes {
			l := len(hamList)
			maxHam := upper / prime
			for i := 0; i < l; i++ {
				if hamList[i] > maxHam {
					continue
				}
				factor := hamList[i] * prime
				if factor > upper {
					continue
				}
				hamList = append(hamList, factor)
			}
		}
		for _, ham := range hamList {
			sum += ham % mod
			sum %= mod
		}
		count += len(hamList)
	}

	return count, sum
}

func main() {
	fmt.Printf("Welcome to 516\n\n")

	upper := 1000 * 1000 * 1000 * 1000
	exp := int(math.Log10(float64(upper)))
	mod := int(math.Pow(2, 32))
	count := 0
	sum := 0

	count, sum = S(upper, mod)
	fmt.Printf("S2(10^%d):  count = %d  sum %% %d = %d\n", exp, count, mod, sum)
}
