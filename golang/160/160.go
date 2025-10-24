package main

// go fmt ./... && go vet ./... && go test && go build 160.go && time ./160
// go fmt ./... && go vet ./... && go test && go build 160.go && ./160 && echo top | go tool pprof cpu.prof

import (
	"160/bins"
	"160/dnc"
	"160/naive"
	"fmt"
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

func main() {
	fmt.Printf("Welcome to 160\n\n")

	fileHandle, _ := os.Create("cpu.prof")
	_ = pprof.StartCPUProfile(fileHandle)
	defer pprof.StopCPUProfile()

	p := message.NewPrinter(language.English)

	upper := 1000 * 1000 * 1000

	// Naive
	for i := 10; i <= upper; i *= 10 {
		f := naive.Factorial(i)
		p.Printf("%18d! = %12d\n", i, f)
	}

	// DNC
	for i := 10; i <= upper; i *= 10 {
		f := dnc.Factorial(i)
		p.Printf("%18d! = %12d\n", i, f)
	}

	// Bins
	for i := 10; i <= upper; i *= 10 {
		f := bins.Factorial(i)
		p.Printf("%18d! = %12d\n", i, f)
	}
}
