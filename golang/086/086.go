package main

// go fmt ./... && go vet ./... && go test && go build 086.go && time ./086

// A spider, S, sits in one corner of a cuboid room, measuring 6 by 5 by 3,
// and a fly, F, sits in the opposite corner. By travelling on the surfaces
// of the room the shortest "straight line" distance from S to F is 10 and
// the path is shown on the diagram.
//
// However, there are up to three "shortest" path candidates for any given
// cuboid and the shortest route doesn't always have integer length.
//
// It can be shown that there are exactly 2060 distinct cuboids, ignoring rotations,
// with integer dimensions, up to a maximum size of M x M x M, for which
// the shortest route has integer length when M = 100. This is the least
// value of M for which the number of solutions first exceeds two thousand;
// the number of solutions when M = 99 is 1975.
//
// Find the least value of M such that the number of solutions first exceeds
// one million.

import (
	"fmt"
	"math"
)

func hypotenuse(a, b int64) float64 {
	return math.Sqrt(float64(a*a + b*b))
}

func shortestPathIntegral(x, y, z int64) bool {
	d1 := hypotenuse(x+y, z)
	d2 := hypotenuse(x+z, y)
	d3 := hypotenuse(x, y+z)

	short := min(d1, d2, d3)

	return math.Trunc(short) == short
}

// cuboidPaths returns the number of unique integral-sided cuboids up to dimension MxMxM with integral paths
func cuboidPaths(M int64) int64 {
	paths := int64(0)

	x := M
	for y := int64(1); y <= x; y++ {
		for z := int64(1); z <= y; z++ {
			p := shortestPathIntegral(x, y, z)
			if p {
				paths++
				//fmt.Printf("%3d, %3d, %3d -> %5t  paths: %5d\n", x, y, z, p, paths)
			}
		}
	}

	return paths
}

func main() {
	fmt.Printf("Welcome to 086\n\n")

	M := int64(1)
	target := int64(1000 * 1000)
	paths := int64(0)
	for {
		paths += cuboidPaths(M)
		if paths > target {
			fmt.Printf("M: %4d -> paths: %5d\n", M, paths)
			break
		}
		M++
	}
}
