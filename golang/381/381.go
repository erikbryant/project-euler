package main

// go fmt ./... && go vet ./... && go test ./... && go build 381.go && time ./381

import (
	"fmt"
	"strconv"

	"github.com/erikbryant/util-golang/primey"
)

// For a prime p let S(p) = (Σ (p-k)!) mod (p) for 1 <= k <= 5.
//
// For example, if p=7,
// (7-1)! + (7-2)! + (7-3)! + (7-4)! + (7-5)! = 6! + 5! + 4! + 3! + 2! = 720 + 120 + 24 + 6 + 2 = 872.
// As 872 mod (7) = 4, S(7) = 4.
//
// It can be verified that Σ S(p) = 480 for 5 <= p < 100.
//
// Find Σ S(p) for 5 <= p < 10^8.

// S returns the result of: For a prime p let S(p) = (∑(p-k)!) mod(p) for 1 ≤ k ≤ 5.
func S(p int) int {
	sum := 0

	// (p-1)!%p == p-1
	sum += p - 1

	// (p-2)!%p == 1
	sum++

	// (p-3)!%p == (p-1)/2
	sum += (p - 1) / 2

	// (p-4)!%p ==
	four := (p + 1) / 6
	if (p+1)%6 != 0 {
		four = p - four
	}
	sum += four

	// (p-5)!%p ==
	var five int
	switch p % 24 {
	case 1:
		five = (p - 1) / 24
	case 5:
		five = p/4 - p/24
	case 7:
		five = p/3 - p/24
	case 11:
		five = p/2 - p/24
	case 13:
		five = (p+1)/2 + (p-1)/24
	case 17:
		five = p - p/3 + p/24
	case 19:
		five = p - p/4 + p/24
	case 23:
		five = (p - 1) - (p-1)/24
	default:
		panic("Yikes! " + strconv.Itoa(p%24))
	}
	sum += five

	return sum % p
}

func sumS(min, max int) int {
	sum := 0

	for i := min; i < max; i++ {
		if primey.Prime(i) {
			sum += S(i)
		}
	}

	return sum
}

// Find ∑S(p) for 5 ≤ p < 10^8.
func main() {
	fmt.Printf("Welcome to 381\n\n")

	sum := sumS(5, 1000*1000*100)
	fmt.Printf("Σ S(p) for 5 <= p < 10^8 = %d\n\n", sum)
}
