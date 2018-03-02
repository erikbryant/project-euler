package main

import (
	"fmt"
)

func baseB(b int, max int) (repunits []int) {
	repunits = append(repunits, 1)

	for i := b + 1; i < max; i = i*b + 1 {
		repunits = append(repunits, i)
	}

	return
}

func isRepunit(b int, r int) bool {
	// In all bases, 1 is defined as a repunit.
	if r == 1 {
		return true
	}

	for r >= b && r%b == 1 {
		r--
		r = r / b
	}

	return r == 1
}

//
// The pattern for repunits is of the form:
// (For max=50 in this example)
//
// base  repunits
//   2   [1 3 7 15 31]
//   3   [1 4 13 40]
//   4   [1 5 21]
//   5   [1 6 31]
//   6   [1 7 43]
//   7   [1 8]
//   8   [1 9]
//   9   [1 10]
//  10   [1 11]
//  11   [1 12]
//  ...
//  47   [1 48]
//  48   [1 49]
//  49   [1]
//  50   [1]
//
// This means:
//   * 1 is a repunit. Always count it.
//   * For bases >= floor(sqrt(max)) there are no new
//     repunits added, so those can be ignored.
//   * The second digit of the candidate list can be ignored.
//
func sumRepunits(max int) int {
	repunits := make(map[int]int)

	// In all bases, 1 is defined as being a repunit
	repunits[1] = 1

	for b := 2; b < max; b++ {
		r := baseB(b, max)
		if len(r) <= 2 {
			break
		}
		for i := 2; i < len(r); i++ {
			val := r[i]
			repunits[val]++
		}
	}

	sum := 0
	for key := range repunits {
		sum += key
	}

	return sum
}

func main() {
	max := 1000 * 1000 * 1000 * 1000
	sum := sumRepunits(max)
	fmt.Println("max:", max, "sum:", sum)
}
