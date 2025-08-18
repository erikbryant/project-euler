package main

import (
	"fmt"
	"log"
	"strconv"
)

// go fmt ./... && go vet ./... && go test && go run 932.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

// Given positive integers a and b, the concatenation ab we call a 2025-number
// if ab = (a + b)^2. Other examples are 3025 and 81. Note 9801 is not a
// 2025-number because the concatenation of 98 and 1 is 981.
// Let T(n) be the sum of all 2025-numbers with n digits or less. You are given
// T(4) = 5131. Find T(16).

// summable returns true if square consists of digits that sum to root
func summable(root, square int) bool {
	s := fmt.Sprintf("%d", square)

	for i := 1; i < len(s); i++ {
		if s[i:][0] == '0' {
			continue
		}
		a, err := strconv.Atoi(s[:i])
		if err != nil {
			log.Fatalf("Error converting %v to int: %v", s, err)
		}
		b, err := strconv.Atoi(s[i:])
		if err != nil {
			log.Fatalf("Error converting %v to int: %v", s, err)
		}
		if a+b == root {
			return true
		}
	}

	return false
}

func main() {
	fmt.Printf("Welcome to 932\n\n")

	// largest 16-digit number
	upper := 9999999999999999
	root := 1
	total := 0

	for {
		square := root * root
		if square > upper {
			break
		}
		if summable(root, square) {
			fmt.Println(root, square)
			total += square
		}
		root++
	}

	fmt.Println("Total:", total)
}
