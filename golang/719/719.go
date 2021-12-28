package main

// go fmt && golint && go test && go run 719.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"runtime/pprof"
	"strconv"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	combos     = [][][]int{}
)

// countBits returns a run length encoded count of the bits in mask.
func countBits(mask, numDigits int) []int {
	runs := []int{}
	bit := mask & 0x01
	count := 0

	for i := 0; i <= numDigits; i++ {
		if i == numDigits {
			runs = append(runs, count)
			break
		}
		if mask&(0x01<<i)>>i == bit {
			count++
			continue
		}
		bit = mask & (0x01 << i) >> i
		runs = append(runs, count)
		count = 1
	}

	return runs
}

// makeCombos returns all possible combinations of digits.
func makeCombos(numDigits int) [][]int {
	if numDigits <= 1 {
		return nil
	}

	combo := [][]int{}
	maxMask := int(math.Pow(2.0, float64(numDigits-1))) - 1

	for mask := 0; mask <= maxMask; mask++ {
		combo = append(combo, countBits(mask, numDigits))
	}

	return combo
}

// initCombos initializes the slice of digit combinations.
func initCombos(max int) {
	numDigits := int(math.Log10(float64(max))) + 1

	combos = make([][][]int, numDigits+1)

	for i := 1; i <= numDigits; i++ {
		combos[i] = makeCombos(i)
		fmt.Println(i, combos[i])
	}
}

// sNumber retuns true if square's root is the sum of square's digits.
func sNumber(root, square int) bool {
	numDigits := int(math.Log10(float64(square))) + 1

	digits := fmt.Sprintf("%d", square)

	// Try each combination of square's digits.
	for _, c := range combos[numDigits] {
		sum := 0
		i := 0
		for _, s := range c {
			t, _ := strconv.Atoi(digits[i : i+s])
			sum += t
			i += s
		}
		if sum == root {
			return true
		}
	}

	return false
}

// sSum returns the sum of all S-numbers from 1 to max.
func sSum(max int) int {
	initCombos(max)

	sum := 0
	for i := 1; i <= int(math.Sqrt(float64(max))); i++ {
		square := i * i
		if sNumber(i, square) {
			fmt.Println(square)
			sum += square
		}
	}

	return sum
}

func main() {
	fmt.Printf("Welcome to 719\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	max := 1000 * 1000 * 1000 * 1000

	fmt.Println("Sum of S numbers:", sSum(max))
}
