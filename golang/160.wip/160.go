package main

// go fmt ./... && go vet ./... && go test && go build 160.go && time ./160
// go fmt ./... && go vet ./... && go test && go build 160.go && ./160 && echo top | go tool pprof cpu.prof

import (
	"fmt"
	"log"
	"math"
	"os"
	"runtime/pprof"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// For any N, let f(N) be the last five digits before the trailing zeroes in N!.
// For example,
//
// 9! = 362880 so f(9)=36288
// 10! = 3628800 so f(10)=36288
// 20! = 2432902008176640000 so f(20)=17664
//
// Find f(1,000,000,000,000).

var (
	// Mod is the global digit mask. Don't change this. Unless you hate yourself.
	Mod = 10000000

	// MaxFives is a value greater than k where k is the largest 5^k factor we expect to encounter
	MaxFives = 16
)

func multiply(x, f, twos int) (int, int) {
	for twos < MaxFives && x%2 == 0 {
		twos++
		x /= 2
	}

	for x%5 == 0 {
		twos--
		x /= 5
	}

	x %= Mod
	f *= x
	f %= Mod

	return f, twos
}

func fix(f, twos int) int {
	if twos < 0 || twos > 32 {
		log.Fatal("Twos outside of expected 0-32 range! ", twos)
	}
	return (f << twos) % Mod
}

func factorial(n int) int {
	f := 1
	twos := 0

	for i := 2; i <= n; i++ {
		f, twos = multiply(i, f, twos)
	}
	f = fix(f, twos)

	return f
}

var (
	noTensCache = map[string]int{}
)

func factorialNoTens(start, n int) int {
	startEnd := fmt.Sprintf("%d-%d", start, n)
	fCache, ok := noTensCache[startEnd]
	if ok {
		return fCache
	}

	f := 1
	twos := 0

	for i := start; i <= n; i++ {
		if i%10 == 0 {
			continue
		}
		f, twos = multiply(i, f, twos)
	}
	f = fix(f, twos)

	noTensCache[startEnd] = f

	return f
}

func power(base, exp int) int {
	f := 1
	twos := 0

	for i := 1; i <= exp; i++ {
		f, twos = multiply(base, f, twos)
	}
	f = fix(f, twos)

	return f
}

func computeDataset(d Dataset) {
	pOfP := 1
	fmt.Printf(" Stage Start     Stage End       Product         Count       Product\n")
	for _, stage := range d.stages {
		rp := factorialNoTens(stage.start, stage.end)
		p := power(rp, stage.count)
		pOfP *= p
		pOfP %= Mod
		fmt.Printf("%12d  %12d  %12d  %12d  %12d\n", stage.start, stage.end, rp, stage.count, p)
	}
	fmt.Printf("Product of products: %d  last5: %d\n", pOfP, pOfP%100000)
	if d.expected > 0 && pOfP != d.expected {
		fmt.Printf("FAIL!!!!!!! expected: %d  got: %d\n", d.expected, pOfP%100000)
	}
	fmt.Println()
}

var (
	Elevens = []int{
		1, 1, 11, 101, 1001, 10001, 100001, 1000001, 10000001, 100000001, 1000000001, 10000000001, 100000000001,
	}
	Nines = []int{
		1, 9, 99, 999, 9999, 99999, 999999, 9999999, 99999999, 999999999, 9999999999, 99999999999, 999999999999,
	}
	Offsets = []int{
		0, 10, 110, 1110, 11110, 111110, 1111110, 11111110,
	}
	Expected = []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1946112, 8167808, 3416576,
	}
)

type Stage struct {
	start int
	end   int
	count int
}
type Dataset struct {
	upper    int
	stages   []Stage
	expected int
}

func printDataset(d Dataset) {
	p := message.NewPrinter(language.English)
	fmt.Println(p.Sprintf("%d", d.upper))

	for _, stage := range d.stages {
		fmt.Printf("  %d %d %d\n", stage.start, stage.end, stage.count)
	}

	fmt.Printf("  %d\n\n", d.expected)
}

