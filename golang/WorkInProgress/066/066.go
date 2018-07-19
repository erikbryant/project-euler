package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func square(y int) bool {
	root := math.Sqrt(float64(y))
	return root == float64(int(root)) && int(root)*int(root) == y
}

func solve(D, maxX int) int {
	// There are no solutions in positive integers when D is square.
	if square(D) {
		return 0
	}

	// x^2 – Dy^2 = 1
	// x^2 = Dy^2 + 1
	y := 2
	root := 0.0
	for {
		x2 := D*y*y + 1
		root = math.Sqrt(float64(x2))
		if root == float64(int(root)) && int(root)*int(root) == x2 {
			if int(root) < maxX {
				return int(root)
			}
			break
		}
		y++
	}

	x := int(root) - 1
	lastX := int(root)
	for x > 1 {
		if (x*x-1)%D == 0 {
			y2 := (x*x - 1) / D
			root := math.Sqrt(float64(y2))
			if root == float64(int(root)) {
				if x <= maxX {
					return x
				}
				lastX = x
			}
		}
		x--
	}

	return lastX
}

// Find the value of D ≤ max in minimal solutions of x for which the largest value of x is obtained.
func maxSolution(max int) (int, int) {
	maxX := 0
	maxD := 0
	for D := 2; D <= max; D++ {
		x := solve(D, maxX)
		if x >= maxX {
			maxD = D
			maxX = x
		}
	}

	return maxD, maxX
}

func main() {
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
