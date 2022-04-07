package main

// go fmt ./... && go vet ./... && go test && go run 119.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

import (
	"../util"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"runtime/pprof"
	"sort"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
)

// The number 512 is interesting because it is equal to the sum of its digits
// raised to some power: 5 + 1 + 2 = 8, and 83 = 512. Another example of a
// number with this property is 614656 = 284.
//
// We shall define a(n) to be the nth term of this sequence and insist that a
// number must contain at least two digits to have a sum.
//
// You are given that a(2) = 512 and a(10) = 614656.
//
// Find a(30).

// powerSums returns a sorted list of all numbers that are a power of the sum
// of their digits where the solutions are <= maxDigits in length
func powerSums(maxDigits int) []int {
	maxSum := maxDigits * 9
	maxD := int(math.Pow(10.0, float64(maxDigits))) - 1
	solutions := []int{}

	for base := 2; base <= maxSum; base++ {
		power := base
		for {
			power *= base
			if power <= 0 {
				fmt.Println("UNDERFLOW!!!!")
			}
			if power > maxD {
				break
			}
			if util.DigitSum(power) == base {
				solutions = append(solutions, power)
			}
		}
	}

	sort.Ints(solutions)

	return solutions
}

func main() {
	fmt.Printf("Welcome to 119\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	solutions := powerSums(16)

	fmt.Println("Len:", len(solutions))

	fmt.Println(solutions[1])
	fmt.Println(solutions[9])
	fmt.Println(solutions[29])
}
