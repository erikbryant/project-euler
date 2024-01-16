package main

// go fmt ./... && go vet ./... && go test && go run 749.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"runtime/pprof"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	powers     = [][]int{}
)

//
// A positive integer, n, is a near power sum if there exists a positive
// integer, k, such that the sum of the kth powers of the digits in its decimal
// representation is equal to either n+1 or n-1. For example 35 is a near power
// sum number because 3^2+5^2=34.
//
// Define S(d) to be the sum of all near power sum numbers of d digits or less.
// Then S(2)=110 and S(6)=2,562,701.
//
// Find S(16).
//

// generatePowers populates 'powers' with powers of each digit.
func generatePowers() {
	digits := 16

	powers = make([][]int, 10)

	for digit := 0; digit <= 9; digit++ {
		// For 16-digit numbers 54 is the highest power.
		powers[digit] = make([]int, 54+1)
	}

	maxFound := int(math.Pow(10.0, float64(digits))) - 1

	for digit := 2; digit <= 9; digit++ {
		val := 0
		for power := 1; val < maxFound; power++ {
			val = int(math.Pow(float64(digit), float64(power)))
			powers[digit][power] = val
			powers[0][power] = 0
			powers[1][power] = 1
		}
	}
}

// solvable returns true if at least one of the digits is not zero or one
func solvable(digits []int) bool {
	// If the only digits are zeroes and ones then there are no solvable powers
	for _, d := range digits {
		if d > 1 {
			return true
		}
	}

	return false
}

// canMakeSumFromDigits returns true if the digits raised to k sum to a permutation of the digits
func canMakeSumFromDigits(digits []int, sum int) bool {
	s := sum
	dd := make([]int, len(digits))
	copy(dd, digits)

	for s > 0 && len(dd) > 0 {
		digit := s % 10
		s /= 10
		found := false
		for i := range dd {
			if dd[i] == digit {
				found = true
				dd = append(dd[:i], dd[i+1:]...)
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

// powerSums finds each power sum for a given set of digits
func powerSums(digits []int) int {
	if !solvable(digits) {
		return 0
	}

	total := 0

	d := len(digits) - 1
	minPower := d - 1
	minSum := int(math.Pow(10.0, float64(d)))
	maxSum := int(math.Pow(10.0, float64(d+1))) - 1

	sum := 0
	for power := minPower; ; power++ {
		sum = 0
		for i := 0; i <= d; i++ {
			digit := digits[i]
			sum += powers[digit][power]
		}
		if sum < minSum {
			continue
		}
		if sum > maxSum {
			break
		}
		if sum-1 > minSum && canMakeSumFromDigits(digits, sum-1) {
			total += sum - 1
		}
		if sum+1 < maxSum && canMakeSumFromDigits(digits, sum+1) {
			total += sum + 1
		}
	}

	return total
}

// increment increments a slice of d digits and returns it
func increment(digits []int) []int {
	// This is not a traditional increment. Instead, we are cycling through
	// all possible unique sets of digits. The pattern looks like:
	//
	// 0 0 0 1
	// 0 0 0 2
	// ...
	// 0 0 0 9
	// 0 0 1 1
	// 0 0 1 2
	// ...

	lastD := len(digits) - 1

	if digits[lastD] < 9 {
		digits[lastD]++
		return digits
	}

	// Carry
	var d int
	for d = lastD; d >= 0 && digits[d] == 9; d-- {
	}
	digits[d]++
	carry := digits[d]
	for ; d <= lastD; d++ {
		digits[d] = carry
	}

	return digits
}

// digitCombinations finds each combination of d digits (ignoring duplicates)
func digitCombinations(d int) int {
	sum := 0

	digits := make([]int, d)

	digits[d-1] = 1

	for digits[0] < 9 {
		sum += powerSums(digits)
		increment(digits)
	}

	return sum
}

func main() {
	fmt.Printf("Welcome to 749\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	generatePowers()

	sum := 0
	for d := 2; d <= 16; d++ {
		sum += digitCombinations(d)
		fmt.Println(d, sum)
	}
}
