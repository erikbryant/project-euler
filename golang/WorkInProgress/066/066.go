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
	maxX := 1000000.0

	// 1 = x^2 - Dy^2
	// 1 + Dy^2 = x^2
	// Dy^2 = x^2 - 1
	// y^2 = (x^2 - 1)/D
	// y = sqrt((x^2 - 1)/D)

	// 1 = x^2 - Dy^2
	// Dy^2 + 1 = x^2
	// sqrt(Dy^2 + 1) = x
	//    or
	// (Dy^2 + 1)/x = x
	// Dy^2/x + 1/x = x
	// Dy(y/x) + 1/x = x

	for x := 2.0; x < maxX; x++ {
		y := math.Sqrt((x*x - 1.0) / d)
		if y == math.Trunc(y) {
			return int(x)
		}
	}

	// Large values of x overflow x^2. Move to an approximation function.
	// Check the result of that approximation against the original equation.
	for x := maxX; ; x++ {
		// For large x, x^2/D approaches (x^2-1)/D
		y := math.Trunc(x / math.Sqrt(d))
		// y^2 overflows, so don't calculate that!
		if x == d*y*(y/x)+1.0/x {
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
