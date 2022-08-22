package main

// go fmt ./... && go vet ./... && go test && go run 131.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

import (
	"flag"
	"fmt"
	"github.com/erikbryant/project-euler/golang/primes"
	"log"
	"os"
	"runtime/pprof"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
)

// There are some prime values, p, for which there exists a positive integer, n,
// such that the expression n^3 + n^2p is a perfect cube.
//
// For example, when p = 19, 8^3 + 8^2×19 = 123.
//
// What is perhaps most surprising is that for each prime with this property the
// value of n is unique, and there are only four such primes below one-hundred.
//
// How many primes below one million have this remarkable property?

// This sequence is https://oeis.org/A002407
// Also can be calculated as the difference of two consecutive cubes

func init() {
	primes.Load("../primes.gob")
}

// looper counts the difference between consecutive cubes where the differernce
// is prime
func looper(maxP int) int {
	i := 0
	cube := i * i * i
	count := 0

	for {
		prevCube := cube
		i++
		cube = i * i * i
		p := cube - prevCube
		if p >= maxP {
			break
		}
		if primes.Prime(p) {
			fmt.Println(p)
			count++
		}
	}

	return count
}

func main() {
	fmt.Printf("Welcome to 131\n\n")

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
	count := looper(maxP)
	fmt.Println("There are", count, "perfect cube solutions for n^3 + n^2*p where p <", maxP)
}
