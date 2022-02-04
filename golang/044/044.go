package main

import (
	"fmt"
)

// MaxPent is the maximum pentagonal number we will consider
const MaxPent = 10000

// Pent is the list of pentagonal numbers
var Pent [MaxPent]int

func makePent() {
	for n := 1; n < MaxPent; n++ {
		Pent[n] = n * (3*n - 1) / 2
	}
}

func pentagonal(n int) bool {
	for i := 0; i < MaxPent; i++ {
		if n == Pent[i] {
			return true
		}
	}
	return false
}

func main() {
	makePent()

	minDist := 100000000
	minI := 0
	minJ := 0

	var i, j int

	for i = 1; i < MaxPent-1; i++ {
		for j = i + 1; j < MaxPent; j++ {
			p1 := Pent[i]
			p2 := Pent[j]
			if pentagonal(p1+p2) && pentagonal(p2-p1) {
				if p2-p1 < minDist {
					minDist = p2 - p1
					minI = i
					minJ = j
					fmt.Println("MinDist: ", minDist, minI, minJ)
				}
			}
		}
	}
	fmt.Println("MinDist: ", minDist, minI, minJ)
}
