package main

// go fmt ./... && go vet ./... && go test ./... && go build 220.go && time ./220
// go fmt ./... && go vet ./... && go test ./... && go build 220.go && ./220 && echo top | go tool pprof cpu.prof

import (
	"fmt"
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

// https://en.wikipedia.org/wiki/Dragon_curve

// isLeft returns true if the robot should turn left
func isLeft(n int) bool {
	// https://oeis.org/A091067
	n /= n & -n
	return n&0x3 == 3
}

// runSteps returns the ending coordinates at n steps
func runSteps(step, n, x, y, xDelta, yDelta int) (int, int) {
	for ; step <= n; step++ {
		// Take a step
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

func flipper(n int) (int, int) {
	x, y := 0, 0
	xDelta, yDelta := 0, 1

	if n == 0 {
		return x, y
	}

	// Take the first step
	step := 1
	x, y = runSteps(step, 1, x, y, xDelta, yDelta)

	// Rotationally duplicate the dragon
	for step*2 <= n {
		x, y = x+y, y-x
		step *= 2
	}

	// Complete any remaining steps to get to n
	step++
	xDelta, yDelta = 0, -1
	x, y = runSteps(step, n, x, y, xDelta, yDelta)

	return x, y
}

func main() {
	fmt.Printf("Welcome to 220\n\n")

	fileHandle, _ := os.Create("cpu.prof")
	_ = pprof.StartCPUProfile(fileHandle)
	defer pprof.StopCPUProfile()

	steps := 1000 * 1000 * 1000 * 1000
	x, y := flipper(steps)
	fmt.Printf("After %d steps, position = %d,%d\n\n", steps, x, y)

}
