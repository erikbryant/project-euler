package main

// go fmt ./... && go vet ./... && go test ./... && go build 087.go && time ./087

import (
	"fmt"
	"math"

	"github.com/erikbryant/util-golang/primey"
)

var (
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
	generatePowers()
}

// generatePowers populates all powers of 2, 3, and 4.
func generatePowers() {
	maxFound := 50 * 1000 * 1000

	for _, p := range primey.Iter() {
		v := int(math.Pow(float64(p), 2.0))
		if v >= maxFound {
			break
		}
		primePowers2 = append(primePowers2, v)
	}

	for _, p := range primey.Iter() {
		v := int(math.Pow(float64(p), 3.0))
		if v >= maxFound {
			break
		}
		primePowers3 = append(primePowers3, v)
	}

	for _, p := range primey.Iter() {
		v := int(math.Pow(float64(p), 4.0))
		if v >= maxFound {
			break
		}
		primePowers4 = append(primePowers4, v)
	}
}

// generatePrimePowerSums finds all power prime sums that are sums of powers of 2,3, and 4.
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

	fmt.Printf("Number of n where n < 50,000,000 and n = p1^2 + p2^3 + p3^4 : %d\n\n", generatePrimePowerSums(50*1000*1000))
}
