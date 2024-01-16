package main

import (
	"fmt"
)

func figurate(f int) []int {
	if f < 3 || f > 8 {
		return []int{}
	}

	figurates := []int{}

	n := 0
	for {
		v := 0
		switch f {
		case 3:
			v = n * (n + 1) / 2
		case 4:
			v = n * n
		case 5:
			v = n * (3*n - 1) / 2
		case 6:
			v = n * (2*n - 1)
		case 7:
			v = n * (5*n - 3) / 2
		case 8:
			v = n * (3*n - 2)
		}
		n++
		if v < 1000 {
			continue
		}
		if v > 9999 {
			break
		}
		if v%100 < 10 {
			// No number can begin with a zero, so we
			// cannot have a number that ends with 0[0-9].
			continue
		}
		figurates = append(figurates, v)
	}

	return figurates
}

func findMatches(f [][]int, candidate []int, excluded map[int]bool) {
	end := candidate[len(candidate)-1] % 100

	// Termination condition. We have used each of the 6 figurates.
	if len(excluded) == 6 {
		start := candidate[0] / 100
		if start == end {
			sum := 0
			for _, val := range candidate {
				sum += val
			}
			fmt.Println("Found a cycle!", candidate, sum)
		}
	}

	// There are still more figurates to search. Recurse.
	for i := 3; i <= 8; i++ {
		if excluded[i] {
			continue
		}
		for _, val := range f[i] {
			start := val / 100
			if start == end {
				excluded[i] = true
				try := append(candidate, val)
				findMatches(f, try, excluded)
				delete(excluded, i)
			}
		}
	}

	// We reached a dead end with this cycle. Fall back.
}

func main() {
	f := [][]int{}
	for i := 0; i <= 8; i++ {
		f = append(f, figurate(i))
	}

	excluded := map[int]bool{}

	// 8 is the smallest set of numbers, so use it for searching.
	excluded[8] = true
	for _, val := range f[8] {
		findMatches(f, []int{val}, excluded)
	}
}
