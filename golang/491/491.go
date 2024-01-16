package main

// go fmt ./... && go vet ./... && go test && go run 491.go

// We call a positive integer double pandigital if it uses
// all the digits 0 to 9 exactly twice (with no leading zero).
// For example, 40561817703823564929 is one such number.
//
// How many double pandigital numbers are divisible by 11?

// **Sum of Alternating Digits**
//  If the ∑(digits in odd places) – ∑(digits in even places)
//  is divisible by 11 the original number is divisible by 11.
//
//   Example: n=2143
//    ∑(digits in odd places)  = 2 + 4 = 6
//    ∑(digits in even places) = 1 + 3 = 4
//    6 - 4 = 2
//    Thus n is not divisible by 11

// Find all the ways to split the 20 digits into two bins where
// the difference in the sums of the digits in the bins is a multiple
// of 11.
//
// For each way they can be split, calculate how many 20-digit
// permutations those ways would yield, excluding permutations where
// the leading digit is zero.
//
// The sum of the 20 digits is 90. The task is to choose ten
// digits whose sum will be s and where the difference between
// sum(s) and sum(90-s) is divisible by 11. In other words, all
// cases where abs(90-2s) % 11 == 0.
//
// Maximum sum = 9988776655 = 70
// Minimum sum = 0011223344 = 20
// 20 <= s  <= 70
// 20 <= s' <= 70
//
//   (s-s')    s   s'
//      0     45   45
//     11                 no solution (s is not an integer)
//     22     34   56
//     33                 no solution (s is not an integer)
//     44     23   67
//     55                 no solution (s is not an integer)
//     66     12   78     no solution (s is too small)
//     77                 no solution (s is not an integer)
//     88      1   89     no solution (s is too small)
//     99                 no solution (s is not an integer)
//
// Therefore, the digit partitionings to examine are:
//     (23, 67) (34, 56) (45, 45)
//   and their inverses;
//     (56, 34) (67, 23)
// These are the only partitionings that will yield a difference
// of a multiple of 11.
// Count all of these partitionings. Remove any that have a leading zero.

import (
	"fmt"
	"log"
	"math"
	"slices"
)

// permuteSum finds all permutations of the given digits that add to the given sum (including some duplicates)
func permuteSum(targetSum int, digitPool []int, currSum int, prefix []int, found *[][]int) {
	if currSum == targetSum {
		// Append a copy, so we don't clobber the slice later
		*found = append(*found, append([]int{}, prefix...))
		return
	}

	// Move forward in digitPool until we reach a digit we can use
	for i := 0; i < len(digitPool); i++ {
		if currSum+digitPool[i] <= targetSum {
			permuteSum(targetSum, digitPool[i+1:], currSum+digitPool[i], append(prefix, digitPool[i]), found)
		}
	}
}

// verify returns nil if the sum of the values in permutations equals targetSum
func verify(targetSum int, permutations [][]int) error {
	for _, p := range permutations {
		sum := 0
		for _, val := range p {
			sum += val
		}
		if sum != targetSum {
			return fmt.Errorf("sums did not match %d versus %v", targetSum, p)
		}
	}

	return nil
}

// fixLengthsAndZeroes removes permutations that are too long/short and adds missing zeroes
func fixLengthsAndZeroes(permutations [][]int) [][]int {
	pNew := [][]int{}

	for _, p := range permutations {
		// We assume no zeroes in the permutations; confirm that
		for _, val := range p {
			if val == 0 {
				panic("WTF!?")
			}
		}
		if len(p) == 10 {
			// Perfect! Save it.
			tmp := append([]int{}, p...)
			pNew = append(pNew, tmp)
			continue
		}
		if len(p) == 9 {
			// Add a zero and save it
			tmp := append([]int{}, p...)
			tmp = append(tmp, 0)
			pNew = append(pNew, tmp)
			continue
		}
		if len(p) == 8 {
			// Add two zeroes and save it
			tmp := append([]int{}, p...)
			tmp = append(tmp, 0)
			tmp = append(tmp, 0)
			pNew = append(pNew, tmp)
			continue
		}
	}

	return pNew
}

