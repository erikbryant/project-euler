package main

import (
	"fmt"
)

// Consider the sequence 1504170715041707n mod 4503599627370517.
//
// An element of this sequence is defined to be an Eulercoin if it is
// strictly smaller than all previously found Eulercoins.
//
// For example, the first term is 1504170715041707 which is the first Eulercoin.
// The second term is 3008341430083414 which is greater than 1504170715041707
// so is not an Eulercoin. However, the third term is 8912517754604 which is
// small enough to be a new Eulercoin.
//
// The sum of the first 2 Eulercoins is therefore 1513083232796311.
//
// Find the sum of all Eulercoins.
//
func eulerCoin() int64 {
	var left int64
	var right int64
	var sum int64

	// We can slim 'right' down quite a bit. It does not have to be so many
	// multiples of 'left'. That will save a lot of looping time.
	left = 1504170715041707
	right = 4503599627370517 % left

	for ; left > 0 && right > 0; left -= right {
		fmt.Println(left)
		sum += left
		// Keep slimming
		right = right % left
	}

	return sum
}

func main() {
	fmt.Printf("Welcome to 700\n\n")

	fmt.Println("Sum:", eulerCoin())
}
