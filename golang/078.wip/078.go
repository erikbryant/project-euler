package main

// go fmt ./... && go vet ./... && go test && go run 078.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	pcache     = [10 * 1000][5000]int{}
)

// Let p(n) represent the number of different ways in which n coins can be
// separated into piles. For example, five coins can be separated into piles
// in exactly seven different ways, so p(5)=7.
//
// OOOOO
// OOOO   O
// OOO   OO
// OOO   O   O
// OO   OO   O
// OO   O   O   O
// O   O   O   O   O
//
// Find the least value of n for which p(n) is divisible by one million.

func phelper(n, max int) int {
	if pcache[n][max] != 0 {
		return pcache[n][max]
	}

	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	count := 1

	for i := 1; i < max && i < n; i++ {
		count += phelper(n-i, i)
	}

	pcache[n][max] = count

	return count
}

func permutations(n int) int {
	return phelper(n, n)
}

func looper() int {
	for i := 3; ; i++ {
		p := permutations(i)
		fmt.Printf("%10d %10d\n", i, p)
		if p%(1000*1000) == 0 {
			fmt.Println("Found it!", i, p)
			return i
		}
	}

	return 0
}

func main() {
	fmt.Printf("Welcome to 078\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	n := looper()
	fmt.Println("Solution:", n)
}
