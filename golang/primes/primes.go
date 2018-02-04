package primes

// https://primes.utm.edu/howmany.html
//
//                           x                pi(x)
// 1                        10                    4
// 2                       100                   25
// 3                     1,000                  168
// 4                    10,000                1,229
// 5                   100,000                9,592
// 6                 1,000,000               78,498
// 7                10,000,000              664,579
// 8               100,000,000            5,761,455
// 9	         1,000,000,000           50,847,534
// 10	        10,000,000,000          455,052,511
// 11	       100,000,000,000        4,118,054,813
// 12	     1,000,000,000,000       37,607,912,018
// 13	    10,000,000,000,000      346,065,536,839
// 14	   100,000,000,000,000    3,204,941,750,802
// 15	 1,000,000,000,000,000   29,844,570,422,669
// 16	10,000,000,000,000,000  279,238,341,033,925

import (
	"fmt"
	"math"
)

const (
	MAX_PRIME = 1000*1000*10 + 1000
)

var (
	Primes          [MAX_PRIME + 1]bool
	PackedPrimes    [MAX_PRIME + 1]int
	packedPrimesEnd int
)

func Prime(number int) bool {
	return number == PackedPrimes[PackedIndex(number)]
}

func packPrimes() {
	j := 0
	for i := 0; i < len(Primes); i++ {
		if Primes[i] {
			PackedPrimes[j] = i
			j++
		}
	}
	packedPrimesEnd = j - 1
}

func PackedIndex(n int) int {
	upper := packedPrimesEnd
	lower := 0

	for upper > lower {
		mid := (upper + lower) >> 1

		if n > PackedPrimes[mid] {
			if n < PackedPrimes[mid+1] {
				return mid
			}
			lower = mid + 1
		} else {
			if n == PackedPrimes[mid] {
				return mid
			}
			upper = mid - 1
		}

	}
	return upper
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
			break
		}
		Primes[exclude] = false
	}
}

func Init() {
	seive()
	packPrimes()
	fmt.Println("primes.Init() complete")
}
