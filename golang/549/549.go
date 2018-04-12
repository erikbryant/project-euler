package main

import (
	"../primes"
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"runtime/pprof"
)

var (
	cpuprofile     = flag.String("cpuprofile", "", "write cpu profile to file")
	factorCacheLen = 0
	factorCache    = []map[int]uint8{}
)

func init() {
	primes.Load("../primes.gob")
	Load()
	// Save()
}

func Save() {
	if factorCacheLen != 0 {
		fmt.Println("ERROR: The cache already has data in it!")
		return
	}

	fmt.Println("Generating factors ...")

	// 0 has no factors > 1
	factorCache = append(factorCache, nil)
	factorCacheLen++

	// 1 has no factors > 1
	factorCache = append(factorCache, nil)
	factorCacheLen++

	for ; factorCacheLen <= 1000*1000*10; factorCacheLen++ {
		factorCache = append(factorCache, factorSlow(factorCacheLen))
	}

	fmt.Println("Saving factors ...")

	file, err := os.Create("549.gob")
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		panic(err)
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	encoder.Encode(factorCacheLen)
	encoder.Encode(factorCache)

	fmt.Println("... done saving")
}

func Load() {
	file, err := os.Open("549.gob")
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		panic(err)
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&factorCacheLen)
	if err != nil {
		fmt.Printf("error reading factorCacheLen: %v", err)
		panic(err)
	}
	err = decoder.Decode(&factorCache)
	if err != nil {
		fmt.Printf("error reading factorCache: %v", err)
		panic(err)
	}
	fmt.Println("Loaded", factorCacheLen, "factors into the cache")
}

func factorSlow(n int) map[int]uint8 {
	factors := make(map[int]uint8)

	// Find all of the 2 factors, since they are quick
	for (n & 0x01) == 0 {
		factors[2]++
		n = n >> 1
		if n == 1 {
			return factors
		}
	}

	root := int(math.Sqrt(float64(n)))
	for i := 1; primes.PackedPrimes[i] <= root; i++ {
		p := primes.PackedPrimes[i]
		for n%p == 0 {
			factors[p]++
			n = n / p
			if n == 1 {
				return factors
			}
		}
	}

	// We did not find any factors for 'n',
	// so it must be prime.
	factors[n]++
	return factors
}

func factor(n int) map[int]uint8 {
	if n < factorCacheLen {
		return factorCache[n]
	}
	return factorSlow(n)
}

// TODO: Explore this option...
// Create a list of prime factors <= 10^8.
// For each prime factor, create a list of which numbers provide them.
// Since 2^27 > 10^8 we at most need to keep 27.
//
// 2 - 2|1, 4|3, 6|4, 8|7, ..., 2*27|xx
// 3 - 3|1, 6|2, 9|4, 12|5, ..., 3*27|yy
// 5 - 5|1, 10|2, 15|3, ..., 5*27|zz
// ...
// 17 - 17|1, ...

func countFactors(n int, f int) uint8 {
	count := uint8(0)
	for {
		if n%f != 0 {
			break
		}
		n = n / f
		count++
	}
	return count
}

func minProvider(fact int, count uint8) int {
	if count == 1 {
		return fact
	}

	sum := 0
	c := uint8(0)
	for c < count {
		sum += fact
		c += countFactors(sum, fact)
	}

	return sum
}

// Let s(n) be the smallest number m such that n divides m!
func s(n int) int {
	m := 0

	for fact, count := range factor(n) {
		f := minProvider(fact, count)
		if f > m {
			m = f
		}
	}

	return m
}

// Let S(n) be ∑s(i) for 2 ≤ i ≤ n
func sumS(n int) (sum int) {
	var start int = 2

	// Skip over 2 -> 2!
	start = 3
	sum += 2

	// Skip over 3 -> 3!
	start = 4
	sum += 3

	// Skip over 4 -> 4!
	start = 5
	sum += 4

	// Because the loop is unrolled, it must begin
	// on an odd number.
	if start&0x01 != 1 {
		fmt.Println("ERROR: loop is not synchronized")
		return 0
	}

	for a := start; a <= n; a++ {
		// Odd
		if primes.Prime(a) {
			sum += a
		} else {
			sum += s(a)
		}

		// Even
		a++
		if primes.Prime(a >> 1) {
			sum += a >> 1
		} else {
			sum += s(a)
		}
	}

	return
}

// Find S(10^8)
func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	fmt.Println("Welcome to 549")
	n := 1000 * 1000 * 100
	fmt.Println("For n:", n, "answer:", sumS(n))
}
