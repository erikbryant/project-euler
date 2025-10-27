package main

// go fmt ./... && go vet ./... && go test ./... && go build 160.go && time ./160

import (
	"fmt"

	"github.com/erikbryant/util-golang/factorials"
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

	upper := 1000 * 1000 * 1000 * 1000
	p := message.NewPrinter(language.English)
	algorithm := "dnc"

	fmt.Printf("=====  %s  =====\n\n", algorithm)
	for i := 10; i <= upper; i *= 10 {
		iSmall := factorials.Reduce(i)
		f := factorials.Factorial(iSmall, algorithm)
		p.Printf("%18d! = %18d! = %12d\n", i, iSmall, f)
	}

	fmt.Println()
}
