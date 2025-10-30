package main

// go fmt ./... && go vet ./... && go test ./... && go build 070.go && time ./070
// go fmt ./... && go vet ./... && go test ./... && go build 070.go && ./070 && echo top | go tool pprof cpu.prof

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/erikbryant/util-golang/algebra"
	"github.com/erikbryant/util-golang/primey"
	"github.com/erikbryant/util-golang/util"
)

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

func findNinRange(maxN int) (int, int, float64) {
	minRatio := 99999999999.0
	minN := 0
	minT := 0

	for n := 2; n < maxN; n++ {
		if primey.Prime(n) {
			// Totient of a prime n is n-1 which cannot be a permutation of n
			continue
		}

		t := algebra.Totient(n)

		if util.IsDigitPermutation(t, n) {
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

	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	//maxN := 1000 // min n: 291 min t: 192 min ratio: 1.515625
	//maxN := 10 * 1000 // min n: 1975 min t: 1579 min ratio: 1.2507916402786574
	//maxN := 100 * 1000 // min n: 75841 min t: 75184 min ratio: 1.0087385613960418
	//maxN := 1000 * 1000 // min n: 783169 min t: 781396 min ratio: 1.0022690159662961
	maxN := 10 * 1000 * 1000 // min n: 8319823 min t: 8313928 min ratio: 1.0007090511248113

	minN, minT, minRatio := findNinRange(maxN)
	fmt.Println("min n:", minN, "min t:", minT, "min ratio:", minRatio)
	fmt.Printf("\n1 < n < 10^7 for which n/φ(n) is a minimum: %d\n\n", minN)
}
