package main

// go fmt ./... && go vet ./... && go test && go run 108.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

import (
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

// In the following equation x, y, and n are positive integers.
//
// 1/x + 1/y = 1/n
//
// For n = 4 there are exactly three distinct solutions:
//
// x=5, y=20
// x=6, y=12
// x=8, y=8
//
// What is the least value of n for which the number of distinct solutions
// exceeds one-thousand?

// checker returns whether x, y, and n satisfy 1/x + 1/y = 1/n
func checker(x, y, n int) bool {
	// Calculate differently than the algorithm in solver so that we get
	// independent confirmation. Avoid floating point math.
	return x*y == n*(x+y)
}

// solver returns how many pairs of x,y that solve 1/x + 1/y = N
func solver(N int) int {
	n := float64(N)
	var x, y float64
	count := 0

	// x and y must both be > n
	for y = n + 1; ; y++ {
		x = n * y / (y - n)
		if x == math.Trunc(x) {
			// Did we get overflow, floating point imprecision, etc.?
			if !checker(int(x), int(y), int(n)) {
				fmt.Println("Values failed check!!!!", x, y, n)
				continue
			}
			count++
		}
		if math.Trunc(x) <= y {
			break
		}
	}

	return count
}

func looper(minCount int) (int, int) {
	var n int
	var solutions int

	for n = 1; ; n++ {
		solutions = solver(n)
		if solutions > minCount {
			break
		}
	}

	return solutions, n
}

func main() {
	fmt.Printf("Welcome to 108\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	minCount := 1000
	count, n := looper(minCount)
	fmt.Println("First N to exceed", minCount, "solutions is n:", n, "with", count, "solutions")
}
