package main

// go fmt ./... && go vet ./... && go test ./... && go build 549.go && time ./549
// go fmt ./... && go vet ./... && go test ./... && go build 549.go && ./549 && echo top | go tool pprof cpu.prof

import (
	"fmt"
	"log"
	"math"
	"os"
	"runtime/pprof"

	"github.com/erikbryant/util-golang/primey"
)

// The smallest number m such that 10 divides m! is m=5.
// The smallest number m such that 25 divides m! is m=10.
//
// Let s(n) be the smallest number m such that n divides m!.
// So s(10)=5 and s(25)=10.
// Let S(n) be Σ s(i) for 2 <= i <= n.
// S(100)=2012.
//
// Find S(10^8).

// Not my best work, but it does run in under a minute on my Mac notebook.
// There is at least one more mathy simplification to be made, but what
// that is eludes me.

var (
	// factors stores the running total of how many p's (k*p)! provides
	factors = map[int][]int{}

	// primes is a cache of the primes we need, as primey.Iter() is too slow for how many times we would call it
	primes = []int{}
)

// divides returns the number of times f divides into n
func divides(n, f int) int {
	count := 0
	for n%f == 0 {
		count++
		n /= f
	}
	return count
}

// makeFactors fills primes with primes 0..upper/2 and fills factors
func makeFactors(upper int) {
	for _, p := range primey.Iter() {
		primes = append(primes, p)
		if p > upper/2 {
			break
		}
	}

	// Find the running total of how many p's are provided by multiples of (k*p)!
	for _, p := range primes {
		if p > upper/2 {
			break
		}
		total := 1
		multiple := p
		counts := []int{0, total}
		for multiple < upper {
			multiple += p
			total += divides(multiple, p)
			counts = append(counts, total)
		}
		factors[p] = counts
	}
}

// findMultiple returns the multiple of factor that provides count worth of factors in factor!
func findMultiple(factor, count int) int {
	for i := 1; ; i++ {
		if factors[factor][i] >= count {
			return factor * i
		}
	}
}

// s returns the smallest number m such that n divides m!
func s(n int) int {
	// n decomposes into:  p1^j1 * p2^j2 * .. * pn^jn
	// for each pn, find the lowest k for which
	// (k*pn)! provides jn pn's. Once we compute jn
	// k*pn is a simple lookup using findMultiple(pn, jn).

	m := 1
	root := int(math.Sqrt(float64(n)))

	for _, p := range primes {
		if p > root {
			break
		}
		if n%p == 0 {
			count := 1
			n = n / p
			for n%p == 0 {
				count++
				n = n / p
			}
			pMultiple := findMultiple(p, count)
			m = max(m, pMultiple)
			if n == 1 {
				return m
			}
		}
	}

	// We did not find any factors for 'n',
	// so it must be prime.
	return n
}

// S returns ∑s(i) for 2 ≤ i ≤ n
func S(n int) int {
	sum := 0
	for i := 2; i <= n; i++ {
		sum += s(i)
	}
	return sum
}

// Find S(10^8)
func main() {
	fmt.Println("Welcome to 549")

	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	_ = pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	n := 1000 * 1000 * 100
	makeFactors(n)
	fmt.Printf("\nS(%d) = %d\n\n", n, S(n))
}
