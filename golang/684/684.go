package main

// go fmt ./... && go vet ./... && go test && go build 684.go && time ./684

// Define s(n) to be the smallest number that has a digit sum of n.
// For example s(10) = 19.
//
//            k
// Let S(k) = ∑ s(n). You are given S(20) = 1074.
//           n=1
//
// Further let fi be the Fibonacci sequence defined by
// f0 = 0, f1 = 1 and fi = fi-2 + fi-1 for all i ≧ 2.
//
//     90
// Find ∑ S(fi). Give your answer modulo 1,000,000,007.
//     i=2

import (
	"fmt"
)

var (
	tenMod  = int64(0)
	tenMods = []int64{}
)

func exp10Mod(exp, m int64) int64 {
	if len(tenMods) == 0 || tenMod != m {
		tenMod = m
		tenMods = make([]int64, m)
		ten := int64(1)
		for i := int64(0); i < m; i++ {
			tenMods[i] = ten
			ten *= 10
			ten %= m
		}
	}

	return tenMods[exp%(m-1)]
}

func modSub(t1, t2, m int64) int64 {
	t2 %= m
	t1 += m - t2
	return t1 % m
}

func SMod(s, m int64) int64 {
	if s <= 0 {
		return 0
	}

	sum := int64(0)

	// // Sum up to the highest multiple of 9 <= s
	// tens := 1
	// for i := 1; i <= s/9; i++ {
	//	sum += 54*tens - 9
	//	tens *= 10
	// }
	//
	// ^-- Loop unrolling results in a number of the regexp form:
	// 5[9]*4. That is, 54, 594, 5994, 59994, ... and the
	// subtraction of some nines.
	//         54 (- 9)
	//        540 (- 9)
	//       5400 (- 9)
	//      54000 (- 9)
	//   + 540000 (- 9)
	//   --------
	//     599994 (- 9*5)
	//     600000 (- 6) (- 9*5)
	// Add 6 to that to get: 60[0]*. 60, 600, 6000, ...
	// As an equation:
	//    sum = 6 * 10^(s/9) - 6 - (9*(s/9))
	exp := exp10Mod(s/9, m)
	if s/9 >= 1 {
		sum = 6 * exp
		sum %= m
		sum = modSub(sum, 6, m)
		sum = modSub(sum, 9*(s/9), m)
	}

	// Sum from the highest multiple of 9 up to s
	for i := 9*(s/9) + 1; i <= s; i++ {
		sum += (((i % 9) + 1) * exp) % m
		sum %= m
		sum = modSub(sum, 1, m)
	}

	return sum
}

func fib(f int64) int64 {
	if f <= 0 {
		return 0
	}

	a := int64(0)
	b := int64(1)

	for i := int64(2); i <= f; i++ {
		a, b = b, a+b
	}

	return b
}

func sumFibs() {
	lower := int64(2)
	upper := int64(90)
	total := int64(0)
	m := int64(1000000007)

	for s := lower; s <= upper; s++ {
		f := fib(s)
		sum := SMod(f, m)
		//sum := S(int(f)) % m
		total += sum
		total %= m
		fmt.Printf("s: %2d  fib: %20d  sum: %12d  total: %12d\n", s, f, sum, total)
	}

	fmt.Printf("\nFor s = 2..%d  total = %d\n", upper, total)
}

func main() {
	fmt.Printf("Welcome to 684\n\n")

	sumFibs()
}
