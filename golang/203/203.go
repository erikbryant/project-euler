package main

// go fmt ./... && go vet ./... && go test && go run 203.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/erikbryant/util-golang/algebra"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
)

// Find the sum of the distinct squarefree numbers in the first 51 rows of
// Pascal's triangle.

// sumDistinctSquareFree returns the sum of the distinct squarefree numbers in
// the given triangle
func sumDistinctSquareFree(rows [][]int) int {
	distinct := map[int]bool{}

	for _, row := range rows {
		for _, val := range row {
			if algebra.SquareFree(val) {
				distinct[val] = true
			}
		}
	}

	sum := 0
	for key := range distinct {
		sum += key
	}

	return sum
}

func main() {
	fmt.Printf("Welcome to 203\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	max := 51
	triangle := algebra.PascalTriangle(max)
	sum := sumDistinctSquareFree(triangle)

	fmt.Println("Sum of the first", max, "rows of distinct squarefree numbers:", sum)
}
