package main

import (
	"../primes"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"runtime/pprof"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
)

func init() {
	primes.Load("../primes.gob")
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
func harshad(n, sum int) bool {
	return n%sum == 0
}

// rightTruncatableHarshad returns true if n is a right truncatable harshad.
// There are no truncatable values below 10, so don't call this if n < 10.
func rightTruncatableHarshad(n, sum int) bool {
	if !harshad(n, sum) {
		return false
	}

	if n < 100 {
		return true
	}

	n /= 10
	return rightTruncatableHarshad(n, digitSum(n))
}

// strong returns true if n divided by the sum of its digits is prime.
func strong(n, sum int) bool {
	// Only check for prime if it divides evenly. Otherwise we get false positives.
	return n%sum == 0 && primes.Prime(n/sum)
}

// sumSRTHP returns the sum of strong right truncatable Harshad primes <= max.
func sumSRTHP(max int) int {
	// We start searching at 200. That skips 181. Account for it manually.
	sum := 181
	power := 1.0

	for {
		// Results only begin with 2, 4, 6, or 8. But, there are so few 6's
		// that we'll just add those in later.
		for _, base := range []int{2, 4, 8} {
			start := base * int(math.Pow(10, power))
			end := (base + 1) * int(math.Pow(10, power))
			for i := start; i < end; i++ {
				d := digitSum(i)
				if strong(i, d) && rightTruncatableHarshad(i, d) {
					for _, t := range []int{
						1 + i*10,
						3 + i*10,
						7 + i*10,
						9 + i*10,
					} {
						if t > max {
							// Add in the results that begin with 6.
							for _, s := range []int{631, 6037, 60000000037} {
								if max >= s {
									sum += s
								}
							}
							return sum
						}
						if primes.Prime(t) {
							fmt.Println(t)
							sum += t
						}
					}
				}
			}
		}
		power += 1.0
	}
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
