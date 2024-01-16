package main

// go fmt ./... && go vet ./... && go test && go run 122.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

	// Chains stores the minimal addition chains. Each n can have multiple
	// minimal length chains. Store each minimal chain. This way, when we are
	// building the minimal chains for higher numbers we have multiple optimal
	// paths to choose from.
	Chains = [][][]int{}

	// MaxN is the highest n for which we will generate a chain
	MaxN = 200
)

// The most naive way of computing n15 requires fourteen multiplications:
//
// n × n × ... × n = n^15
//
// But using a "binary" method you can compute it in six multiplications:
//
// n × n = n^2
// n^2 × n^2 = n^4
// n^4 × n^4 = n^8
// n^8 × n^4 = n^12
// n^12 × n^2 = n^14
// n^14 × n = n^15
//
// However it is yet possible to compute it in only five multiplications:
//
// n × n = n^2
// n^2 × n = n^3
// n^3 × n^3 = n^6
// n^6 × n^6 = n^12
// n^12 × n^3 = n^15
//
// We shall define m(k) to be the minimum number of multiplications to compute
// n^k; for example m(15) = 5.
//
// For 1 ≤ k ≤ 200, find ∑ m(k).

// OEIS sequence for length of shortest addition chain for n.
// https://oeis.org/A003313

// makeAdditionChains returns all the minimal addition chains for a given n
func makeAdditionChains(n int) [][]int {
	if len(Chains[n]) != 0 {
		return Chains[n]
	}

	candidates := [][]int{}

	// For each i < n...
	//   Iterate over each chain in Chains[i] to find j where j <= i and i+j = n
	//   Save these chains
	// Find the shortest of all evaluated chains

	for i := 1; i < n; i++ {
		for _, chain := range Chains[i] {
			for _, val := range chain {
				if i+val == n {
					candidates = append(candidates, chain)
					break
				}
			}
		}
	}

	// Find the length of the shortest chain(s)
	minLen := len(candidates[0])
	for _, chain := range candidates {
		if len(chain) < minLen {
			minLen = len(chain)
		}
	}

	shortest := [][]int{}

	for _, chain := range candidates {
		if len(chain) <= minLen {
			tmp := make([]int, len(chain))
			copy(tmp, chain)
			tmp = append(tmp, n)
			shortest = append(shortest, tmp)
		}
	}

	return shortest
}

// initChains fills Chains with the minimal addition chains for each value of n
func initChains() {
	Chains = make([][][]int, MaxN+1)

	Chains[0] = append(Chains[0], []int{})
	Chains[1] = append(Chains[1], []int{1})
	Chains[2] = append(Chains[2], []int{1, 2})
	Chains[3] = append(Chains[3], []int{1, 2, 3})
	Chains[4] = append(Chains[4], []int{1, 2, 4})
	Chains[5] = append(Chains[5], []int{1, 2, 3, 5})
	Chains[6] = append(Chains[6], []int{1, 2, 3, 6})
	Chains[7] = append(Chains[7], []int{1, 2, 3, 5, 7})
	Chains[8] = append(Chains[8], []int{1, 2, 4, 8})

	for n := 2; n <= 200; n++ {
		if len(Chains[n]) != 0 {
			continue
		}
		chains := makeAdditionChains(n)
		Chains[n] = chains
	}
}

// lenAdditionChain returns the length of the minimal chain for given n
func lenAdditionChain(n int) int {
	l := len(Chains[n][0])

	if l == 0 {
		return 0
	}

	return l - 1
}

// looper returns the sum of all minimal chains for 1 <= n <= max
func looper(max int) int {
	sum := 0

	for i := 1; i <= max; i++ {
		sum += lenAdditionChain(i)
	}

	return sum
}

func main() {
	fmt.Printf("Welcome to 122\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	initChains()

	maxFound := 200
	sum := looper(maxFound)
	fmt.Println("For 1 <= k <=", maxFound, "the sum of m(k) =", sum)
}
