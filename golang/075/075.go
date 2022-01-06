package main

// go fmt && golint && go test && go run 075.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
)

// https://projecteuler.net/problem=75
//
// It turns out that 12 cm is the smallest length of wire that can be bent to
// form an integer sided right angle triangle in exactly one way, but there are
// many more examples.
//
// 12 cm: (3,4,5)
// 24 cm: (6,8,10)
// 30 cm: (5,12,13)
// 36 cm: (9,12,15)
// 40 cm: (8,15,17)
// 48 cm: (12,16,20)
//
// In contrast, some lengths of wire, like 20 cm, cannot be bent to form an
// integer sided right angle triangle, and other lengths allow more than one
// solution to be found; for example, using 120 cm it is possible to form
// exactly three different integer sided right angle triangles.
//
// 120 cm: (30,40,50), (20,48,52), (24,45,51)
//
// Given that L is the length of the wire, for how many values of L ≤ 1,500,000
// can exactly one integer sided right angle triangle be formed?

// triples counts all lengths <= maxLen that have only one triple solution where a+b+c <= maxLen and a^2+b^2=c^2
func triples(maxLen int) int {
	nSmallEnough := true

	lengths := make([]int, maxLen+1)
	found := make(map[int]map[int]bool)

	for n := 1; nSmallEnough; n++ {
		nSmallEnough = false
		for m := n + 1; ; m++ {
			a := m*m - n*n
			b := 2 * m * n
			c := m*m + n*n
			if a+b+c > maxLen {
				break
			}
			nSmallEnough = true
			for k := 1; ; k++ {
				ak := k * a
				bk := k * b
				ck := k * c
				if ak+bk+ck > maxLen {
					break
				}
				if !found[ak][bk] {
					lengths[ak+bk+ck]++
					if found[ak] == nil {
						found[ak] = make(map[int]bool)
					}
					if found[bk] == nil {
						found[bk] = make(map[int]bool)
					}
					found[ak][bk] = true
					found[bk][ak] = true
				}
			}
		}
	}

	// Find the wire lengths that had only one solution
	count := 0
	for _, l := range lengths {
		if l == 1 {
			count++
		}
	}

	return count
}

func main() {
	fmt.Printf("Welcome to 075\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	L := 1500000
	count := triples(L)
	fmt.Println("For L <=", L, count, "triangles have integer sides")
}
