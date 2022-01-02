package main

// go fmt && golint && go test && go run 621.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

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
	triangles  = []int{}
)

// Gauss famously proved that every positive integer can be expressed as the
// sum of three triangular numbers (including 0, the lowest triangular number).
// In fact most numbers can be expressed as a sum of three triangular numbers
// in several ways.
//
// Let G(n) be the number of ways of expressing n as the sum of three triangular
// numbers, regarding different arrangements of the terms of the sum as distinct.
//
// For example, G(9) = 7, as 9 can be expressed as:
//   3+3+3, 0+3+6, 0+6+3, 3+0+6, 3+6+0, 6+0+3, 6+3+0.
//
// You are given G(1000) = 78 and G(10^6) = 2106. Find G(17526 x 10^9).

// makeTriangles finds all triangular numbers up to a given max
func makeTriangles() {
	t := 0
	for i := 0; t < 17526*1000*1000*1000; i++ {
		t = (i * (i + 1)) >> 1
		triangles = append(triangles, t)
	}
}

// findTriangle returns index of 't' in triangles for range of min<=i<=max
func findTriangle(t, min, max int) int {
	var i int

	for min <= max {
		i = (min + max) >> 1
		if triangles[i] == t {
			return i
		}
		if triangles[i] < t {
			min = i + 1
			continue
		}
		max = i - 1
	}

	return i
}

// trianglar returns true if n is a trianglar number
func triangular(n int) bool {
	// n is triangular if 8*n+1 is a square
	root := math.Sqrt(float64(n<<3 + 1))
	return root == math.Trunc(root)
}

// tSumCount returns the # of combinations of triangular numbers that sum to n
func tSumCount(n int) int {
	count := 0

	for i := findTriangle(n, 0, len(triangles)-1); triangles[i] >= n/3; i-- {
		ti := triangles[i]

		// Find the largest triangular number <= min(ti, need).
		need := n - ti
		if need > ti {
			need = ti
		}
		tjRoot := int(math.Sqrt(float64(8*need+1))-1) >> 1
		tj := (tjRoot * (tjRoot + 1)) >> 1

		for ; tj >= 0 && tjRoot >= 0; tj, tjRoot = tj-tjRoot, tjRoot-1 {
			tk := n - ti - tj
			if tk > tj {
				break
			}
			if triangular(tk) {
				// 2 or 3 being the same is very rare. Nest the checks for speed.
				if ti == tj || tj == tk {
					if ti == tj && tj == tk {
						// (a, a, a)
						count++
						continue
					}
					// (a, a, b) (a, b, a) (b, a, a)
					count += 3
					continue
				}
				// (a, b, c) (a, c, b) (b, a, c) (b, c, a) (c, a, b) (c, b, a)
				count += 6
			}
		}
	}

	return count
}

func main() {
	fmt.Printf("Welcome to 621\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	makeTriangles()

	// t := 10
	// t := 17526
	// t := 17526 * 10
	// t := 17526 * 1000
	// t := 17526 * 1000 * 10
	// t := 17526 * 1000 * 1000
	// t := 17526 * 1000 * 1000 * 10
	// t := 17526 * 1000 * 1000 * 100
	t := 17526 * 1000 * 1000 * 1000

	// G(17526 * 10^9) = 11429712
	fmt.Printf("G(%d) = %d\n", t, tSumCount(t))
}
