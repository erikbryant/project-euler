package main

// go fmt && golint && go test && go run 745.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

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

func maxSquare(n int) int {
	max := 1
	maxRoot := int(math.Sqrt(float64(n)))

	for i := maxRoot; i > 1; i-- {
		square := i * i
		if n%square == 0 {
			max = square
			break
		}
	}

	return max
}

func sumSquares(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		sum += maxSquare(i)
	}

	return sum
}

func seive(n int64) int64 {
	var sum int64
	var i int64
	var maxRoot int64

	maxRoot = int64(math.Sqrt(float64(n)))
	rootCount := make([]int64, maxRoot+1)
	sum = 0

	// Fill the slice with counts of how many times each root would
	// appear in the sum if we were summing *all* roots.
	for i = 1; i <= maxRoot; i++ {
		rootCount[i] = n / (i * i)
	}

	// Subtract the number of times that the higher root...
	for i = maxRoot; i >= 1; i-- {
		sum += i * i * rootCount[i]
		// ...shadows a lower root.
		for j := i / 2; j >= 1; j-- {
			if i%j == 0 {
				rootCount[j] -= rootCount[i]/j + rootCount[i]%j
			}
		}
	}

	return sum
}

func main() {
	fmt.Printf("Welcome to 745\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	// fmt.Println(sumSquares(50))
	// fmt.Println()
	// fmt.Println("Sum:", seive(50))

	fmt.Println("Sum:", seive(100*1000*1000*1000))
	// fmt.Println(seive(100 * 1000 * 1000 * 1000 * 1000))
	// fmt.Println(sumSquares(100))
}
