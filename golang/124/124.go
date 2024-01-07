package main

// go fmt ./... && go vet ./... && go test && go run 124.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"sort"

	"github.com/erikbryant/util-golang/algebra"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
)

// The radical of n, rad(n), is the product of the distinct prime factors of n.
// For example, 504 = 23 × 32 × 7, so rad(504) = 2 × 3 × 7 = 42.
//
// If we calculate rad(n) for 1 ≤ n ≤ 10, then sort them on rad(n), and sorting
// on n if the radical values are equal, we get:
//
//	Unsorted	 	Sorted
// 	n	rad(n)	 	n	rad(n)		k
// 	1	1	 				1		1				1
// 	2	2	 				2		2				2
// 	3	3	 				4		2				3
// 	4	2	 				8		2				4
// 	5	5	 				3		3				5
// 	6	6	 				9		3				6
// 	7	7	 				5		5				7
// 	8	2	 				6		6				8
// 	9	3	 				7		7				9
// 	10	10	 			10	10			10
//
// Let E(k) be the kth element in the sorted n column; for example, E(4) = 8
// and E(6) = 9.
//
// If rad(n) is sorted for 1 ≤ n ≤ 100000, find E(10000).

// radical returns the product of the prime factors of n
func radical(n int) int {
	factors := algebra.Factors(n)

	if len(factors) == 0 {
		return n
	}

	product := 1

	for _, factor := range factors {
		product *= factor
	}

	return product
}

type rad struct {
	n       int
	radical int
}

// looper calculates each radical 1..n and returns the sorted result
func looper(n int) []rad {
	rads := []rad{}

	for i := 1; i <= n; i++ {
		r := radical(i)
		rads = append(rads, rad{i, r})
	}

	sort.Slice(rads, func(i, j int) bool {
		if rads[i].radical < rads[j].radical {
			return true
		}
		if rads[i].radical == rads[j].radical {
			return rads[i].n < rads[j].n
		}
		return false
	})

	return rads
}

func main() {
	fmt.Printf("Welcome to 124\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	rads := looper(100 * 1000)
	fmt.Println("E(10,000) =", rads[10*1000-1])
}
