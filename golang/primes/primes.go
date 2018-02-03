package primes

import (
	"fmt"
	"math"
)

const (
	MAX_PRIME = 1000 * 1000
)

var (
	Primes [MAX_PRIME + 1]bool
)

func Prime(number int) bool {
	return Primes[number]
}

func excludes(upper int, c chan int) {
	c <- 0
	c <- 1
	mid := int(math.Sqrt(float64(upper)))
	for i := 2; i <= mid; i++ {
		for j := i * 2; j <= upper; j += i {
			c <- j
		}
	}
	close(c)
}

func seive() {
	upper := MAX_PRIME
	fmt.Println("upper: ", upper)
	for i := 0; i <= upper; i++ {
		Primes[i] = true
	}
	c := make(chan int)
	go excludes(upper, c)
	for {
		exclude, ok := <-c
		if !ok {
			// Channel is empty
			return
		}
		Primes[exclude] = false
	}
}

func Init() {
	seive()
}
