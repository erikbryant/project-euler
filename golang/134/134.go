package main

// go fmt ./... && go vet ./... && go test && go run 134.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"runtime/pprof"

	"github.com/erikbryant/util-golang/primes"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
)

// Consider the consecutive primes p1 = 19 and p2 = 23. It can be verified that
// 1219 is the smallest number such that the last digits are formed by p1 whilst
// also being divisible by p2.
//
// In fact, with the exception of p1 = 3 and p2 = 5, for every pair of
// consecutive primes, p2 > p1, there exist values of n for which the last
// digits are formed by p1 and n is divisible by p2. Let S be the smallest of
// these values of n.
//
// Find ∑ S for every pair of consecutive primes with 5 ≤ p1 ≤ 1,000,000.

// LCM returns the lowest common multiple of a and b where the digits in LCM
// end with the digits in a
func LCM(a, b int) int {
	// We start with NNN where NNN are the digits in a. We add 1000 to it on each
	// iteration, resulting in 1NNN, 2NNN, 3NNN, etc. When we add the 1000 mask we
	// need to also step up bMult correspondingly, so calculate a bStep that is
	// close to, but less than, 1000.
	lenSuffix := math.Trunc(math.Log10(float64(a))) + 1.0
	mask := int(math.Pow(10.0, lenSuffix))
	bStep := b * (mask / b)

	candidate := a
	bMult := 0

	for {
		candidate += mask
		bMult += bStep
		for ; bMult < candidate; bMult += b {
			// Adding bStep gets us close to candidate.
			// Make sure we get all the way there.
		}
		if candidate == bMult {
			break
		}
	}

	return candidate
}

func looper(maxP int) int {
	sum := 0

	// fmt.Printf("     p1    p2            LCM            sum\n")

	for i := 0; primes.Primes[i] <= maxP; i++ {
		p1 := primes.Primes[i]
		if p1 < 5 {
			continue
		}
		p2 := primes.Primes[i+1]
		lcm := LCM(p1, p2)
		sum += lcm
		// fmt.Printf("%6d %6d %14d %14d\n", p1, p2, lcm, sum)
	}

	return sum
}

func main() {
	fmt.Printf("Welcome to 134\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	maxP := 1000 * 1000
	sum := looper(maxP)
	fmt.Println("For 5 <= p <=", maxP, "∑ S =", sum)
}
