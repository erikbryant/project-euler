package main

// go fmt ./... && go vet ./... && go test ./... && go build 160.go && time ./160
// go fmt ./... && go vet ./... && go test ./... && go build 160.go && ./160 && echo top | go tool pprof cpu.prof

import (
	"160/bins"
	"160/dnc"
	"160/moessner"
	"160/naive"
	"160/swing"
	"fmt"
	"math"
	"os"
	"runtime/pprof"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// For any N, let f(N) be the last five digits before the trailing zeroes in N!.
// For example,
//
//  9! = 362880   so f(9)=36288
// 10! = 3628800 so f(10)=36288
// 20! = 2432902008176640000 so f(20)=17664
//
// Find f(1,000,000,000,000).

var (
	p = message.NewPrinter(language.English)
)

// reduce returns nSmall where nSmall <= n and nSmall! % dnc.Mod == n! % dnc.Mod
func reduce(n int) int {
	// Idea from: https://euler.stephan-brumme.com/160/
	nDigits := int(math.Log10(float64(n)))
	modDigits := int(math.Log10(float64(dnc.Mod)))

	exp := nDigits - modDigits
	for exp > 0 && n%5 == 0 {
		n /= 5
		exp--
	}

	return n
}

// factorial runs the given algorithm and displays the results for various values
func factorial(algorithm string) {
	var factorial func(int) int
	upper := 1000 * 1000 * 1000 * 1000

	switch algorithm {
	case "naive":
		factorial = naive.Factorial
	case "dnc":
		factorial = dnc.Factorial
	case "bins":
		factorial = bins.Factorial
	case "swing":
		// Very fast, but memory intensive (requires a list of primes up to n)
		upper = min(1000*1000*1000*10, upper)
		factorial = swing.Factorial
	case "moessner":
		// Nifty algorithm, but not very powerful
		upper = min(10, upper)
		factorial = moessner.Factorial
	default:
		fmt.Printf("Not a supported algorithm: %s\n", algorithm)
		os.Exit(1)
	}

	fmt.Printf("=====  %s  =====\n\n", algorithm)
	for i := 10; i <= upper; i *= 10 {
		iSmall := reduce(i)
		f := factorial(iSmall)
		p.Printf("%18d! = %18d! = %12d\n", i, iSmall, f)
	}
	fmt.Println()
}

func main() {
	fmt.Printf("Welcome to 160\n\n")

	fileHandle, _ := os.Create("cpu.prof")
	_ = pprof.StartCPUProfile(fileHandle)
	defer pprof.StopCPUProfile()

	factorial("dnc")
}
