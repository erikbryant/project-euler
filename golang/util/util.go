package util

import (
  "math"
)

// IsSquare returns true if f is a square
func IsSquare(n int) bool {
	root := math.Sqrt(float64(n))
	return root == math.Trunc(root)
}

// Generating permutation using Heap Algorithm
// https://www.geeksforgeeks.org/heaps-algorithm-for-generating-permutations/
func heapPermutation(digits []int, size int, c chan []int) {
	if size == 1 {
		var temp []int
		for i := 0; i < len(digits); i++ {
			temp = append(temp, digits[i])
		}
		c <- temp
		return
	}

	for i := 0; i < size; i++ {
		heapPermutation(digits, size-1, c)

		// if size is odd, swap first and last element
		// If size is even, swap ith and last element
		swap := 0
		if size%2 == 0 {
			swap = i
		}
		digits[swap], digits[size-1] = digits[size-1], digits[swap]
	}
}

func MakeDigits(n int, c chan []int) {
	defer close(c)

	var digits []int
	for i := 1; i <= n; i++ {
		digits = append(digits, i)
	}

	heapPermutation(digits, len(digits), c)
}
