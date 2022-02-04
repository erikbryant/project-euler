package main

// go fmt && golint && go test && go run 187.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

import (
	"../primes"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"runtime/pprof"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
)

// A composite is a number containing at least two prime factors. For example,
// 15 = 3 × 5; 9 = 3 × 3; 12 = 2 × 2 × 3.
//
// There are ten composites below thirty containing precisely two, not
// necessarily distinct, prime factors: 4, 6, 9, 10, 14, 15, 21, 22, 25, 26.
//
// How many composite integers, n < 10^8, have precisely two, not necessarily
// distinct, prime factors?
//
// A composite number with precisely two prime factors is called "semiprime".
// https://en.wikipedia.org/wiki/Semiprime
//
//

func init() {
	primes.Load("../primes.gob")
}

// composite2 returns true if n has precisely two prime factors
func composite2(n int) bool {
	if primes.Prime(n) {
		return false
	}

	f := 0

	// 2 is much faster to do bitwise than division is
	for n&0x01 == 0 {
		f++
		if f > 2 {
			return false
		}
		n >>= 1
	}

	// Test using division
	for i := 1; primes.PackedPrimes[i] <= n; i++ {
		for {
			test := float64(n) / float64(primes.PackedPrimes[i])
			if test != math.Trunc(test) {
				break
			}
			f++
			if f > 2 {
				return false
			}
			n = int(test)
			if primes.Prime(n) {
				return (f + 1) == 2
			}
		}
	}

	return f == 2
}

func looper(max int) int {
	count := 0

	for i := 1; i <= max; i++ {
		if composite2(i) {
			count++
		}
	}

	return count
}

func main() {
	fmt.Printf("Welcome to 187\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	// max := 2 * 1000 * 1000
	max := 10 * 1000 * 1000
	// max := 100 * 1000 * 1000
	count := looper(max)

	fmt.Println("There are ", count, "2-composite integers <", max)
}
