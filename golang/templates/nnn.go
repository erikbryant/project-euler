package main

// go fmt ./... && go vet ./... && go test && go run nnn.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

import (
	// "github.com/erikbryant/util-golang/primes"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
)

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
