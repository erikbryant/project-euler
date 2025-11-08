package main

// go fmt ./... && go vet ./... && go test ./... && go build 220.go && time ./220
// go fmt ./... && go vet ./... && go test ./... && go build 220.go && ./220 && echo top | go tool pprof cpu.prof

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

// Let D_0 be the two-letter string "Fa".  For n\ge 1, derive D_n from D_{n-1}
// by the string-rewriting rules:
//
//   "a" → "aRbFR"
//   "b" → "LFaLb"
//
// Thus, D_0 = "Fa", D_1 = "FaRbFR", D_2 = "FaRbFRRLFaLbFR", and so on.
//
// These strings can be interpreted as instructions to a computer graphics program,
// with "F" meaning "draw forward one unit", "L" meaning "turn left 90 degrees",
// "R" meaning "turn right 90 degrees", and "a" and "b" being ignored. The initial
// position of the computer cursor is (0,0), pointing up towards (0,1).
//
// Then D_n is an exotic drawing known as the Heighway Dragon of order n. For example,
// D_{10} is shown below; counting each "F" as one step, the highlighted spot at (18,16)
// is the position reached after 500 steps.
//
// https://projecteuler.net/resources/images/0220.gif?1678992055
//
// What is the position of the cursor after 10^12 steps in D_50?
// Give your answer in the form x,y with no spaces.

// isLeft returns true if the robot should turn left
func isLeft(n int) bool {
	// https://oeis.org/A091067
	n /= n & -n
	return n&0x3 == 3
}

// runSteps returns the ending coordinates after n steps
func runSteps(n int) (int, int) {
	if n%2 == 1 {
		log.Fatal("runSteps: n must be even ", n)
	}

	x := 0
	y := 0
	xDelta := 0
	yDelta := 1

	// The loop is unrolled one layer to get faster heading check
	for step := 1; step <= n; step++ {
		// Take a step
		x += xDelta
		y += yDelta

		// Change heading
		if step&0x3 == 3 {
			xDelta, yDelta = -yDelta, xDelta
		} else {
			xDelta, yDelta = yDelta, -xDelta
		}

		step++

		//Take a step
		x += xDelta
		y += yDelta

		// Change heading
		if isLeft(step) {
			xDelta, yDelta = -yDelta, xDelta
		} else {
			xDelta, yDelta = yDelta, -xDelta
		}
	}

	return x, y
}

func main() {
	fmt.Printf("Welcome to 220\n\n")

	fileHandle, _ := os.Create("cpu.prof")
	_ = pprof.StartCPUProfile(fileHandle)
	defer pprof.StopCPUProfile()

	steps := 1000 * 1000 * 1000
	x, y := runSteps(steps)
	fmt.Printf("steps: %d  x: %d  y: %d\n\n", steps, x, y)
}
