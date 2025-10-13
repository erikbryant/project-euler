package main

// go fmt ./... && go vet ./... && go test && go build 174.go && time ./174

import (
	// "github.com/erikbryant/util-golang/primes"
	"fmt"
)

// We shall define a square lamina to be a square outline with a square "hole" so that the shape
// possesses vertical and horizontal symmetry.

// Given eight tiles it is possible to form a lamina in only one way: 3 x 3 square with a 1 x 1
// hole in the middle. However, using thirty-two tiles it is possible to form two distinct laminae.
//
// https://projecteuler.net/resources/images/0173_square_laminas.gif?1678992055
//
// If t represents the number of tiles used, we shall say that t = 8 is type L(1) and t = 32 is type L(2).
// Let N(n) be the number of t <= 1,000,000 such that t is type L(n); for example, N(15) = 832.
//         10
// What is  Σ  N(n)?
//         n=1

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
	fmt.Printf("Welcome to 174\n\n")

	upper := 1000 * 1000
	count := 0
	tilesUsedBin := map[int]int{}

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
			tilesUsedBin[frameTiles]++
			//fmt.Printf("%4d: holeWidth: %5d  frameWidth: %5d  frameTiles: %6d\n", count, holeWidth, frameWidth, frameTiles)
		}
		if !found {
			break
		}
	}

	fmt.Printf("Using up to %d tiles, %d different square laminae can be formed.\n", upper, count)

	L := map[int]int{}
	for _, freq := range tilesUsedBin {
		L[freq]++
	}

	N := 0
	for n := 1; n <= 10; n++ {
		N += L[n]
	}
	fmt.Println("Sum of N from 1..10 = ", N)
}
