package main

// go fmt ./... && go vet ./... && go test && go build 142.go && time ./142

import (
	"fmt"

	"github.com/erikbryant/util-golang/algebra"
)

// Find the smallest x + y + z with integers x > y > z > 0 such that
//   x + y, x - y, x + z, x - z, y + z, y - z
// are all perfect squares.

// check returns true if all the problem's conditions have been met
func check(x, y, z int) bool {
	// x > y > z > 0
	if x <= y || y <= z || z <= 0 {
		return false
	}

	//   x + y, x + z, y + z are all perfect squares
	if !algebra.IsSquare(x + y) {
		return false
	}

	if !algebra.IsSquare(x + z) {
		return false
	}

	if !algebra.IsSquare(y + z) {
		return false
	}

	//   x - y, x - z, y - z are all perfect squares
	if !algebra.IsSquare(x - y) {
		return false
	}

	if !algebra.IsSquare(x - z) {
		return false
	}

	if !algebra.IsSquare(y - z) {
		return false
	}

	return true
}

// findTriple returns the lowest (x, y, z) triple that passes check()
func findTriple() (int, int, int) {
	for x := 1; ; x++ {
		for i := 1; i*i < x; i++ {
			// IsSquare(x - y)
			y := x - i*i
			if !algebra.IsSquare(x + y) {
				continue
			}
			for j := 1; j*j < y; j++ {
				// IsSquare(y - z)
				z := y - j*j
				if check(x, y, z) {
					return x, y, z
				}
			}
		}
	}
}

func main() {
	fmt.Printf("Welcome to 142\n\n")

	x, y, z := findTriple()
	fmt.Printf("sum( x: %d  y: %d  z: %d ) = %d\n", x, y, z, x+y+z)
}
