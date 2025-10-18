package main

// go fmt ./... && go vet ./... && go test && go build 139.go && time ./139

import (
	"fmt"

	"github.com/erikbryant/util-golang/algebra"
)

// Let (a, b, c) represent the three sides of a right angle triangle with integral length sides.
// It is possible to place four such triangles together to form a square with length c.
//
// For example, (3, 4, 5) triangles can be placed together to form a 5 by 5 square with a 1 x 1 hole
// in the middle, and it can be seen that the 5 by 5 square can be tiled with twenty-five 1 x 1 squares.
//
// https://projecteuler.net/resources/images/0139.png?1678992052
//
// However, if (5, 12, 13) triangles were used then the hole would measure 7 by 7 and these could not
// be used to tile the 13 by 13 square.
//
// Given that the perimeter of the right triangle is less than one-hundred million, how many Pythagorean
// triangles would allow such a tiling to take place?

func CanTile(a, b, c int) bool {
	a, b = max(a, b), min(a, b)
	sideInner := a - b
	sideOuter := c
	return sideOuter%sideInner == 0
}

func main() {
	fmt.Printf("Welcome to 139\n\n")

	upper := 1000 * 1000 * 100
	count := 0
	for _, ptp := range algebra.PythagoreanTriples(upper) {
		a, b, c := ptp[0], ptp[1], ptp[2]
		if a+b+c >= upper {
			continue
		}
		if CanTile(a, b, c) {
			k := (upper - 1) / (a + b + c)
			count += k
			fmt.Printf("Pythagorean primitive: %v adds %d unique solutions for a total of %d\n", ptp, k, count)
		}
	}

	fmt.Printf("\nCount: %d unique solutions.\n", count)
}