func removeDuplicatePermutations(p [][]int) [][]int {
	unique := [][]int{}
	dupes := map[int]bool{}

	for _, permutation := range p {
		asInt := makeInt(permutation)
		if dupes[asInt] {
			continue
		}
		dupes[asInt] = true
		unique = append(unique, permutation)
	}
	return unique
}

// permutations returns permutations of digitPool that add to targetSum
func permutations(targetSum int, digitPool []int) [][]int {
	perms := [][]int{}

	permuteSum(targetSum, digitPool, 0, []int{}, &perms)
	err := verify(targetSum, perms)
	if err != nil {
		log.Fatal(err)
	}
	perms = removeDuplicatePermutations(perms)
	perms = fixLengthsAndZeroes(perms)

	return perms
}

// otherHalf returns the remaining digits from digitPool
func otherHalf(digits, digitPool []int) []int {
	pool := map[int]int{}
	result := []int{}

	for _, digit := range digitPool {
		pool[digit]++
	}

	for _, digit := range digits {
		// digitPool is assumed to be a proper superset of digits
		if pool[digit] <= 0 {
			panic("Doh!")
		}
		pool[digit]--
	}

	for key, val := range pool {
		for i := 0; i < val; i++ {
			result = append(result, key)
		}
	}

	// Double check that we didn't drop/add any digits
	if len(result)+len(digits) != len(digitPool) {
		panic("Yikes!")
	}

	slices.Sort(result)

	return result
}

func makeInt(digits []int) int {
	result := 0

	for _, digit := range digits {
		result *= 10
		result += digit
	}

	return result
}

func pairCount(digits []int) (int, int) {
	dupes := map[int]int{}

	for _, d := range digits {
		dupes[d]++
	}

	pairs := 0
	for _, val := range dupes {
		if val == 2 {
			pairs++
		}
	}

	return pairs, dupes[0]
}

func factorial(n int) int {
	product := 1

	for n > 0 {
		product *= n
		n -= 1
	}

	return product
}

// comboCompute returns how many combinations would result from the given digits
func comboCompute(digits []int, ignoreLeadingZeroes bool) int {
	// Given l=len(digits), combinations = l!
	// If ignoreLeadingZeroes
	//   One zero   combinations = 9 * (l-1)!
	//   Two zeroes combinations = 8 * (l-1)!

	// If there are pairs of duplicates, each additional pair
	// cuts the combinations in half.

	l := len(digits)
	pairs, zeroCount := pairCount(digits)
	if !ignoreLeadingZeroes {
		zeroCount = 0
	}
	combos := (l - zeroCount) * factorial(l-1) / int(math.Pow(2.0, float64(pairs)))

	return combos
}

// combinations returns the number of unique combinations of the given digits
func combinations(digitsOdd, digitsEven []int) int {
	lenOdd := comboCompute(digitsEven, true)
	lenEven := comboCompute(digitsEven, false)
	return lenOdd * lenEven
}

func main() {
	digitPool := []int{9, 9, 8, 8, 7, 7, 6, 6, 5, 5, 4, 4, 3, 3, 2, 2, 1, 1, 0, 0}

	divisibleCount := 0

	// The digit partitionings to examine are:
	//   (23, 67) (34, 56) (45, 45) (56, 34) (67, 23)
	for _, targetSum := range []int{23, 34, 45, 56, 67} {
		p := permutations(targetSum, digitPool)
		fmt.Println("targetSum:", targetSum, "#permutations:", len(p))
		for _, pn := range p {
			o := otherHalf(pn, digitPool)
			divisibleCount += combinations(pn, o)
		}
	}

	fmt.Println("# of double pandigital numbers divisible by 11:", divisibleCount)
}
