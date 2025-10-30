package main

// go fmt ./... && go vet ./... && go test ./... && go build 069.go && time ./069

import (
	"fmt"

	"github.com/erikbryant/util-golang/algebra"
	"github.com/erikbryant/util-golang/primey"
)

// Euler's totient function, ɸ(n) [sometimes called the phi function], is defined as the number of positive integers not exceeding n which are relatively prime to n. For example, as 1, 2, 4, 5, 7, and 8, are all less than or equal to nine and relatively prime to nine, ɸ(9)=6.
//
// _n_   _Relatively Prime_   _ɸ(n)_     _n/ɸ(n)_
//  2           1                1           2
//  3          1,2               2         1.5
//  4          1,3               2           2
//  5        1,2,3,4             4        1.25
//  6          1,5               2           3
//  7      1,2,3,4,5,6           6      1.1666...
//  8        1,3,5,7             4           2
//  9      1,2,4,5,7,8           6         1.5
// 10        1,3,7,9             4         2.5
//
// It can be seen that n = 6 produces a maximum n/ɸ(n) for n<=q 10.
// Find the value of n<=q 1,000,000 for which n/ɸ(n) is a maximum.

// Our goal is to find the highest ratio of n/phi.
// This means that we are looking to minimize phi.
// Numbers with a minimum phi are ones that have
// a maximum number of factors. Find each of the
// numbers that have a maximum set of factors
// (that is, they are factorials of the first
// prime numbers).

func main() {
	fmt.Printf("Welcome to 069\n\n")

	n := 1
	fmt.Printf("  n        n/ɸ(n)\n")
	fmt.Printf("------    --------\n")
	for _, p := range primey.Iter() {
		if n*p > 1000*1000 {
			break
		}
		n *= p
		phi := algebra.Totient(n)
		ratio := float64(n) / float64(phi)
		fmt.Printf("%6d  %10f\n", n, ratio)
	}
	fmt.Println()
}
