package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func reverse(n int) int {
	s := strconv.Itoa(n)

	i := 0
	shift := 1
	for _, digit := range s {
		i += (int(digit) - 48) * shift
		shift *= 10
	}

	return i
}

func odd(n int) bool {
	s := strconv.Itoa(n)

	for _, val := range s {
		if (val-48)&0x01 != 1 {
			return false
		}
	}

	return true
}

func reversible(n int) bool {
	// Leading zeroes are not allowed in either n or reverse(n).
	if n%10 == 0 {
		return false
	}

	return odd(n + reverse(n))
}

func countReversible(end int) int {
	count := 0

	for i := 1; i < end; i++ {
		if reversible(i) {
			count++
		}
	}

	return count
}

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	end := 1000 * 1000 * 1000
	count := countReversible(end)
	fmt.Println("End:", end, "Count:", count)
}
