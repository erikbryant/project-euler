package main

// go fmt ./... && go vet ./... && go test && go run 072.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/erikbryant/util-golang/algebra"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
)

// Consider the fraction, n/d, where n and d are positive integers. If n<d and
// HCF(n,d)=1, it is called a reduced proper fraction.
//
// If we list the set of reduced proper fractions for d ≤ 8 in ascending order
// of size, we get:
//
// 1/8, 1/7, 1/6, 1/5, 1/4, 2/7, 1/3, 3/8, 2/5, 3/7, 1/2, 4/7, 3/5, 5/8, 2/3,
// 5/7, 3/4, 4/5, 5/6, 6/7, 7/8
//
// It can be seen that there are 21 elements in this set.
//
// How many elements would be contained in the set of reduced proper fractions
// for d ≤ 1,000,000?

// A reduced proper fraction is one where the numerator and denominator are
// coprime. To solve, we just need to count the number of coprime pairs of
// 2 <= n < d where d <= 1,000,000

// countNumerators returns the number of numerators that are coprime to d
func countNumerators(d int) int {
	return algebra.Totient(d)
}

func looper(maxD int) int {
	count := 0

	for d := 2; d <= maxD; d++ {
		count += countNumerators(d)
	}

	return count
}

func main() {
	fmt.Printf("Welcome to 072\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	maxD := 1000 * 1000
	count := looper(maxD)
	fmt.Println("Set of reduced fractions for d <=", maxD, "is:", count)
}
