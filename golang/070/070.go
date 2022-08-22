package main

// go fmt ./... && go vet ./... && go test && go run 070.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

import (
	"flag"
	"fmt"
	"github.com/erikbryant/project-euler/golang/primes"
	"github.com/erikbryant/project-euler/golang/util"
	"log"
	"os"
	"runtime/pprof"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
)

func init() {
	primes.Load("../primes.gob")
}

// https://projecteuler.net/problem=70
//
// Euler's Totient function, φ(n) [sometimes called the phi function], is used
// to determine the number of positive numbers less than or equal to n which
// are relatively prime to n. For example, as 1, 2, 4, 5, 7, and 8, are all less
// than nine and relatively prime to nine, φ(9)=6.
//
// The number 1 is considered to be relatively prime to every positive number,
// so φ(1)=1.
//
// Interestingly, φ(87109)=79180, and it can be seen that 87109 is a permutation
// of 79180.
//
// Find the value of n, 1 < n < 10^7, for which φ(n) is a permutation of n and
// the ratio n/φ(n) produces a minimum.

// isPermutation returns whether the two numbers are permutations of each other
func isPermutation(a, b int) bool {
	digits := make(map[int]int)

	for a > 0 {
		r := a % 10
		digits[r]++
		a /= 10
	}

	for b > 0 {
		r := b % 10
		digits[r]--
		if digits[r] == 0 {
			delete(digits, r)
		}
		b /= 10
	}

	return len(digits) == 0
}

func looper(maxN int) (int, int, float64) {
	minRatio := 99999999999.0
	minN := 0
	minT := 0

	for n := 2; n < maxN; n++ {
		if primes.Prime(n) {
			// Totient of a prime n is n-1 which cannot be a permutation of n
			continue
		}

		t := util.Totient(n)

		if isPermutation(t, n) {
			ratio := float64(n) / float64(t)
			if ratio < minRatio {
				minRatio = ratio
				minN = n
				minT = t
			}
		}
	}

	return minN, minT, minRatio
}

func main() {
	fmt.Printf("Welcome to 070\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	// maxN := 1000 // min n: 291 min t: 192 min ratio: 1.515625
	// maxN := 10 * 1000 // min n: 1975 min t: 1579 min ratio: 1.2507916402786574
	// maxN := 100 * 1000 // min n: 75841 min t: 75184 min ratio: 1.0087385613960418
	// maxN := 1000 * 1000 // min n: 783169 min t: 781396 min ratio: 1.0022690159662961
	maxN := 10 * 1000 * 1000 // min n: 8319823 min t: 8313928 min ratio: 1.0007090511248113

	minN, minT, minRatio := looper(maxN)
	fmt.Println("min n:", minN, "min t:", minT, "min ratio:", minRatio)
}
