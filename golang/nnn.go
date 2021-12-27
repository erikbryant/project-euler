package main

// go fmt && golint && go test && go run nnn.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

import (
	// "../primes"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
)

// func init() {
// 	primes.Load("../primes.gob")
// }

func main() {
	fmt.Printf("Welcome to nnn\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

}
