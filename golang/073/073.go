package main

// go fmt ./... && go vet ./... && go test ./... && go build 073.go && time ./073

import (
	"fmt"

	"github.com/erikbryant/util-golang/primey"
)

// Consider the fraction, n/d, where n and d are positive integers. If n < d and HCF(n, d)=1,
// it is called a reduced proper fraction.
//
// If we list the set of reduced proper fractions for d <= 8 in ascending order of size, we get:
// 1/8, 1/7, 1/6, 1/5, 1/4, 2/7, 1/3, _3/8_, _2/5_, _3/7_, 1/2, 4/7, 3/5, 5/8, 2/3, 5/7,
// 3/4, 4/5, 5/6, 6/7, 7/8
//
// It can be seen that there are 3 fractions between 1/3 and 1/2.
//
// How many fractions lie between 1/3 and 1/2 in the sorted set of reduced proper fractions for
// d <= 12,000?

type fraction struct {
	n int
	d int
}

// Reduce a fraction n/d such that the HCF(n, d) == 1.
func reduce(n, d int) (rN, rD int) {
	for _, p := range primey.Iter() {
		if p > n {
			break
		}
		for n%p == 0 && d%p == 0 {
			n = n / p
			d = d / p
		}
	}

	return n, d
}

// For 2 <= d <= max, find all fractions such that: 1/3 < n/d < 1/2
func rpf(max int) int {
	if max <= 4 {
		return 0
	}

	foundFractions := make(map[fraction]int)

	for d := 5; d <= max; d++ {
		start := d/3 + 1
		end := d / 2
		if d%2 == 0 {
			end--
		}

		for n := start; n <= end; n++ {
			rN, rD := reduce(n, d)
			foundFractions[fraction{n: rN, d: rD}] = 1
		}
	}

	return len(foundFractions)
}

func main() {
	fmt.Printf("Welcome to 073\n\n")

	fmt.Printf("# fractions between 1/3 and 1/2 for d <= 12,000: %d\n\n", rpf(12000))
}
