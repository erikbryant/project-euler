package main

// go fmt ./... && go vet ./... && go run 205.go

import (
	"fmt"
)

// Peter has nine four-sided dice, each with faces numbered 1, 2, 3, 4.
// Colin has six six-sided dice, each with faces numbered 1, 2, 3, 4, 5, 6.
//
// Peter and Colin roll their dice and compare totals: the highest total wins.
// The result is a draw if the totals are equal.
//
// What is the probability that Pyramidal Pete beats Cubic Colin? Give your
// answer rounded to seven decimal places in the form 0.abcdefg

// sumDice returns the sum of the dice in the slice.
func sumDice(dice []int) int {
	sum := 0

	for i := 0; i < len(dice); i++ {
		sum += dice[i]
	}

	return sum
}

// countRolls returns a map of how many different rolls can make each sum.
func countRolls(d, sides int) map[int]int {
	rolls := make(map[int]int)
	dice := make([]int, d)

	for i := 0; i < d; i++ {
		dice[i] = 1
	}
	rolls[sumDice(dice)]++

	// Find each combination of dice in a single roll.
	for {
		dice[0]++
		// Carry
		i := 0
		for {
			if dice[i] <= sides {
				break
			}
			dice[i] = 1
			i++
			if i >= d {
				return rolls
			}
			dice[i]++
		}
		rolls[sumDice(dice)]++
	}

	return rolls
}

func main() {
	fmt.Printf("Welcome to 205\n\n")

	pRolls := countRolls(9, 4)
	fmt.Println(pRolls)
	cRolls := countRolls(6, 6)
	fmt.Println(cRolls)

	pCombos := 0
	for _, i := range pRolls {
		pCombos += i
	}

	cCombos := 0
	for _, i := range cRolls {
		cCombos += i
	}

	totalOutcomes := pCombos * cCombos

	cMin := 6 * 1
	pMax := 9 * 4

	pWins := 0
	for i := cMin + 1; i <= pMax; i++ {
		cLosses := 0
		for j := 0; j < i; j++ {
			cLosses += cRolls[j]
		}
		pWins += pRolls[i] * cLosses
	}

	fmt.Println("Combinations Peter can roll:", pCombos)
	fmt.Println("Combinations Colin can roll:", cCombos)
	fmt.Println("Total outcomes:", totalOutcomes)
	fmt.Println("Peter win chance: ", float64(pWins)/float64(totalOutcomes))
}
