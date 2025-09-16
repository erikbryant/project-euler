package main

// go fmt ./... && go vet ./... && go test && go build 118.go && time ./118

// Using all the digits 1 through 9 and concatenating them freely to
// form decimal integers, different sets can be formed. Interestingly with
// the set {2,5,47,89,631}, all the elements belonging to it are prime.
//
// How many distinct sets containing each of the digits one through nine
// exactly once contain only prime elements?

import (
	"fmt"
	"math"

	"github.com/erikbryant/util-golang/primes"
	"github.com/erikbryant/util-golang/util"
)

func dupeDigits(n int) bool {
	digits := map[int]bool{
		0: true, // Only looking for 1-9 pandigital
	}

	for n > 0 {
		_, ok := digits[n%10]
		if ok {
			return true
		}
		digits[n%10] = true
		n /= 10
	}

	return false
}

func digitCount(n int) int {
	return int(math.Trunc(math.Log10(float64(n)))) + 1
}

func loadPrimes() [][]int {
	candidates := [][]int{
		{}, // primes of length 0 digits (there will be none here)
		{}, // 1 digit
		{}, // 2   "
		{}, // 3   "
		{}, // 4   "
		{}, // 5   "
		{}, // 6   "
		{}, // 7   "
		{}, // 8   "
		{}, // 9   "
	}

	for _, p := range primes.PackedPrimes {
		digits := digitCount(p)
		if digits > 9 {
			break
		}
		if dupeDigits(p) {
			continue
		}
		candidates[digits] = append(candidates[digits], p)
	}

	for digits, p := range candidates {
		fmt.Printf("#candidates[%d] : %7d\n", digits, len(p))
	}

	return candidates
}

func digitsNeeded(have []int) map[int]bool {
	need := map[int]bool{}
	for i := 1; i <= 9; i++ {
		need[i] = true
	}

	for _, n := range have {
		for n > 0 {
			delete(need, n%10)
			n /= 10
		}
	}

	return need
}

func pandigitalSet(s []int) bool {
	digits := 0
	for _, n := range s {
		digits += digitCount(n)
	}

	return digits == 9 && len(digitsNeeded(s)) == 0
}

func makeSets(primesByDigit [][]int, pattern []int, set []int, sets [][]int) [][]int {
	if len(pattern) == 0 {
		if pandigitalSet(set) {
			sets = append(sets, set)
		}
		return sets
	}

	digit := pattern[0]
	for i := 0; i < len(primesByDigit[digit]); i++ {
		if len(set) > 0 {
			if set[len(set)-1] <= primesByDigit[digit][i] {
				continue
			}
		}
		set := append(set, primesByDigit[digit][i])
		sets = makeSets(primesByDigit, pattern[1:len(pattern)], set, sets)
	}

	return sets
}

func main() {
	fmt.Printf("Welcome to 118\n\n")

	sets := [][]int{}
	primesByDigit := loadPrimes()

	for _, pattern := range util.Partitions(9) {
		sets = makeSets(primesByDigit, pattern, []int{}, sets)
		fmt.Printf("Pattern: %v sets: %d\n", pattern, len(sets))
	}

	fmt.Printf("\nTotal pandigital sets: %d\n", len(sets))
}
