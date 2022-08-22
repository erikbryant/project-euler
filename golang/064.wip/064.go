package main

// go fmt ./... && go vet ./... && go test && go run 064.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

import (
	"github.com/erikbryant/project-euler/golang/util"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"runtime/pprof"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
)

// https://projecteuler.net/problem=64

// https://math.stackexchange.com/questions/265690/continued-fraction-of-a-square-root
//
// Scroll down to the recursive formula...
// https://en.wikipedia.org/wiki/Continued_fraction#Basic_formula

func period(n int) int {
	if util.IsSquare(n) {
		return 0
	}

	d := math.Sqrt(float64(n))
	a := math.Floor(d)
	aZero := a
	sequence := []int{}

	// Series starts at second term, terminates at 2x first term.
	// https://www.johndcook.com/blog/2020/08/04/continued-fraction-sqrt/
	for a != 2*aZero {
		d = 1.0 / (d - a)
		a = math.Floor(d)
		sequence = append(sequence, int(a))
	}

	// The sequence will be a palindrome followed by the final term
	if !util.IsPalindromeInt(sequence[:len(sequence)-1]) {
		fmt.Println("Palindrome fail!", n)
	}

	return len(sequence)
}

func main() {
	fmt.Printf("Welcome to 064\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	odd := 0
	for N := 0; N <= 1000; N++ {
		p := period(N)
		if p&0x01 == 1 {
			odd++
		}
	}

	fmt.Println("Odd periods:", odd)
}
