package main

import (
	"../primes"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
)

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

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func init() {
	primes.Load("../primes.gob")
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

// SamDisplayCost() calculates the power consumed to transition from a blank display to displaying 'next' and back to a blank display.
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

// MaxDisplayCost() calculates the power consumed to transition from an 'initial' display state to the 'next' display state turning on/off only those segments that change.
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
		if primes.Prime(i) {
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
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	// The two clocks are fed all the prime numbers between A = 107 and B = 2×107.
	samCost, maxCost := costComparison(1000*1000*10, 1000*1000*20)

	fmt.Println("Sam:", samCost, "Max:", maxCost, "Sam-Max:", samCost-maxCost)
}
