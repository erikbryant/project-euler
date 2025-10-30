package main

// go fmt ./... && go vet ./... && go test && go build 132.go && time ./132
// go fmt ./... && go vet ./... && go test && go build 132 && ./132 && echo top | go tool pprof cpu.prof

import (
	"fmt"
	"math"
	"slices"

	"github.com/erikbryant/util-golang/algebra"
	"github.com/erikbryant/util-golang/primey"
)

// A number consisting entirely of ones is called a repunit. We shall define R(k) to be a repunit of length k.
// For example, R(10) = 1111111111 = 11 x 41 x 271 x 9091, and the sum of these prime factors is 9414.
// Find the sum of the first forty prime factors of R(10^9).

// It turns out the divisibility of R(n) by a given prime is periodic on n. If we know the period k for a given
// p then:
//
//   p is a prime factor of R(n) if n%k == 0

// Prime factors:
//   R(10^1): [11 41 271 9091]
//   R(10^2): [11 41 101 251 271 3541 5051 9091 21401 25601 27961 60101 7019801]
//   R(10^3): [11 41 73 101 137 251 271 401 751 1201 1601 3541 4001 5051 9091 21001 21401 24001 25601 27961 60101 76001 162251 1378001 1610501 1676321 7019801]

// ItoA converts an int to an []int8
func ItoA(n int) []int8 {
	if n == 0 {
		return []int8{0}
	}

	a := make([]int8, 20)
	i := len(a)

	for n >= 1 {
		i--
		a[i] = int8(n % 10)
		n /= 10
	}

	return a[i:]
}

// AtoI converts an []int8 to an int
func AtoI(s []int8) int {
	n := 0
	for _, ch := range s {
		n *= 10
		n += int(ch)
	}
	return n
}

// compare returns -1:s1<s2 0:s1==s2 1:s1>s2
func compare(s1, s2 []int8) int {
	if len(s1) < len(s2) {
		return -1
	}

	if len(s1) > len(s2) {
		return 1
	}

	return slices.Compare(s1, s2)
}

// subtract returns a-b as an []int8 and whether it is negative
func subtract(a []int8, bInt int) ([]int8, bool) {
	negative := false

	b := ItoA(bInt)

	switch compare(a, b) {
	case -1:
		negative = true
		b, a = a, b
	case 0:
		return []int8{0}, false
	}

	// a minus b
	aLen := len(a) - 1
	bLen := len(b) - 1
	borrow := int8(0)

	for ; bLen >= 0; bLen-- {
		a[aLen] -= b[bLen] + borrow
		borrow = 0
		if a[aLen] < 0 {
			a[aLen] += 10
			borrow = 1
		}
		aLen--
	}

	for ; borrow == 1; aLen-- {
		a[aLen] -= borrow
		borrow = 0
		if a[aLen] < 0 {
			a[aLen] += 10
			borrow = 1
		}
	}

	// Ignore leading zeroes
	i := 0
	for ; a[i] == 0; i++ {
	}

	return a[i:], negative
}

// PDRStr returns true if n is divisible by prime p NOTE: p must be prime, p != 2 && p != 5
func PDRStr(nOrig []int8, p int) bool {
	pStr := ItoA(p)

	switch compare(nOrig, pStr) {
	case -1:
		return false
	case 0:
		return true
	}

	// Make a copy because we are going to do update 'n' in place
	n := make([]int8, len(nOrig))
	copy(n, nOrig)

	// https://www.johndcook.com/blog/2021/02/17/divisibility-by-any-prime/
	// [1] R. A. Watson. Tests for Divisibility. The Mathematical Gazette, Vol. 87, No. 510 (Nov., 2003), pp. 493-494
	//
	// Let p be an odd prime and n a number we want to test for divisibility by p.
	// Write n as 10a + b where b is a single digit. Then there is a number k,
	// depending on p, such that n is divisible by p if and only if
	//
	//                  | a   b |
	//                  | k   1 |
	//
	// is divisible by p.
	//
	// So how do we find k?
	//
	// If p ends in 1 k = ⌊1p / 10⌋
	// If p ends in 3 k = ⌊7p / 10⌋
	// If p ends in 7 k = ⌊3p / 10⌋
	// If p ends in 9 k = ⌊9p / 10⌋
	//
	// Example:
	// Is 3293 divisible by 37? Since 37*3 = 111, k = 11.
	//   329 − 11×3 = 296
	//    29 − 11×6 = 37
	// Yes. 3293 is divisible by 37.

	k := p
	switch pStr[len(pStr)-1] {
	case 3:
		k *= 7
	case 7:
		k *= 3
	case 9:
		k *= 9
	}
	k /= 10

	negative := false
	for compare(n, pStr) == 1 {
		// n' = 10a - k*b where n == 10a + b
		b := int(n[len(n)-1])
		n = n[:len(n)-1]
		if len(n) == 0 {
			n = []int8{0}
		}

		// Reduce n
		n, negative = subtract(n, k*b)
		if negative {
			if compare(n, pStr) == 1 {
				return AtoI(n)%p == 0
			}
			break
		}
	}

	return slices.Equal(n, pStr) || slices.Equal(n, []int8{0})
}

