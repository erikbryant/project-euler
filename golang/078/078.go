package main

// go fmt ./... && go vet ./... && go test && go build 078.go && time ./078

import (
	"fmt"

	"github.com/erikbryant/util-golang/figurate"
)

// Let p(n) represent the number of different ways in which n coins can be
// separated into piles. For example, five coins can be separated into piles
// in exactly seven different ways, so p(5)=7.
//
// OOOOO
// OOOO   O
// OOO   OO
// OOO   O   O
// OO   OO   O
// OO   O   O   O
// O   O   O   O   O
//
// Find the least value of n for which p(n) is divisible by one million.

// Subtract returns (t1-t2)%m
func modSub(t1, t2, m int) int {
	t2 %= m
	t1 += m - t2
	return t1 % m
}

// P returns p[k] or zero if k < 0 (safe indexing)
func P(p []int, n int) int {
	if n < 0 {
		return 0
	}
	return p[n]
}

// partitionCountFind returns the first n for which the partition count is divisible by m
func partitionCountFind(m int) int {
	// 𝑝(𝑛)=𝑝(𝑛−1)+𝑝(𝑛−2)−𝑝(𝑛−5)−𝑝(𝑛−7)+𝑝(𝑛−12)+𝑝(𝑛−15)+…
	// where the integers are pentagonal numbers

	p := []int{}
	p = append(p, 1)

	for n := 1; ; n++ {
		p = append(p, 0)
		for k := 1; k < n+1; {
			p[n] += P(p, n-figurate.Pentagonal(k))
			p[n] += P(p, n-figurate.Pentagonal(-k))
			p[n] %= m
			k++
			p[n] = modSub(p[n], P(p, n-figurate.Pentagonal(k)), m)
			p[n] = modSub(p[n], P(p, n-figurate.Pentagonal(-k)), m)
			k++
		}
		if p[n]%m == 0 {
			return n
		}
	}
}

func main() {
	fmt.Printf("Welcome to 078\n\n")

	n := partitionCountFind(1000 * 1000)
	fmt.Printf("\np(%d) is divisible by 1,000,000\n", n)
}
