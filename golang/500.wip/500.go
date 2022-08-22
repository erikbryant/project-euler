package main

// go fmt ./... && go vet ./... && go test && go run 500.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

import (
	"github.com/erikbryant/project-euler/golang/primes"
	"flag"
	"fmt"
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

// The number of divisors of 120 is 16.
// In fact 120 is the smallest number having 16 divisors.
//
// Find the smallest number with 2^500500 divisors.
// Give your answer modulo 500500507.

// Not sure what use this data is, but this is the pattern I have
// see so far. For any given number (#) the number of factors is
// often a power of two. The number divided by the previous number
// that had a power-of-two number of factors is prime (for numbers
// >= 120).
//
// 2-power     			 #	#factors    #/#prev
//   1               2		  2					 -
//   2               6		  4					 3
//   3              24		  8					 4
//   4             120     16          5
//   5             840		 32					 7
//   7           83160		128					11
//   8         1081080		256					13
//   9        18378360		512					17

func main() {
	fmt.Printf("Welcome to 500\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

}
