package main

import (
	"fmt"
	"strconv"
)

type pair struct {
	dividend int
	divisor  int
}

func longDivision(numerator, denominator int) (answer string, cycleLen int) {
	const maxDigits = 10000

	seen := make(map[pair]int)

	dividend := numerator
	divisor := denominator
	loopCount := 0

	answer = ""

	for {
		if dividend < divisor {
			dividend *= 10
		}

		remainder := dividend % divisor
		quotient := dividend / divisor
		answer += strconv.Itoa(quotient)

		// Termination conditions:
		// * The division is complete (no remainder)
		// * A cycle is detected (we have see the dividend and the divisor before)
		// * We have generated > maxDigits
		if remainder == 0 {
			cycleLen = 0
			break
		}
		if loopCount > maxDigits {
			fmt.Println("ERROR: timed out searching for a repetition")
			cycleLen = maxDigits
			break
		}
		index, ok := seen[pair{dividend: dividend, divisor: divisor}]
		if ok {
			cycleLen = loopCount - index
			answer = answer[index:loopCount]
			break
		}
		seen[pair{dividend: dividend, divisor: divisor}] = loopCount

		dividend = remainder
		loopCount++
	}

	return answer, cycleLen
}

func main() {
	maxLen := 0
	maxD := 0

	for d := 2; d < 1000; d++ {
		_, x := longDivision(1, d)
		if x > maxLen {
			maxLen = x
			maxD = d
		}
	}

	fmt.Println("Longest cycle length:", maxLen, "at d =", maxD)
}
