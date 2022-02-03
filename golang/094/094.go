package main

// go fmt && golint && go test && go run 094.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"runtime/pprof"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
)

// It is easily proved that no equilateral triangle exists with integral length
// sides and integral area. However, the almost equilateral triangle 5-5-6 has
// an area of 12 square units.
//
// We shall define an almost equilateral triangle to be a triangle for which two
// sides are equal and the third differs by no more than one unit.
//
// Find the sum of the perimeters of all almost equilateral triangles with
// integral side lengths and area and whose perimeters do not exceed one billion
// (1,000,000,000).

// perimeter returns the perimeter of a triangle with sides s1, s1, s2
func perimeter(s1, s2 int) int {
	return s1 + s1 + s2
}

// areaIntHero returns true if the area of the triangle is an integer
func areaIntHero(s1, s2 int) (int, bool) {
	// https://en.wikipedia.org/wiki/Integer_triangle#Area_of_an_integer_triangle
	// If T is the area of the triangle then:
	//   T = sqrt( (a+b+c)*(a+b-c)*(a-b+c)*(-a+b+c) ) / 4
	// In our case a==b, so this reduces to:
	//   T = sqrt( (a+a+c)*(a+a-c)*(c)*(c) ) / 4
	//   T = c * sqrt( (a+a+c)*(a+a-c) ) / 4

	term := (s1 + s1 + s2) * (s1 + s1 - s2)
	root := int(math.Sqrt(float64(term)))

	// If term is not a square then its root will be irrational, which cannot
	// make for an integer area
	if root*root != term {
		return 0, false
	}

	area := float64(root) * float64(s2) / 4.0

	// Is area an integer?
	if math.Trunc(area) != area {
		return 0, false
	}

	return int(area), true
}

func looper() int {
	sum := 0

	for s1 := 2; ; s1++ {
		s2 := s1 - 1
		if perimeter(s1, s2) > 1000*1000*1000 {
			break
		}
		if _, ok := areaIntHero(s1, s2); ok {
			fmt.Printf("Integral area -1: %10d %10d\n", s1, s2)
			sum += perimeter(s1, s2)
		}

		s2 = s1 + 1
		if perimeter(s1, s2) > 1000*1000*1000 {
			break
		}
		if _, ok := areaIntHero(s1, s2); ok {
			fmt.Printf("Integral area +1: %10d %10d\n", s1, s2)
			sum += perimeter(s1, s2)
		}
	}

	return sum
}

func main() {
	fmt.Printf("Welcome to 094\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	perimeterSum := looper()
	fmt.Println("Perimeter sum:", perimeterSum)
}
