package main

// go fmt ./... && go vet ./... && go test ./... && go build 187.go && time ./187

import (
	"fmt"
	"math"

	"github.com/erikbryant/util-golang/primey"
)

// A composite is a number containing at least two prime factors. For example,
// 15 = 3 × 5; 9 = 3 × 3; 12 = 2 × 2 × 3.
//
// There are ten composites below thirty containing precisely two, not
// necessarily distinct, prime factors: 4, 6, 9, 10, 14, 15, 21, 22, 25, 26.
//
// How many composite integers, n < 10^8, have precisely two, not necessarily
// distinct, prime factors?
//
// A composite number with precisely two prime factors is called "semiprime".
// https://en.wikipedia.org/wiki/Semiprime

func f(n, k int) int {
	return primey.Pi(n/int(primey.Nth(k-1))) - k + 1
}

// semiprimes returns the number of semiprimes less than or equal to n
// https://en.wikipedia.org/wiki/Semiprime
func semiprimes(n int) int {
	root := int(math.Sqrt(float64(n)))
	maxFound := primey.Pi(root)
	count := 0

	for k := 1; k <= maxFound; k++ {
		count += f(n, k)
	}

	return count
}

func main() {
	fmt.Printf("Welcome to 187\n\n")

	maxFound := 100 * 1000 * 1000
	count := semiprimes(maxFound - 1) // We need the count _less than_ maxFound

	fmt.Printf("Number of 2-composite integers < %d = %d\n\n", maxFound, count)
}
