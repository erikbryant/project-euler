package main

// go fmt ./... && go vet ./... && go test ./... && go build 549.go && time ./549
// go fmt ./... && go vet ./... && go test ./... && go build 549.go && ./549 && echo top | go tool pprof cpu.prof

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/erikbryant/util-golang/algebra"
	"github.com/erikbryant/util-golang/primes"
)

// The smallest number m such that 10 divides m! is m=5.
// The smallest number m such that 25 divides m! is m=10.
//
// Let s(n) be the smallest number m such that n divides m!.
// So s(10)=5 and s(25)=10.
// Let S(n) be Σ s(i) for 2 <= i <= n.
// S(100)=2012.
//
// Find S(10^8).

// TODO: Explore this option...
// Create a list of prime factors <= 10^8.
// For each prime factor, create a list of which numbers provide them.
// Since 2^27 > 10^8 we at most need to keep 27.
//
// 2 - 2|1, 4|3, 6|4, 8|7, ..., 2*27|xx
// 3 - 3|1, 6|2, 9|4, 12|5, ..., 3*27|yy
// 5 - 5|1, 10|2, 15|3, ..., 5*27|zz
// ...
// 17 - 17|1, ...

func countFactors(n int, f int) uint8 {
	count := uint8(0)
	for {
		if n%f != 0 {
			break
		}
		n = n / f
		count++
	}
	return count
}

func minProvider(fact int, count uint8) int {
	if count == 1 {
		return fact
	}

	sum := 0
	c := uint8(0)
	for c < count {
		sum += fact
		c += countFactors(sum, fact)
	}

	return sum
}

// Let s(n) be the smallest number m such that n divides m!
func s(n int) int {
	m := 0

	for fact, count := range algebra.FactorsCounted(n) {
		f := minProvider(fact, uint8(count))
		m = max(m, f)
	}

	return m
}

// Let S(n) be ∑s(i) for 2 ≤ i ≤ n
func sumS(n int) (sum int) {
	var start int = 2

	// Skip over 2 -> 2!
	start = 3
	sum += 2

	// Skip over 3 -> 3!
	start = 4
	sum += 3

	// Skip over 4 -> 4!
	start = 5
	sum += 4

	// Because the loop is unrolled, it must begin
	// on an odd number.
	if start&0x01 != 1 {
		fmt.Println("ERROR: loop is not synchronized")
		return 0
	}

	for a := start; a <= n; a++ {
		// Odd
		if primes.Prime(a) {
			sum += a
		} else {
			sum += s(a)
		}

		// Even
		a++
		if primes.Prime(a >> 1) {
			sum += a >> 1
		} else {
			sum += s(a)
		}
	}

	return
}

// Find S(10^8)
func main() {
	fmt.Println("Welcome to 549")

	flag.Parse()
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	_ = pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	n := 1000 * 1000 * 100
	fmt.Printf("For n: %d answer: %d\n\n", n, sumS(n))
}
