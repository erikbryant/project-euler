package main

// go fmt && golint && go test && go run 087.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

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
	cpuprofile   = flag.String("cpuprofile", "", "write cpu profile to file")
	primePowers2 = []int{}
	primePowers3 = []int{}
	primePowers4 = []int{}
)

// The smallest number expressible as the sum of a prime square, prime cube,
// and prime fourth power is 28. In fact, there are exactly four numbers below
// fifty that can be expressed in such a way:
//
// 28 = 2^2 + 2^3 + 2^4
// 33 = 3^2 + 2^3 + 2^4
// 47 = 2^2 + 3^3 + 2^4
// 49 = 5^2 + 2^3 + 2^4
//
// How many numbers below fifty million can be expressed as the sum of a prime
// square, prime cube, and prime fourth power?

func init() {
	primes.Load("../primes.gob")
	generatePowers()
}

// generatePowers populates all of the powers of 2, 3, and 4.
func generatePowers() {
	max := 50 * 1000 * 1000

	for _, p := range primes.PackedPrimes {
		v := int(math.Pow(float64(p), 2.0))
		if v >= max {
			break
		}
		primePowers2 = append(primePowers2, v)
	}

	for _, p := range primes.PackedPrimes {
		v := int(math.Pow(float64(p), 3.0))
		if v >= max {
			break
		}
		primePowers3 = append(primePowers3, v)
	}

	for _, p := range primes.PackedPrimes {
		v := int(math.Pow(float64(p), 4.0))
		if v >= max {
			break
		}
		primePowers4 = append(primePowers4, v)
	}
}

// generatePrimePowerSums finds all PPS's that are sums of powers of 2,3, and 4.
func generatePrimePowerSums(max int) int {
	// Put the results in a map to remove duplicates.
	results := make(map[int]bool)

	for _, i := range primePowers4 {
		for _, j := range primePowers3 {
			for _, k := range primePowers2 {
				val := i + j + k
				if val >= max {
					break
				}
				results[val] = true
			}
		}
	}

	return len(results)
}

func main() {
	fmt.Printf("Welcome to 087\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	fmt.Println("Count:", generatePrimePowerSums(50*1000*1000))
}
