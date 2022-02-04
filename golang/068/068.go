package main

// go fmt && golint && go test && go run 068.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

import (
	"../util"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
)

// https://projecteuler.net/problem=68
//
// Consider the following "magic" 3-gon ring, filled with the numbers 1 to 6,
// and each line adding to nine.
//
// Working clockwise, and starting from the group of three with the numerically
// lowest external node (4,3,2 in this example), each solution can be described
// uniquely. For example, the above solution can be described by the set:
// 4,3,2; 6,2,1; 5,1,3.
//
// It is possible to complete the ring with four different totals: 9, 10, 11,
// and 12. There are eight solutions in total.
//
// Total	Solution Set
// 9	4,2,3; 5,3,1; 6,1,2
// 9	4,3,2; 6,2,1; 5,1,3
// 10	2,3,5; 4,5,1; 6,1,3
// 10	2,5,3; 6,3,1; 4,1,5
// 11	1,4,6; 3,6,2; 5,2,4
// 11	1,6,4; 5,4,2; 3,2,6
// 12	1,5,6; 2,6,4; 3,4,5
// 12	1,6,5; 3,5,4; 2,4,6
//
// By concatenating each group it is possible to form 9-digit strings; the
// maximum string for a 3-gon ring is 432621513.
//
// Using the numbers 1 to 10, and depending on arrangements, it is possible to
// form 16- and 17-digit strings. What is the maximum 16-digit string for a
// "magic" 5-gon ring?

// findMax returns the max representation in n digits of the given solution
func findMax(d []int, n int) int {
	min := d[0]
	minI := 0

	for i := 0; i < len(d)-1; i += 2 {
		if d[i] < min {
			min = d[i]
			minI = i
		}
	}

	s := ""
	for i := minI; ; {
		a := i
		b := i + 1
		c := i + 3
		if c >= len(d) {
			c = 1
		}
		s += fmt.Sprintf("%d%d%d", d[a], d[b], d[c])
		i += 2
		if i >= len(d) {
			i = 0
		}
		if i == minI {
			break
		}
	}

	f := 0
	if len(s) == n {
		f, _ = strconv.Atoi(s)
	}

	return f
}

// solved tests whether the given representation is a solution
func solved(d []int) bool {
	// Leg 0
	sum := d[0] + d[1] + d[3]

	done := false
	for leg := 1; !done; leg++ {
		a := leg * 2
		b := leg*2 + 1
		c := leg*2 + 3
		if c >= len(d) {
			c = 1
			done = true
		}
		sum2 := d[a] + d[b] + d[c]
		if sum2 != sum {
			return false
		}
	}

	return true
}

// solve returns the max integer representation of n digits
func solve(d, n int) int {
	max := 0

	c := make(chan []int, 1000)
	go util.MakeDigits(d, c)
	for {
		next, ok := <-c
		if !ok {
			break
		}
		if solved(next) {
			m := findMax(next, n)
			if m > max {
				max = m
			}
		}
	}

	return max
}

func main() {
	fmt.Printf("Welcome to 068\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	max := solve(10, 16)
	fmt.Println("Max:", max)
}
