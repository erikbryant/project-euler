package main

import (
	"fmt"
)

func prefixMatch(prefix, target int64) bool {
	for {
		if prefix == target {
			return true
		}
		if prefix > target {
			break
		}
		target /= 10
	}

	return false
}

// Let p(L,n) be the nth-smallest value of j such that the base-10
// representation of 2^j begins with the digits of L.
// p(12,1) = 7
// p(12,2) = 80
// p(123,45) = 12710
func p(L, n int64) int64 {
	var val int64
	var j int64
	var found int64

	val = 1
	j = 0
	found = 0

	for {
		if prefixMatch(L, val) {
			found++
			// fmt.Printf("%10d %20d %10d\n", j, val, found)
			if found == n {
				return j
			}
		}

		val <<= 1
		j++

		// The 'trick' to this problem is that only some of the leftmost
		// digits matter when multiplying. Digits too far to the right
		// will never change the results of the leftmost digits. So,
		// freely truncate the rightmost digits. Of course, this assumption
		// fails for very long values of L. There may be a way to determine
		// the minimum leftmost digits needed based on the length of L.
		if val > 2305843009213693952 {
			val /= 100
		}
	}
}

// Find p(123,678910)
func main() {
	fmt.Println("Welcome to 686")
	fmt.Println("p(123, 678910) =", p(123, 678910))
}