// R returns a repunit of length n
func R(n int) []int8 {
	r := make([]int8, n)
	for i := 0; i < n; i++ {
		r[i] = 1
	}
	return r
}

// repunitFactorsSlow finds the first prime factors of R(10^9)
func repunitFactorsSlow() {
	upper := 1000 * 1000 * 1000
	maxCount := 40
	count := 0
	sum := 0

	// The prime factors of 10^9 are {2,5} so only periods
	// of multiples of 2^i*5^j need to be evaluated
	large := int(math.Pow(2.0, 9.0)) * int(math.Pow(2.0, 9.0))
	ks25 := []int{}
	for _, ks := range algebra.KSmooths(large, 5) {
		if ks == 1 || upper%ks != 0 {
			continue
		}
		ks25 = append(ks25, ks)
	}
	fmt.Printf("KS25: %v\n\n", ks25)

	// Remove values that are multiples of other values
	toDelete := map[int]bool{}
	for i := len(ks25) - 1; i >= 1; i-- {
		for j := i - 1; j >= 0; j-- {
			if toDelete[ks25[j]] {
				continue
			}
			if ks25[i]%ks25[j] == 0 {
				toDelete[ks25[j]] = true
			}
		}
	}
	ksFinal := []int{}
	for _, ks := range ks25 {
		if !toDelete[ks] {
			ksFinal = append(ksFinal, ks)
		}
	}
	ks25 = ksFinal
	fmt.Printf("KS25: %v\n\n", ks25)

	rBig := R(ks25[len(ks25)-1])

	fmt.Printf("Count   (ith)           factor      period      Σ(factors)\n")

	for i, p := range primey.Iterr(3, primey.Len()-1) { // p:{2,3,5} are not candidates
		if p > 160001 {
			break
		}
		for _, ks := range ks25 {
			r := rBig[:ks]
			if PDRStr(r, p) {
				count++
				sum += p
				fmt.Printf("%4d: %7dth    %10d     %8d    %10d\n", count, i, p, ks, sum)
				break
			}
		}
		if count >= maxCount {
			break
		}
	}
}

// ------------------------------------------------------------------------------------------------v
// math-based solution

// repunitFactors finds the first prime factors of R(10^9)
func repunitFactors() {
	// https://projecteuler.net/thread=132;page=2#6245
	//
	// The repunit generator is R(k)=(10^k-1)/9
	//
	// If n is a factor of R(k) then the remainder of 10^k/9*n has to be 1.
	// Find which primes divide 10^k, where k=10^9.
	sum := 0
	count := 0
	maxCount := 40
	exp := 1000 * 1000 * 1000
	for _, p := range primey.Iterr(3, primey.Len()-1) { // p:{2,3,5} are not candidates
		if algebra.PowerMod(10, exp, 9*p) == 1 {
			count++
			sum += p
			fmt.Printf("%2d:  %8d  %8d\n", count, p, sum)
			if count == maxCount {
				break
			}
		}
	}
}

// ------------------------------------------------------------------------------------------------^

func main() {
	fmt.Printf("Welcome to 132\n\n")

	// Math-based solution, as seen in the Project Euler forum.
	repunitFactors()
	fmt.Println()
	// My solution. More on the brute-force side. Works, but is slow.
	repunitFactorsSlow()
}
