package main

import (
	"fmt"
)

const MAX_PENT = 10000

var Pent [MAX_PENT]int

func makePent() {
	for n := 1; n < MAX_PENT; n++ {
		Pent[n] = n * (3*n - 1) / 2
	}
}

func pentagonal(n int) bool {
	for i := 0; i < MAX_PENT; i++ {
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

	for i = 1; i < MAX_PENT-1; i++ {
		for j = i + 1; j < MAX_PENT; j++ {
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
