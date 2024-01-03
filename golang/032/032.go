package main

import (
	"fmt"

	"github.com/erikbryant/util-golang/util"
)

var (
	productsFound map[int]bool
)

func init() {
	productsFound = make(map[int]bool)
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
	go util.MakeDigits(9, c)
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
