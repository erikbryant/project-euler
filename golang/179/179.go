package main

// go fmt && golint && go test && go run 179.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

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

// Find the number of integers 1 < n < 107, for which n and n + 1 have the same
// number of positive divisors. For example, 14 has the positive divisors 1, 2,
// 7, 14 while 15 has 1, 3, 5, 15.

func divisors(n int) int {
	d := 1 // 1

	root := int(math.Sqrt(float64(n)))

	for i := 2; i <= root; i++ {
		if n%i == 0 {
			d += 2 // i and its n/i compliment
		}
	}

	if root*root == n {
		d-- // n is a perfect square, so we overcounted above
	}

	d++ // n

	return d
}

func looper(max int) int {
	cpd := 0

	current := divisors(1)
	for i := 2; i < max; i++ {
		next := divisors(i)
		if current == next {
			cpd++
		}
		current = next
	}

	return cpd
}

func main() {
	fmt.Printf("Welcome to 179\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	max := 10 * 1000 * 1000
	cpd := looper(max)

	fmt.Println("Number of integers 1 < n <", max, " for which n and n + 1 have the same number of positive divisors:", cpd)
}
