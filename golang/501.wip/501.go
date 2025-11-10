package main

// go fmt ./... && go vet ./... && go test ./... && go build 501.go && time ./501
// go fmt ./... && go vet ./... && go test ./... && go build 501.go && ./501 && echo top | go tool pprof cpu.prof

import (
	"fmt"
	"os"
	"runtime/pprof"

	"github.com/erikbryant/util-golang/primey"
)

// The eight divisors of 24 are 1, 2, 3, 4, 6, 8, 12 and 24.
// The ten numbers not exceeding 100 having exactly eight divisors are
// 24, 30, 40, 42, 54, 56, 66, 70, 78 and 88.
// Let f(n) be the count of numbers not exceeding n with exactly eight divisors.
// You are given f(100) = 10, f(1000) = 180 and f(10^6) = 224427.
// Find f(10^12).

var (
	primes = []int{}
)

func init() {
	for _, p := range primey.Iter() {
		primes = append(primes, p)
	}
}

// factor1 returns the number of single-prime-factor numbers with 8 divisors
func factor1(n int) int {
	// Numbers with 1 prime factor
	// 8 = (a+1)
	// a=7
	// p^7 <= n

	count := 0

	for _, p := range primes {
		p7 := p * p * p * p * p * p * p
		if p7 > n {
			break
		}
		count++
	}

	return count
}

// factor2 returns the number of two-prime-factor numbers with 8 divisors
func factor2(n int) int {
	// Numbers with 2 prime factors
	// 8 = (a+1)(b+1)
	// a=1 and b=3  OR  a=3 and b=1
	// p1 * p2^3 <= n  OR  p1^3 * p2 <= n

	count := 0

	// Count p1^3 * p2
	for i := 0; ; i++ {
		p1Cubed := primes[i] * primes[i] * primes[i]
		if p1Cubed > n/primes[i+1] {
			break
		}
		// TODO: FIX OVERFLOW ----v
		for j := i + 1; primes[j] <= n/p1Cubed; j++ {
			count++
		}
		// TODO: FIX OVERFLOW ----^
	}

	// Count p1 * p2^3
	for i := 0; ; i++ {
		p1 := primes[i]
		p2 := primes[i+1]
		if p2*p2*p2 > n/p1 {
			break
		}
		for j := i + 1; ; j++ {
			p2 := primes[j]
			if p2*p2*p2 > n/p1 {
				break
			}
			count++
		}
	}

	return count
}

// factor3 returns the number of three-prime-factor numbers with 8 divisors
func factor3(n int) int {
	// Numbers with 3 prime factors
	// 8 = (a+1)(b+1)(c+1)
	// a=1 b=1 c=1
	// p1 * p2 * p3 <= n

	count := 0

	for i := 0; primes[i]*primes[i+1] <= n/primes[i+2]; i++ {
		p1 := primes[i]
		for j := i + 1; p1*primes[j] <= n/primes[j+1]; j++ {
			p2 := primes[j]
			// TODO: FIX OVERFLOW ----v
			for k := j + 1; primes[k] <= n/(p1*p2); k++ {
				count++
			}
			// TODO: FIX OVERFLOW ----^
		}
	}

	return count
}

// divisors8 returns the count of all numbers <= n with 8 divisors
func divisors8(n int) int {
	// The number of divisors for n is the product of powers(plus one) of its prime factors
	// primeFactors(n) = p1^a * p2^b * p3^c * ...
	// count(divisors(n)) = (a+1)(b+1)(c+1) ...
	return factor3(n) + factor2(n) + factor1(n)
}

func main() {
	fmt.Printf("Welcome to 501\n\n")

	fileHandle, _ := os.Create("cpu.prof")
	_ = pprof.StartCPUProfile(fileHandle)
	defer pprof.StopCPUProfile()

	upper := 1000 * 1000 * 100
	divisors := divisors8(upper)
	fmt.Printf("Numbers <= %d with exactly 8 divisors: %d\n\n", upper, divisors)
}
