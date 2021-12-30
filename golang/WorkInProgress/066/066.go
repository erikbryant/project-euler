package main

// go fmt && golint && go test && go run 066.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

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

// isSquare returns true if f is a square
func isSquare(n int) bool {
	root := math.Sqrt(float64(n))
	return root == math.Trunc(root)
}

// solution returns the lowest value of x that satisfies x^2 - Dy^2 = 1
func solution(D int) int {
	if isSquare(D) {
		return 0
	}

	d := float64(D)

	var x float64

	for x = 2.0; ; x++ {
		// We should be doing:
		// y = math.Sqrt(x*x-1) / math.Sqrt(d)
		// But, that is too hard to solve. Instead, ignore the '-1' in the
		// equation. It is so small compared to the large values of x and d
		// that it can be a rounding error. This will allow us to easily
		// calculate candidates for y. From those we can see which actually
		// solve the equation.
		y := math.Round(x / math.Sqrt(d))
		if x == math.Sqrt(d*y*y+1.0) {
			return int(x)
		}
	}

	return 0
}

// maxSolution returns the value of D ≤ max in minimal solutions of x for which
// the largest value of x is obtained.
func maxSolution(max int) (int, int) {
	maxX := 0
	maxD := 0

	for D := 2; D <= max; D++ {
		x := solution(D)
		if x >= maxX {
			maxD = D
			maxX = x
		}
		fmt.Println(D, maxD, maxX)
	}

	return maxD, maxX
}

func main() {
	fmt.Printf("Welcome to 066\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	D := 1000
	maxD, maxX := maxSolution(D)
	fmt.Println("For D =", D, "MaxX =", maxX, "at D =", maxD)
}
