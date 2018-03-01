package main

import (
	"fmt"
	"strconv"
)

func bouncy(b int) bool {
	if b <= 100 {
		return false
	}

	s := strconv.Itoa(b)

	ascending := false
	descending := false

	for i := len(s) - 1; i >= 1; i-- {
		if s[i] == s[i-1] {
			continue
		}
		if s[i] > s[i-1] {
			ascending = true
		} else {
			descending = true
		}
		if ascending && descending {
			return true
		}

	}

	return false
}

func countBouncy(max int) int {
	count := 0
	for b := 1; b <= max; b++ {
		if bouncy(b) {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println("Welcome to 112")

	n := 1
	count := 0
	for {
		if bouncy(n) {
			count++
		}

		// We need to find *exactly* 99%, so use integer math.
		if n-count == n/100 {
			pctBouncy := float64(count) / float64(n) * 100.0
			fmt.Println("For:", n, "bouncy =", count, "pct bouncy =", pctBouncy)
			break
		}

		n++
	}
}
