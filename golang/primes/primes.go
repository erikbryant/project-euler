package primes

// https://primes.utm.edu/howmany.html
//
//                         x                pi(x)
// 1                      10                    4
// 2                     100                   25
// 3                   1,000                  168
// 4                  10,000                1,229
// 5                 100,000                9,592
// 6               1,000,000               78,498
// 7              10,000,000              664,579
// 8             100,000,000            5,761,455
// 9	         1,000,000,000           50,847,534
// 10	        10,000,000,000          455,052,511
// 11	       100,000,000,000        4,118,054,813
// 12	     1,000,000,000,000       37,607,912,018
// 13	    10,000,000,000,000      346,065,536,839
// 14	   100,000,000,000,000    3,204,941,750,802
// 15	 1,000,000,000,000,000   29,844,570,422,669
// 16	10,000,000,000,000,000  279,238,341,033,925

import (
	"encoding/gob"
	"fmt"
	"math"
	"os"
)

const (
	// MaxPrime is the highest value up to which we will search for primes
	MaxPrime = 100*1000*1000 + 1000
)

var (
	// Primes indicates whether the index is prime or not
	Primes []bool
	// PackedPrimes is a list of the first n prime numbers
	PackedPrimes []int
	// PackedPrimesEnd is the index of the final value ini the PackedPrimes slice
	PackedPrimesEnd int
)

// Pi is the prime counting function, returning the number of primes below n
// https://en.wikipedia.org/wiki/Prime-counting_function
func Pi(n int) int {
	if n < PackedPrimes[0] {
		return 0
	}

	if n > PackedPrimes[PackedPrimesEnd] {
		err := fmt.Errorf("Pi(%d) exceeded max prime. Did you call Init()?", n)
		panic(err)
	}

	return PackedIndex(n) + 1
}

// SlowPrime returns whether a number is prime or not, using a bute force search
func SlowPrime(n int) bool {
	root := int(math.Sqrt(float64(n)))

	if root > PackedPrimes[PackedPrimesEnd] {
		err := fmt.Errorf("SlowPrime(%d) exceeded max prime. Did you call Init()?", n)
		panic(err)
	}

	// Check each potential divisor to see if number divides evenly (i.e., is not prime).
	for i := 0; PackedPrimes[i] <= root; i++ {
		if n%PackedPrimes[i] == 0 {
			return false
		}
	}

	return true
}

// Prime returns true if number is prime
func Prime(number int) bool {
	if number > PackedPrimes[PackedPrimesEnd] {
		return SlowPrime(number)
	}
	return number == PackedPrimes[PackedIndex(number)]
}

// packPrimes fills PackedPrimes with prime numbers
func packPrimes() {
	for i := 0; i < len(Primes); i++ {
		if Primes[i] {
			PackedPrimes = append(PackedPrimes, i)
		}
	}
	PackedPrimesEnd = len(PackedPrimes) - 1
}

// PackedIndex returns the index in PackedPrimes of n
func PackedIndex(n int) int {
	upper := PackedPrimesEnd
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
			if mid == 0 {
				return mid
			}
			upper = mid - 1
		}

	}
	return upper
}

// excludes generates the numbers to exclude from the seive (the non-primes)
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

// seive implements the Seive of Eranthoses to find prime numbers
func seive() {
	upper := MaxPrime
	fmt.Println("upper: ", upper)
	for i := 0; i <= upper; i++ {
		Primes = append(Primes, true)
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

// factors returns a list of the prime factors of n
func factors(n int) []int {
	f := make([]int, 0)

	root := int(math.Sqrt(float64(n)))
	for i := 0; PackedPrimes[i] <= root; i++ {
		if n%PackedPrimes[i] == 0 {
			f = append(f, PackedPrimes[i])
			// Since we are iterating only up to root (as opposed to n/2)
			// we need to also add the 'reciprocal' factors. For instance,
			// when n=10 we iterate up to 3, which would miss 5 as a factor.
			d := n / PackedPrimes[i]
			if d != PackedPrimes[i] && Prime(d) {
				f = append(f, d)
			}
		}
	}

	return f
}

// seiveLowMemory Implements the seive of Eranthoses using an array of counters
//       1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29
// -2:   1 2 3   5   7   9    11    13    15    17    19    21    23    25    27    29
// -3:   1 2 3   5   7        11    13          17    19          23    25          29
// -5:   1 2 3       7        11    13          17    19          23                29
func seiveLowMemory(product int) {
	f := factors(product)
	counters := make([]int, len(f))

	for i := 1; i < product; i++ {
		keep := true
		// Increment each counter one tick.
		for c := 0; c < len(counters); c++ {
			counters[c]++
			if counters[c] == f[c] {
				// If any counter is zero, delete this number.
				counters[c] = 0
				keep = false
			}
		}
		if keep {
			// TODO: Keep this number; it is prime.
		}
	}
}

// Save writes PackedPrimes to a file
func Save() {
	file, err := os.Create("primes.gob")
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		panic(err)
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	encoder.Encode(PackedPrimesEnd)
	encoder.Encode(PackedPrimes)
}

// Load reads the contents of a file into PackedPrimes
func Load(fName string) {
	file, err := os.Open(fName)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		panic(err)
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&PackedPrimesEnd)
	if err != nil {
		fmt.Printf("error reading PackedPrimesEnd: %v", err)
		panic(err)
	}
	err = decoder.Decode(&PackedPrimes)
	if err != nil {
		fmt.Printf("error reading packedPrimes: %v", err)
		panic(err)
	}
}

// Init initializes the primes package
func Init(save bool) {
	seive()
	packPrimes()
	if save {
		Save()
	}
}
