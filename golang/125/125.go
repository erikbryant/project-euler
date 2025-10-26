package main

// go fmt ./... && go vet ./... && go test ./... && go build 125.go && time ./125

import (
	"fmt"
	"math"

	"github.com/erikbryant/util-golang/algebra"
	"github.com/erikbryant/util-golang/util"
)

// The palindromic number 595 is interesting because it can be written as the
// sum of consecutive squares: 6^2 + 7^2 + 8^2 + 9^2 + 10^2 + 11^2 + 12^2.
//
// There are exactly eleven palindromes below one-thousand that can be written
// as consecutive square sums, and the sum of these palindromes is 4164. Note
// that 1 = 02 + 12 has not been included as this problem is concerned with the
// squares of positive integers.
//
// Find the sum of all the numbers less than 10^8 that are both palindromic and
// can be written as the sum of consecutive squares.

func looper(max int) int {
	totalSum := 0
	alreadySeen := map[int]bool{}

	// Iterate through each square number
	for start := 1; start <= int(math.Sqrt(float64(max))); start++ {
		sum := start * start
		for next := start + 1; sum < max; next++ {
			sum += next * next
			if sum < max {
				digits := algebra.IntToDigits(sum)
				if util.IsPalindromeInt(digits) {
					if alreadySeen[sum] {
						fmt.Println("Found a duplicate!", sum)
					} else {
						totalSum += sum
						fmt.Printf("Palindrome sum: %10d, total sum: %12d\n", sum, totalSum)
						alreadySeen[sum] = true
					}
				}
			}
		}
	}

	return totalSum
}

func main() {
	fmt.Printf("Welcome to 125\n\n")

	maxFound := 100 * 1000 * 1000
	sum := looper(maxFound)
	fmt.Printf("\nSum of palindromic and consecutive squares < 10^8: %d\n\n", sum)
}
