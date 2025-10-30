package main

// go fmt ./... && go vet ./... && go test ./... && go build 387.go && time ./387

import (
	"fmt"

	"github.com/erikbryant/util-golang/algebra"
	"github.com/erikbryant/util-golang/primey"
)

// A Harshad or Niven number is a number that is divisible by the sum of its digits.
// 201 is a Harshad number because it is divisible by 3 (the sum of its digits.)
// When we truncate the last digit from 201, we get 20, which is a Harshad number.
// When we truncate the last digit from 20, we get 2, which is also a Harshad number.
// Let's call a Harshad number that, while recursively truncating the last digit, always
// results in a Harshad number a right truncatable Harshad number.
//
// Also:
// 201/3=67 which is prime.
// Let's call a Harshad number that, when divided by the sum of its digits, results
// in a prime a strong Harshad number.
//
// Now take the number 2011 which is prime.
// When we truncate the last digit from it we get 201, a strong Harshad number that
// is also right truncatable.
// Let's call such primes strong, right truncatable Harshad primes.
//
// You are given that the sum of the strong, right truncatable Harshad primes less than 10000 is 90619.
//
// Find the sum of the strong, right truncatable Harshad primes less than 10^14.

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
	return n%sum == 0 && primey.Prime(n/sum)
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
				if primey.Prime(t) {
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

	maxFound := 100 * 1000 * 1000 * 1000 * 1000

	// Open channel and start go routine
	c := make(chan int, 10)
	go findRTH(maxFound, c)

	// Find the sum
	fmt.Printf("\nSum of strong, right truncatable Harshad primes < 10^14 = %d\n\n", sumSRTHP(maxFound, c))
}
