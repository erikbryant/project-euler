package main

// go fmt ./... && go vet ./... && go test && go build 173.go && time ./173

import (
	"fmt"
)

// We shall define a square lamina to be a square outline with a square "hole" so that the shape
// possesses vertical and horizontal symmetry. For example, using exactly thirty-two square tiles
// we can form two different square laminae:
//
// https://projecteuler.net/resources/images/0173_square_laminas.gif?1678992055
//
// With one-hundred tiles, and not necessarily using all tiles at one time, it is possible to form
// forty-one different square laminae.
//
// Using up to one million tiles how many different square laminae can be formed?

// There is an outer square made of tiles. Call this 'frame'.
// There is an inner square that has no tiles. Call this 'hole'.
//
// The number of spaces in the inner square is a square number.
//
// The inner square must be centered and aligned on a tile boundary. Thus, if the
// outer square has an even root then the inner square must also have an even root.
//
// Number of tiles in outer frame = (width(hole) + frameWidth*2)^2 - width(hole)^2

func main() {
	fmt.Printf("Welcome to 173\n\n")

	upper := 1000 * 1000
	count := 0

	for holeWidth := 1; ; holeWidth++ {
		holeTiles := holeWidth * holeWidth
		found := false
		for frameWidth := holeWidth + 2; ; frameWidth += 2 {
			frameTiles := frameWidth*frameWidth - holeTiles
			if frameTiles > upper {
				break
			}
			count++
			found = true
			//fmt.Printf("%4d: holeWidth: %5d  frameWidth: %5d  frameTiles: %6d\n", count, holeWidth, frameWidth, frameTiles)
		}
		if !found {
			break
		}
	}

	fmt.Printf("Using up to %d tiles, %d different square laminae can be formed.\n\n", upper, count)
}
