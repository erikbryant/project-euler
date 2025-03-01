package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"runtime/pprof"
)

// go fmt ./... && go vet ./... && go test && go run 872.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof
//
// A sequence Tn of rooted trees is constructed such that Tn has n nodes
// numbered from 1 to n.
//
// The sequence starts at T1, a tree with a single node as a root with
// the number 1.
//
// For n > 1, Tn is constructed from Tn-1 using the following procedure:
//
//   1. Trace a path from the root of Tn-1 to a leaf by following the
//   largest-numbered child at each node.
//   2. Remove all edges along the traced path, disconnecting all nodes
//   along it from their parents.
//   3. Connect all orphaned nodes directly to a new node numbered n,
//   which becomes the root of Tn.
//
// Let f(n, k) be the sum of the node numbers along the path connecting the
// root of Tn to the node k, including the root and the node k. For example,
// f(6, 1) = 6+5+1 = 12 and f(10, 3) =29.
//
// Find f(10^17, 9^17).

// After computing trees T1 through T16 it became clear that the nodes
// arrange themselves in a way such that the parent of any node k is k+2^n
// where n is some integer value. For all nodes from k to the root of the tree
// each n is unique and forms a minimal set.

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
)

// pow returns base^exp
func pow(base, exp int) int64 {
	// int64(math.Pow(9, 17)) returns a value one less than the actual answer,
	// so we will implement our own Power function.
	sum := int64(1)
	for i := 0; i < exp; i++ {
		sum = sum * int64(base)
	}
	return sum
}

// f returns the sum of the nodes (inclusive) between k and the root of n
func f(n, k int64) int64 {
	i := k
	sum := k
	for i < n {
		p := math.Log2(float64(n - i))
		i = i + int64(math.Pow(2.0, math.Trunc(p)))
		sum += i
	}
	return sum
}

func main() {
	fmt.Printf("Welcome to 872\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	var n, k int64
	n = pow(10, 17)
	k = pow(9, 17)
	a := f(n, k)
	fmt.Printf("f(%d, %d) = %d\n", n, k, a)
}
