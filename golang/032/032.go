package main

import (
	"fmt"
)

var (
	productsFound map[int]bool
)

func init() {
	productsFound = make(map[int]bool)
}

// Generating permutation using Heap Algorithm
// https://www.geeksforgeeks.org/heaps-algorithm-for-generating-permutations/
func heapPermutation(digits []int, size int, c chan []int) {
	if size == 1 {
		var temp []int
		for i := 0; i < len(digits); i++ {
			temp = append(temp, digits[i])
		}
		c <- temp
		return
	}

	for i := 0; i < size; i++ {
		heapPermutation(digits, size-1, c)

		// if size is odd, swap first and last element
		// If size is even, swap ith and last element
		swap := 0
		if size%2 == 0 {
			swap = i
		}
		digits[swap], digits[size-1] = digits[size-1], digits[swap]
	}
}

func makeDigits(n int, c chan []int) {
	defer close(c)

	var digits []int
	for i := 1; i <= n; i++ {
		digits = append(digits, i)
	}

	heapPermutation(digits, len(digits), c)
}

func mulDigits(digits []int, m2, product int) bool {
	M1 := 0
	M2 := 0
	Product := 0

	for i := 0; i < m2; i++ {
		M1 *= 10
		M1 += digits[i]
	}
	for i := m2; i < product; i++ {
		M2 *= 10
		M2 += digits[i]
	}
	for i := product; i < len(digits); i++ {
		Product *= 10
		Product += digits[i]
	}

	if M1*M2 == Product {
		productsFound[Product] = true
		return true
	}

	return false
}

func pandigitalProduct(digits []int) {
	for m2 := 1; m2 < len(digits)-1; m2++ {
		for product := m2 + 1; product < len(digits); product++ {
			if mulDigits(digits, m2, product) {
				fmt.Println(digits, m2, product, digits[product:])
			}
		}
	}
}

func main() {
	c := make(chan []int, 1000)
	go makeDigits(9, c)
	for {
		next, ok := <-c
		if !ok {
			break
		}
		pandigitalProduct(next)
	}
	fmt.Println("Unique pandigital products:", len(productsFound))
	sum := 0
	for key := range productsFound {
		sum += key
	}
	fmt.Println("Sum of products:", sum)
}
