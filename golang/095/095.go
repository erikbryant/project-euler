package main

// go fmt ./... && go vet ./... && go test && go run 095.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

var (
	cpuprofile  = flag.String("cpuprofile", "", "write cpu profile to file")
	divisorSums = []int{}
	// Max is the largest value allowed in the chain
	Max = 1000 * 1000
)

// The proper divisors of a number are all the divisors excluding the number
// itself. For example, the proper divisors of 28 are 1, 2, 4, 7, and 14. As the
// sum of these divisors is equal to 28, we call it a perfect number.
//
// Interestingly the sum of the proper divisors of 220 is 284 and the sum of the
// proper divisors of 284 is 220, forming a chain of two numbers. For this
// reason, 220 and 284 are called an amicable pair.
//
// Perhaps less well known are longer chains. For example, starting with 12496,
// we form a chain of five numbers:
//
// 12496 → 14288 → 15472 → 14536 → 14264 (→ 12496 → ...)
//
// Since this chain returns to its starting point, it is called an amicable
// chain.
//
// Find the smallest member of the longest amicable chain with no element
// exceeding one million.

func init() {
	divisorSums = make([]int, Max+1)

	for i := 1; i <= Max/2; i++ {
		for j := 2; i*j <= Max; j++ {
			divisorSums[i*j] += i
		}
	}
}

func in(slice []int, target int) bool {
	for _, val := range slice {
		if val == target {
			return true
		}
	}

	return false
}

// chainer returns min and length for the amicable chain with the given start
func chainer(start int) (int, int, bool) {
	i := start
	chain := []int{start}

	for {
		i = divisorSums[i]
		if i == 0 || i > Max {
			return 0, 0, false
		}
		if in(chain, i) {
			chain = append(chain, i)
			break
		}
		chain = append(chain, i)
	}

	// For this problem a chain is defined as a proper loop. Count only the length
	// of the loop. Discount any tail that led into the loop.
	l := 1
	head := len(chain) - 2
	min := chain[len(chain)-1]
	for ; head >= 0 && chain[head] != chain[len(chain)-1]; head-- {
		l++
		if chain[head] < min {
			min = chain[head]
		}
	}

	return min, l, true
}

// looper returns the min element, max length, and start number if it is a  chain, otherwise Max,0,0
func looper(max int) (int, int, int) {
	maxLength := 0
	minElement := Max
	start := 0

	for i := 1; i <= max; i++ {
		min, length, isChain := chainer(i)
		if !isChain {
			continue
		}
		if length > maxLength || (length == maxLength && min < minElement) {
			maxLength = length
			minElement = min
			start = i
		}
	}

	return minElement, maxLength, start
}

func main() {
	fmt.Printf("Welcome to 095\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	smallest, length, start := looper(Max)
	fmt.Println("The smallest member is:", smallest, "from a chain of length:", length, "started from:", start)
}
