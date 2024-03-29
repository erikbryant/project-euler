package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/erikbryant/util-golang/algebra"
	"github.com/erikbryant/util-golang/primes"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
)

// rightTruncatableHarshad returns true if n is a right truncatable harshad.
// There are no truncatable values below 10, so don't call this if n < 10.
func rightTruncatableHarshad(n, sum int) bool {
	if n%sum != 0 {
		// This is not a Harshad number
		return false
	}

	if n < 100 {
		return true
	}

	n /= 10
	return rightTruncatableHarshad(n, algebra.DigitSum(n))
}

// strong returns true if n divided by the sum of its digits is prime.
func strong(n, sum int) bool {
	// Only check for prime if it divides evenly. Otherwise, we get false positives.
	return n%sum == 0 && primes.Prime(n/sum)
}

// sumSRTHP returns the sum of strong right truncatable Harshad primes <= max.
func sumSRTHP(max int, c chan int) int {
	sum := 0

	for {
		done := false

		// Read a harshad number from channel.
		h, ok := <-c
		if !ok {
			break
		}
		if strong(h, algebra.DigitSum(h)) {
			for _, t := range []int{
				1 + h*10,
				3 + h*10,
				7 + h*10,
				9 + h*10,
			} {
				if t > max {
					done = true
					break
				}
				if primes.Prime(t) {
					fmt.Println(t)
					sum += t
				}
			}
		}
		if done {
			break
		}
	}

	return sum
}

// findTH finds right truncatable harshad numbers and sends them to a channel
func findRTH(max int, c chan int) {
	defer close(c)

	// Seed the queue with the first known values.
	queue := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	i := 0
	for {
		done := false
		rth := queue[i]
		queue = queue[1:]

		// Push 'rth' to channel.
		c <- rth

		sum := algebra.DigitSum(rth)
		for d := 0; d <= 9; d++ {
			c := rth*10 + d
			if c > max {
				done = true
				break
			}
			if rightTruncatableHarshad(c, sum+d) {
				queue = append(queue, c)
			}
		}

		if done {
			break
		}
	}

	// Flush the rest of the queue.
	for _, rth := range queue {
		c <- rth
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

	maxFound := 100 * 1000 * 1000 * 1000 * 1000

	// Open channel and start go routine
	c := make(chan int, 10)
	go findRTH(maxFound, c)

	// Find the sum
	fmt.Println("Sum: ", sumSRTHP(maxFound, c))
}
