package main

// go fmt ./... && go vet ./... && go test && go run 187.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"runtime/pprof"

	"github.com/erikbryant/util-golang/primes"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
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
//
//

func f(n, k int) int {
	return primes.Pi(n/primes.Primes[k-1]) - k + 1
}

// semiprimes returns the number of semiprimes less than or equal to n
// https://en.wikipedia.org/wiki/Semiprime
func semiprimes(n int) int {
	root := int(math.Sqrt(float64(n)))
	maxFound := primes.Pi(root)
	count := 0

	for k := 1; k <= maxFound; k++ {
		count += f(n, k)
	}

	return count
}

func main() {
	fmt.Printf("Welcome to 187\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	maxFound := 100 * 1000 * 1000
	count := semiprimes(maxFound - 1) // We need the count _less than_ maxFound
	fmt.Println("There are ", count, "2-composite integers <", maxFound)
}
