package main

// go fmt ./... && go vet ./... && go test ./... && go build 046.go && time ./046

import (
	"fmt"

	"github.com/erikbryant/util-golang/primey"
)

// It was proposed by Christian Goldbach that every odd composite number can be written
// as the sum of a prime and twice a square.
//
// 9 = 7 + 2 x 1^2
// 15 = 7 + 2 x 2^2
// 21 = 3 + 2 x 3^2
// 25 = 7 + 2 x 3^2
// 27 = 19 + 2 x 2^2
// 33 = 31 + 2 x 1^2
//
// It turns out that the conjecture was false.
// What is the smallest odd composite that cannot be written as the sum of a prime and
// twice a square?

// goldbach returns prime, square if composite can be written as the sum of a prime and twice a square; 0, 0 otherwise
func goldbach(composite int) (prime, square int) {
	i := 1
	for {
		s := 2 * i * i
		if s >= composite {
			break
		}
		p := composite - s
		if primey.Prime(p) {
			return p, i
		}
		i++
	}

	return 0, 0
}

func oddComposite(c chan int) {
	defer close(c)

	// Return the odd, non-prime numbers
	for i := 3; i <= 1000*1000; i += 2 {
		if !primey.Prime(i) {
			c <- i
		}
	}
}

func main() {
	fmt.Println("Welcome to 046")
	c := make(chan int)
	go oddComposite(c)

	for {
		composite, ok := <-c
		if !ok {
			break
		}
		prime, square := goldbach(composite)
		if prime == 0 && square == 0 {
			fmt.Println("Composite is not Goldbachian:", composite)
			break
		}
	}
}
