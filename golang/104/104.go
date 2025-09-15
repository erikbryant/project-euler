package main

// go fmt ./... && go vet ./... && go test && go build 104.go && time ./104

// The Fibonacci sequence is defined by the recurrence relation:
//   F(n) = F(n-1) + F(n-2) where F(1) = 1 and F(2) = 1
//
// It turns out that F(541), which contains 113 digits, is the first
// Fibonacci number for which the last nine digits are 1-9 pandigital
// (contain all the digits 1 to 9, but not necessarily in order). And
// F(2749), which contains 575 digits, is the first Fibonacci number
// for which the first nine digits are 1-9 pandigital.
//
// Given that F(k) is the first Fibonacci number for which the first
// nine digits AND the last nine digits are 1-9 pandigital, find k.

import (
	"fmt"
)

const (
	Digits = 1000 * 1000
)

var (
	Fibs = [3][Digits]byte{
		{1},
		{1},
		{0},
	}
	FibLens = [3]int{1, 1, 1}
	Indexer = [3]int{0, 1, 2}
)

// pandigital returns true if the slice contains digits 1-9
func pandigital(s []byte) bool {
	if len(s) != 9 {
		return false
	}
	return 0b1111111110 == 1<<s[0]|1<<s[1]|1<<s[2]|1<<s[3]|1<<s[4]|1<<s[5]|1<<s[6]|1<<s[7]|1<<s[8]
}

func pandigitalFirst() bool {
	i := Indexer[2]
	iLen := FibLens[i]
	if iLen < 9 {
		return false
	}
	return pandigital(Fibs[i][iLen-9 : iLen])
}

func pandigitalLast() bool {
	i := Indexer[2]
	iLen := FibLens[i]
	if iLen < 9 {
		return false
	}
	return pandigital(Fibs[i][0:9])
}

func fixLengths() {
	for i, _ := range Fibs {
		// Initializer / single-digit case
		if FibLens[i] <= 1 {
			FibLens[i] = 1
			continue
		}
		// Shrink length down to first non-zero digit
		for Fibs[i][FibLens[i]-1] == 0 {
			FibLens[i]--
		}
	}
}

func nextFib() {
	// Add: Fibs[Indexer[2]] = Fibs[Indexer[0]] + Fibs[Indexer[1]]

	// Make room for the sum
	iLen := max(FibLens[Indexer[0]], FibLens[Indexer[1]]) + 1 // Be loose here; fix later
	FibLens[Indexer[2]] = iLen

	// Compute sum
	carry := byte(0)
	for i := 0; i <= iLen; i++ {
		Fibs[Indexer[2]][i] = Fibs[Indexer[0]][i] + Fibs[Indexer[1]][i] + carry
		carry = Fibs[Indexer[2]][i] / 10
		Fibs[Indexer[2]][i] %= 10
	}

	fixLengths()

	// Shift
	Indexer[0], Indexer[1], Indexer[2] = Indexer[1], Indexer[2], Indexer[0]
}

func stringify() string {
	i := Indexer[2]
	iLen := FibLens[i]
	str := ""

	for _, val := range Fibs[i][0:iLen] {
		str = string('0'+val) + str
	}

	return str
}

func fibonacci() {
	for k := 0; ; k++ {
		first := pandigitalFirst()
		last := pandigitalLast()
		if first || last {
			fmt.Printf("At k=%8d, pandigitalFirst=%5t, pandigitalLast=%5t\n", k, first, last)
			if first && last {
				break
			}
		}
		nextFib()
	}
}

func main() {
	fmt.Printf("Welcome to 104\n\n")

	fibonacci()
}