func makeDataset(upper int) Dataset {
	// The goal is that dataset will look like this:
	// 1,000,000,000,000
	//   1 1 111118
	//   2 9 111117
	//   11 99 111116
	//   101 999 111115
	//   1001 9999 111114
	//   10001 99999 111113
	//   100001 999999 111112
	//   1000001 9999999 111111
	//   -1

	// We count up to upper, but get capped by Mod
	maskDigits := int(math.Log10(float64(Mod))) // + 7 // <===== Set maskDigits to >= 12 and it all works!
	upperDigits := int(math.Log10(float64(upper)))
	limit := min(upperDigits, maskDigits) + 1
	offset := Offsets[max(upperDigits-maskDigits, 0)]

	d := Dataset{upper: upper}
	if upper <= 1000*1000*1000 {
		d.expected = factorial(upper)
	} else {
		d.expected = Expected[upperDigits]
	}

	for r := 0; r < limit; r++ {
		count := (limit - r) + offset
		stage := Stage{start: Elevens[r], end: Nines[r], count: count}
		d.stages = append(d.stages, stage)
	}

	return d
}

func main() {
	fmt.Printf("Welcome to 160\n\n")

	fileHandle, _ := os.Create("cpu.prof")
	pprof.StartCPUProfile(fileHandle)
	defer pprof.StopCPUProfile()

	//upper := 1000 * 1000 * 100
	//dBins := make([]map[int]int, 15)
	//for i := range dBins {
	//	dBins[i] = make(map[int]int)
	//}
	//
	//for i := 1; i <= upper; i++ {
	//	x := i
	//	for x%10 == 0 {
	//		x /= 10
	//	}
	//	x %= Mod
	//	digits := int(math.Log10(float64(x))) + 1
	//	dBins[digits][x]++
	//}
	//
	//for digit, bin := range dBins {
	//	if len(bin) == 0 {
	//		continue
	//	}
	//	fmt.Printf("dBins[%d]: ", digit)
	//	for k, v := range bin {
	//		fmt.Printf("map[%d:%d .. ]\n", k, v)
	//		for _, v2 := range bin {
	//			if v2 != v {
	//				fmt.Println("ERROR!!!!!!! varying bin sizes: ", digit, v, v2)
	//			}
	//		}
	//		break
	//	}
	//}
	//
	//d := makeDataset(upper)
	//printDataset(d)
	//computeDataset(d)

	//upper := 1000 * 1000 * 100
	//d := makeDataset(upper)
	//printDataset(d)
	//computeDataset(d)

	f := factorialNoTens(10001, 99999)
	fmt.Println(f)

	//upper := 1000 * 1000 * 1000 * 1000
	//for tens := 10; tens <= upper; tens *= 10 {
	//	d := makeDataset(tens)
	//	printDataset(d)
	//	computeDataset(d)
	//}

	//upper := 1000 * 1000 * 1000 * 10
	//f := factorial(upper)
	//fmt.Printf("%d! = %d\n", upper, f)

	// If we calculate 99*98*97*... We can expand that to 3-digit numbers
	// Each 2-digit term expands to nine 3-digit terms: For a=99 999*..*991 = (a+1)(a+2)(a+3)(a+4)(a+5)(a+6)(a+7)(a+8)(a+9)
	// a^9 + 45 a^8 + 870 a^7 + 9450 a^6 + 63273 a^5 + 269325 a^4 + 723680 a^3 + 1172700 a^2 + 1026576 a + 362880
	//
	//    999       9
	//     Π a  =   Π (99*10 + k)   =  (a+1)(a+2)(a+3)(a+4)(a+5)(a+6)(a+7)(a+8)(a+9)  for a=99*10
	//   a=991     k=1
	//
	//    999      99    9
	//     Π a  =  Π     Π (a*10 + k)
	//   a=101    a=11  k=1
	//
}
