package main

// go fmt ./... && go vet ./... && go test && go build 204.go && time ./204

import (
	"fmt"

	"github.com/erikbryant/util-golang/algebra"
)

// A Hamming number is a positive number which has no prime factor larger than 5.
// So the first few Hamming numbers are 1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 15.
// There are 1105 Hamming numbers not exceeding 10^8.
//
// We will call a positive number a generalised Hamming number of type n, if it
// has no prime factor larger than n.
// Hence, the Hamming numbers are the generalised Hamming numbers of type 5.
//
// How many generalised Hamming numbers of type 100 are there which don't exceed 10^9?

func main() {
	fmt.Printf("Welcome to 204\n\n")

	k := algebra.KSmooths(1000*1000*1000, 100)
	fmt.Println(len(k))
}
