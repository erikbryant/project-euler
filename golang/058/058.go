package main

// go fmt ./... && go vet ./... && go test ./... && go build 058.go && time ./058

import (
	"fmt"

	"github.com/erikbryant/util-golang/primey"
)

// Starting with 1 and spiralling anticlockwise in the following way, a square spiral with side length 7 is formed.
//
// _37_  36   35  34  33   32  _31_
//  38  _17_  16  15  14  _13_  30
//  39   18   _5_  4  _3_  12   29
//  40   19    6   1   2   11   28
//  41   20   _7_  8   9   10   27
//  42   21   22  23  24   25   26
// _43_  44   45  46  47   48   49
//
// It is interesting to note that the odd squares lie along the bottom right diagonal, but what is more interesting
// is that 8 out of the 13 numbers lying along both diagonals are prime; that is, a ratio of 8/13 ≅ 62%.
// If one complete new layer is wrapped around the spiral above, a square spiral with side length 9 will be formed.
// If this process is continued, what is the side length of the square spiral for which the ratio of primes along
// both diagonals first falls below 10%?

func diagonals(sideLength int) [4]int {
	d := [4]int{0, 0, 0, 0}

	if sideLength < 3 {
		return d
	}

	// Bottom right
	d[3] = sideLength * sideLength

	// Bottom left
	d[2] = d[3] - sideLength + 1

	// Top left
	d[1] = d[2] - sideLength + 1

	// Top right
	d[0] = d[1] - sideLength + 1

	return d
}

func main() {
	fmt.Println("Welcome to 058")

	numbers := 1
	sideLength := 3
	prime := 0

	for {
		d := diagonals(sideLength)
		numbers += len(d)
		// d[3] is never prime; don't check it.
		for i := 0; i <= 2; i++ {
			if primey.Prime(d[i]) {
				prime++
			}
		}
		primePct := float64(prime) / float64(numbers) * 100.0
		fmt.Printf("Side length: %d  Numbers: %d  Primes: %d (%2f%%)\n", sideLength, numbers, prime, primePct)
		if primePct < 10.0 {
			break
		}
		sideLength += 2
	}

	fmt.Printf("\nSide length at which ratio falls below 10%% = %d\n\n", sideLength)
}
