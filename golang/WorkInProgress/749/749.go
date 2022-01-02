package main

// go fmt && golint && go test && go run 749.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

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

// I'm stuck with a too-slow algorithm. Optimizing is unlikely to make it
// fast enough. I think it needs to be a new algorithm. Next steps:
//
// * Generate near power sums instead of searching for them
// * Work out the equations and see if there is a way to find solutions
//     a*10 + b + 1 = a^k + b^k
//        or
//     a*10 + b - 1 = a^k + b^k
//
//     a*100 + b*10 + c + 1 = a^k + b^k + c^k
//        or
//     a*100 + b*10 + c - 1 = a^k + b^k + c^k

// Results:
//
// 1 0
//
//               35          2
//               75          2
// 2 110
//
// 3 0
//
// 4 0
//
// 5 0
//
//           528757          6
//           629643          6
//           688722          6
//           715469          6
// 6 2562591
//
// 7 0
//
//         30405525         10
//         31672867          8
//         44936324          8
//         63645890          8
//         63645891          8
//         71419078          8
//         73495876          8
// 8 379221451
//
//        116079879          8
//        647045075         10
// 9 763124954
//
//       1136483324         10
// 10 1136483324
//
//      83311557354         12
// 11 83311557354

// generatePowers populates 'powers' with powers of each digit.
func generatePowers() {
	digits := 16

	powers = make([][]int, 10)

	for digit := 0; digit <= 9; digit++ {
		// For 16 digit numbers 54 is the highest power.
		powers[digit] = make([]int, 54+1)
	}

	max := int(math.Pow(10.0, float64(digits))) - 1

	for digit := 2; digit <= 9; digit++ {
		val := 0
		for power := 1; val < max; power++ {
			val = int(math.Pow(float64(digit), float64(power)))
			powers[digit][power] = val
			powers[0][power] = 0
			powers[1][power] = 1
		}
	}
}

// nearPowerSum returns whether 'n' is a near power sum.
func nearPowerSum(n int, digits []int, minPower int) bool {
	// If the only digits are zeroes and ones then there are no solvable powers.
	// Exit or the loop below will not terminate.
	if digits[0] == 1 {
		powerable := false
		for _, d := range digits {
			if d > 1 {
				powerable = true
				break
			}
		}
		if !powerable {
			return false
		}
	}

	sum := 0
	// Stop at n/2, as no power++ will be less than that distance.
	for power := minPower; sum < n>>1; power++ {
		sum = 0
		for _, d := range digits {
			sum += powers[d][power]
		}
		if sum == n-1 || sum == n+1 {
			fmt.Printf("%16d %10d\n", n, power)
			return true
		}
	}

	return false
}

// increment increments a slice of digits and returns it.
func increment(digits []int, d int) []int {
	d--

	digits[d]++

	// Carry, if needed.
	for d > 0 {
		if digits[d] <= 9 {
			break
		}
		digits[d] = 0
		d--
		digits[d]++
	}

	return digits
}

// s2 returns the sum of all near power sums of exactly 'd' digits.
func s2(d int) int {
	sum := 0

	digits := make([]int, d)

	digits[0] = 1
	n := int(math.Pow(10.0, float64(d-1)))
	max := int(math.Pow(10.0, float64(d))) - 1

	// The default case I was able to prove. The special cases are cheats based
	// on results of previous runs (until I can prove them, too).
	minPower := d - 1
	switch d {
	case 2:
		minPower = 2
	case 6:
		minPower = 6
	case 8:
		minPower = 8
	case 10:
		minPower = 10
	case 11:
		minPower = 12
	}

	for ; n < max; n++ {
		if nearPowerSum(n, digits, minPower) {
			sum += n
		}
		digits = increment(digits, d)
	}

	return sum
}

// s returns the sum of all near power sums of 'd' or fewer digits.
func s(d int) int {
	sum := 0

	for l := 1; l <= d; l++ {
		sum += s2(l)
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

	// For each number of digit length i, find its near power sum.
	for i := 1; i <= 9; i++ {
		fmt.Println(i, s2(i))
		fmt.Println()
	}
}
