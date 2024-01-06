package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// compare() determines which value is bigger.
// Returns -1 if b1,e1 is bigger, 0 if they are equal, and 1 if b2,e2 is bigger
func compare(b1 float64, e1 int, b2 float64, e2 int) (int, float64, int) {
	// fmt.Printf("%.2f^%d <=> %.2f^%d\n", b1, e1, b2, e2)
	if e1 < 2 && e2 < 2 {
		answer1 := math.Pow(b1, float64(e1))
		answer2 := math.Pow(b2, float64(e2))

		if answer1 > answer2 {
			return -1, b1, e1
		}
		if answer1 == answer2 {
			return 0, b1, e1
		}
		return 1, b2, e2
	}

	if b1 == b2 {
		if e1 == e2 {
			return 0, b1, e1
		}
		if e1 > e2 {
			return -1, b1, e1
		}
		return 1, b2, e2
	}

	if e1 == e2 {
		if b1 == b2 {
			return 0, b1, e1
		}
		if b1 > b2 {
			return -1, b1, e1
		}
		return 1, b2, e2
	}

	if b1 > b2 && e1 > e2 {
		return -1, b1, e1
	}

	if b2 > b1 && e2 > e1 {
		return 1, b2, e2
	}

	// Select the higher exponent. call that b1, e1.
	// Divide both powers by b1^e2
	// Subtract e2 from e1

	answer := 0
	if e1 > e2 {
		answer, _, _ = compare(b1, e1-e2, b2/b1, e2)
	} else {
		answer, _, _ = compare(b1/b2, e1, b2, e2-e1)
	}
	if answer == -1 {
		return -1, b1, e1
	}
	if answer == 0 {
		return 0, b1, e1
	}
	return 1, b2, e2

}

func readFile() [][2]int {
	raw, _ := os.ReadFile("p099_base_exp.txt")
	lines := strings.Split(string(raw), "\n")

	powers := make([][2]int, 0)

	for _, line := range lines {
		if line == "" {
			continue
		}
		be := strings.Split(line, ",")
		b, _ := strconv.Atoi(be[0])
		e, _ := strconv.Atoi(be[1])
		powers = append(powers, [2]int{b, e})
	}

	return powers
}

func max(powers [][2]int) int {
	b1 := 0.0
	e1 := 0
	maxLine := 0
	answer := 0

	for i := 0; i < len(powers); i++ {
		b2 := float64(powers[i][0])
		e2 := powers[i][1]
		answer, b1, e1 = compare(b1, e1, b2, e2)
		if answer == 1 {
			maxLine = i
		}
	}

	// Our array was zero-based, but the answer is going
	// to be 1-based. Add one to account for that.
	maxLine++

	return maxLine
}

func main() {
	fmt.Println("Welcome to 099")

	lines := readFile()
	maxLine := max(lines)
	fmt.Println("The max line is:", maxLine)
}
