package main

import (
	"../primes"
	"flag"
	"fmt"
	"log"
	// "math"
	"os"
	"runtime/pprof"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	cacheRTH   = map[int]bool{}
	maxRTH     = 0
)

func init() {
	primes.Load("../primes.gob")
	cacheRTH = make(map[int]bool)
}

// digitSum returns the sum of the digits in the number.
func digitSum(n int) (sum int) {
	for n > 0 {
		sum += n % 10
		n /= 10
	}

	return
}

// harshad returns true if n is divisible by the sum of its digits.
func harshad(n int) bool {
	return n%digitSum(n) == 0
}

// rightTruncatableHarshad returns true if n is a right truncatable harshad.
// There are no truncatable values below 10, so don't call this if n < 10.
func rightTruncatableHarshad(n int) bool {
	if !harshad(n) {
		return false
	}

	if n < 100 || rightTruncatableHarshad(n/10) {
		return true
	}

	return false
}

// strong returns true if n divided by the sum of its digits is prime.
func strong(n int) bool {
	sum := digitSum(n)

	// Only check for prime if it divides evenly.
	return n%sum == 0 && primes.Prime(n/sum)
}

// strongRightTruncatableHarshad returns true if n is strong and is right truncatable.
func strongRightTruncatableHarshad(n int) bool {
	return rightTruncatableHarshad(n) && strong(n)
}

// strongRightTruncatableHarshadPrime returns true if p is prime and the first truncation is a strong right truncatable Harshad.
func strongRightTruncatableHarshadPrime(p int) bool {
	return strongRightTruncatableHarshad(p/10) && primes.Prime(p)
}

// sumSRTHP returns the sum of strong right truncatable Harshad primes <= max.
func sumSRTHP(max int) int {
	// There are no SRTHP below 10 and starting at 10 eliminates edge cases.
	i := 10

	sum := 0

	for {
		if rightTruncatableHarshad(i) {
			for _, t := range []int{
				1 + i*10,
				3 + i*10,
				5 + i*10,
				7 + i*10,
				9 + i*10,
			} {
				if t > max {
					return sum
				}
				if strong(t/10) && primes.Prime(t) {
					fmt.Println(t)
					sum += t
				}
			}
		}
		i++
	}

	return sum
}

func main() {
	fmt.Printf("Welcome to 387\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	fmt.Println("Sum: ", sumSRTHP(100*1000*1000*1000))
	// fmt.Println("Sum: ", sumSRTHP(100*1000*1000*1000*1000))
}
