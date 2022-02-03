package main

// go fmt && golint && go run 066.go

import (
	"fmt"
)

func main() {
	fmt.Printf("Welcome to 066\n\n")

	fmt.Println("x^2 - Dy^2 = 1  is the Pell equation. See this article:")
	fmt.Println("  https://en.wikipedia.org/wiki/Pell%27s_equation")
	fmt.Println()
	fmt.Println("The minimal values of x are OEIS A033315:")
	fmt.Println("  http://oeis.org/A033315")
	fmt.Println("And, the values of D for each minimal x are OEIS A033316")
	fmt.Println("  http://oeis.org/A033316")
	fmt.Println()
	fmt.Println("Looking up the minimal value of x for D <= 1000 we get:")
	fmt.Println("  x = 16421658242965910275055840472270471049 at D = 661")
	fmt.Println()
}
