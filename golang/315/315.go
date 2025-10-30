package main

// go fmt ./... && go vet ./... && go test ./... && go build 315.go && time ./315

import (
	"fmt"
	"strconv"

	"github.com/erikbryant/util-golang/primey"
)

// https://projecteuler.net/resources/images/0315_clocks.gif?1678992056
//
// Sam and Max are asked to transform two digital clocks into two "digital root" clocks.
// A digital root clock is a digital clock that calculates digital roots step by step.
//
// When a clock is fed a number, it will show it, and then it will start the calculation,
// showing all the intermediate values until it gets to the result.
// For example, if the clock is fed the number 137, it will show: "137" → "11" → "2" and
// then it will go black, waiting for the next number.
//
// Every digital number consists of some light segments: three horizontal (top, middle, bottom)
// and four vertical (top-left, top-right, bottom-left, bottom-right).
// Number "1" is made of vertical top-right and bottom-right, number "4" is made by middle
// horizontal and vertical top-left, top-right and bottom-right. Number "8" lights them all.
//
// The clocks consume energy only when segments are turned on/off.
// To turn on a "2" will cost 5 transitions, while a "7" will cost only 4 transitions.
//
// Sam and Max built two different clocks.
//
// Sam's clock is fed e.g. number 137: the clock shows "137", then the panel is turned off,
// then the next number ("11") is turned on, then the panel is turned off again and finally
// the last number ("2") is turned on and, after some time, off.
// For the example, with number 137, Sam's clock requires:
//
// "137": (2 + 5 + 4) × 2 = 22 transitions ("137" on/off).
//  "11": (2 + 2) × 2 = 8 transitions ("11" on/off).
//   "2": (5) × 2 = 10 transitions ("_2_" on/off).
// For a grand total of 40 transitions.
//
// Max's clock works differently. Instead of turning off the whole panel, it is smart enough
// to turn off only those segments that won't be needed for the next number.
// For number 137, Max's clock requires:
//
//"137": 2 + 5 + 4 = 11 transitions ("137" on)
//       7 transitions (to turn off the segments that are not needed for number "11").
// "11": 0 transitions (number "11" is already turned on correctly)
//       3 transitions (to turn off the first "1" and the bottom part of the second "1";
//         the top part is common with number "_2_").
//  "2": 4 transitions (to turn on the remaining segments in order to get a "_2_")
//       5 transitions (to turn off number "_2_").
// For a grand total of 30 transitions.
//
// Of course, Max's clock consumes less power than Sam's one.
// The two clocks are fed all the prime numbers between A = 107 and B = 2×107.
// Find the difference between the total number of transitions needed by Sam's clock and
// that needed by Max's one.

// DigitCosts is the count of the number of segments in that digital number.
var DigitCosts = [10]int{
	6, // 0
	2, // 1
	5, // 2
	5, // 3
	4, // 4
	5, // 5
	6, // 6
	4, // 7
	7, // 8
	6, // 9
}

// TransitionCosts is the cost to switch the display from one digit to the other digit using Max's least-cost method.
var TransitionCosts = [10][10]int{
	//  1  2  3  4  5  6  7  8  9
	{0, 4, 3, 3, 4, 3, 2, 2, 1, 2}, // 0 -> ...
	{4, 0, 5, 3, 2, 5, 6, 2, 5, 4}, // 1 -> ...
	{3, 5, 0, 2, 5, 4, 3, 5, 2, 3}, // 2 -> ...
	{3, 3, 2, 0, 3, 2, 3, 3, 2, 1}, // 3 -> ...
	{4, 2, 5, 3, 0, 3, 4, 2, 3, 2}, // 4 -> ...
	{3, 5, 4, 2, 3, 0, 1, 3, 2, 1}, // 5 -> ...
	{2, 6, 3, 3, 4, 1, 0, 4, 1, 2}, // 6 -> ...
	{2, 2, 5, 3, 2, 3, 4, 0, 3, 2}, // 7 -> ...
	{1, 5, 2, 2, 3, 2, 1, 3, 0, 1}, // 8 -> ...
	{2, 4, 3, 1, 2, 1, 2, 2, 1, 0}, // 9 -> ...
}

// digitalRoot() calculates the next number in the series of digital roots for 'n'.
// The digital root of a number is the sum of its digits, repeated until the sum
// has only one digit:
// 137 -> 11 -> 2
func digitalRoot(n string) string {
	sum := 0

	for _, digit := range n {
		sum += int(digit) - 48
	}

	return strconv.Itoa(sum)
}

// digitalRoots() calculates the series of numbers that make up the digital root of 'n'.
func digitalRoots(n string) []string {
	roots := []string{n}

	for len(n) > 1 {
		n = digitalRoot(n)
		roots = append(roots, n)
	}

	return roots
}

// SamDisplayCost calculates the power consumed to transition from a blank display to displaying 'next' and back to a blank display.
func SamDisplayCost(next string) int {
	cost := 0

	// Turn on the 'next' value.
	for _, digit := range next {
		cost += DigitCosts[int(digit)-48]
	}

	// Turn off the 'next' value.
	cost *= 2

	return cost
}

// MaxDisplayCost calculates the power consumed to transition from an 'initial' display state to the 'next' display state turning on/off only those segments that change.
func MaxDisplayCost(initial, next string) int {
	cost := 0

	// Turn off the excess 'initial' runes.
	for len(initial) > len(next) {
		cost += DigitCosts[initial[0]-48]
		initial = initial[1:]
	}

	// Turn on the excess 'next' runes.
	for len(next) > len(initial) {
		cost += DigitCosts[next[0]-48]
		next = next[1:]
	}

	// They are now of equal length. Transition the remaining runes.
	for i := 0; i < len(initial); i++ {
		cost += TransitionCosts[initial[i]-48][next[i]-48]
	}

	return cost
}

// costComparison() finds the total number of transitions needed by
// Sam's clock and that needed by Max's clock for a given range.
func costComparison(start, end int) (int, int) {
	samCost := 0
	maxCost := 0

	for i := start; i <= end; i++ {
		if primey.Prime(i) {
			last := ""
			for _, root := range digitalRoots(strconv.Itoa(i)) {
				samCost += SamDisplayCost(root)
				maxCost += MaxDisplayCost(last, root)
				last = root
			}
			maxCost += MaxDisplayCost(last, "")
		}
	}

	return samCost, maxCost
}

func main() {
	fmt.Printf("Welcome to 315\n\n")

	// The two clocks are fed all the prime numbers between A = 107 and B = 2×107.
	samCost, maxCost := costComparison(1000*1000*10, 1000*1000*20)

	fmt.Println("Transitions needed for Sam =", samCost, "Max =", maxCost)
	fmt.Printf("Difference between transitions: %d\n\n", samCost-maxCost)
}
