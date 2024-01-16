package main

// go fmt ./... && go vet ./... && go test && go run 091.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

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

// sideLen returns the distance between two points
func sideLen(X1, Y1, X2, Y2 int) float64 {
	x1 := float64(X1)
	x2 := float64(X2)
	y1 := float64(Y1)
	y2 := float64(Y2)

	return math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
}

// minMax returns a,b,c such that c is the max of a,b,c
func minMax(a, b, c float64) (float64, float64, float64) {
	// c is largest
	if c >= a && c >= b {
		return a, b, c
	}

	// b is largest
	if b >= a && b >= c {
		return a, c, b
	}

	// a is largest
	return b, c, a
}

// epsilonEqual returns whether two values are equal (within a given epsilon)
func epsilonEqual(a, b float64) bool {
	e := 0.0001
	return math.Abs(a-b) <= e
}

// rightTriangle returns whether a triangle with the two given points and (0,0) is right
func rightTriangle(x1, y1, x2, y2 int) bool {
	OP := sideLen(0, 0, x1, y1)
	PQ := sideLen(x1, y1, x2, y2)
	OQ := sideLen(0, 0, x2, y2)

	a, b, c := minMax(OP, PQ, OQ)

	return epsilonEqual(a*a+b*b, c*c)
}

// looper returns how many triangles are right
func looper(max int) int {
	count := 0

	for x1 := 0; x1 <= max; x1++ {
		for y1 := 0; y1 <= max; y1++ {
			if x1 == 0 && y1 == 0 {
				continue
			}
			for x2 := 0; x2 <= max; x2++ {
				for y2 := 0; y2 <= max; y2++ {
					if x2 == 0 && y2 == 0 {
						continue
					}
					if x1 == x2 && y1 == y2 {
						continue
					}
					if rightTriangle(x1, y1, x2, y2) {
						count++
					}
				}
			}
		}
	}

	return count
}

func main() {
	fmt.Printf("Welcome to 091\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	// maxFound := 2
	maxFound := 50
	count := looper(maxFound)
	fmt.Println("For maxFound =", maxFound, "there are", count, "right triangles,", count/2, "of which are distinct")
}
