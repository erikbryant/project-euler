package main

import (
	"./ints_413"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"sort"
	"strconv"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

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

	count := ints_413.OneChild()
	fmt.Println(count)
	return

	// // Digits using division
	// sum := 0
	// for i := 10; i <= 10; i += 2 {
	// 	count := router(i, i)
	// 	sum += count
	// 	fmt.Println(i, count, sum)
	// }
	// return

	// // Tabular one-child counts using division
	// maxLen := 8
	// fmt.Println("One-child counts")
	// fmt.Println()
	// fmt.Printf("Divisor \\ Length")
	// for length := 1; length <= maxLen; length++ {
	// 	fmt.Printf("%12d", length)
	// }
	// fmt.Println()
	// for divisor := 11; divisor <= 19; divisor++ {
	// 	fmt.Printf("%16d   ", divisor)
	// 	for length := 1; length <= maxLen; length++ {
	// 		count := router(length, divisor)
	// 		fmt.Printf("%12d", count)
	// 	}
	// 	fmt.Println()
	// }
	// return

	// // Tabular one-child counts using strings
	// maxLenA := 4
	// fmt.Println("One-child counts")
	// fmt.Println()
	// fmt.Printf("Divisor \\ Length")
	// for length := 2; length <= maxLenA; length++ {
	// 	fmt.Printf("%12d", length)
	// }
	// fmt.Println()
	// for divisor := 2; divisor <= 19; divisor++ {
	// 	fmt.Printf("%16d   ", divisor)
	// 	for length := 2; length <= maxLenA && length <= divisor; length++ {
	// 		count := strChildrenByRank(length, divisor)
	// 		fmt.Printf("%12d", count)
	// 	}
	// 	fmt.Println()
	// }
	// return

	// // Strings using multiples by rank
	// suma := 0
	// for i := 1; i <= 6; i++ {
	// 	count := strChildrenByRank(i, i)
	// 	suma += count
	// 	fmt.Println(i, count, suma)
	// }
	// return

}
