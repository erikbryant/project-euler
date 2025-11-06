package main

// go fmt ./... && go vet ./... && go test ./... && go build 518.go && time ./518
// go fmt ./... && go vet ./... && go test ./... && go build 518.go && ./518 && echo top | go tool pprof cpu.prof

import (
	"fmt"
	"os"
	"runtime/pprof"

	"github.com/erikbryant/util-golang/algebra"
	"github.com/erikbryant/util-golang/primey"
)

// Let S(n) = Σ a + b + c over all triples (a, b, c) such that:
//
// a, b and c are prime numbers.
// a < b < c < n.
// a+1, b+1, and c+1 form a geometric sequence.
// For example, S(100) = 1035 with the following triples:
//
// (2, 5, 11), (2, 11, 47), (5, 11, 23), (5, 17, 53), (7, 11, 17),
// (7, 23, 71), (11, 23, 47), (17, 23, 31), (17, 41, 97), (31, 47, 71),
// (71, 83, 97)
//
// Find S(10^8).

// numerators given a and denom, finds all (a*k, b*k, c*k) triplets where k=num/denom and num=(denom+1)..n
func numerators(a, denom, n int) int {
	sum := 0

	for num := denom + 1; ; num++ {
		b := a * num / denom
		c := b * num / denom
		if c > n {
			break
		}
		if algebra.GCD(num, denom) != 1 {
			continue
		}
		if !primey.Prime(b - 1) {
			continue
		}
		if !primey.Prime(c - 1) {
			continue
		}
		sum += a + b + c - 3
	}

	return sum
}

// generator for each minimal (a,b,c) generates all (a, b*k, c*k) expansions
func generator(n int) int {
	sum := 0

	// For each a, generate (a,b,c) where b and c are minimal
	// and generate all (a, b*k, c*k) expansions, including
	// fractional values of k
	for _, p := range primey.Iter() {
		a := p + 1

		if a > (n - 3) {
			break
		}

		// Fractional multiples
		for denom := 2; ; denom++ {
			denomSquare := denom * denom
			if denomSquare > a {
				break
			}
			if a%denomSquare == 0 {
				sum += numerators(a, denom, n)
			}
		}

		// Integral multiples
		for k := 2; ; k++ {
			b := a * k
			c := b * k
			if c > n {
				break
			}
			if !primey.Prime(b - 1) {
				continue
			}
			if !primey.Prime(c - 1) {
				continue
			}
			sum += a + b + c - 3
		}
	}

	return sum
}

func S(n int) int {
	return generator(n)
}

func main() {
	fmt.Printf("Welcome to 518\n\n")

	fileHandle, _ := os.Create("cpu.prof")
	_ = pprof.StartCPUProfile(fileHandle)
	defer pprof.StopCPUProfile()

	upper := 1000 * 1000 * 100
	sum := S(upper)
	fmt.Printf("\nS(%d): %d\n\n", upper, sum)
}
