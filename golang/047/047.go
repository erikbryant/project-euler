package main

import (
	"github.com/erikbryant/project-euler/golang/primes"
	"fmt"
)

func factor(n int) []int {
	factors := make([]int, 0, 40)

	if n == 1 {
		return factors
	}

	for i := 2; i <= n; i++ {
		if primes.Prime(i) && n%i == 0 {
			factors = append(factors, i)
			f := factor(int(n / i))
			if f == nil {
				return nil
			}
			for _, val := range f {
				factors = append(factors, val)
			}
			return factors
		}
	}

	return nil
}

func countDistinct(factors []int) int {
	if factors == nil {
		return 0
	}

	distinct := make(map[int]int)
	for i := 0; i < len(factors); i++ {
		distinct[factors[i]]++
	}
	return len(distinct)
}

func main() {
	primes.Init(false)

	var distinct [4]int

	fmt.Println("Starting factor search ...")

	consecutive := 4
	i := 1
	depth := 0
	for {
		i++
		factors := factor(i)
		if len(factors) < consecutive || countDistinct(factors) != consecutive {
			depth = 0
			continue
		}
		if depth > 0 && distinct[depth-1]+1 != i {
			depth = 0
		}
		distinct[depth] = i
		depth++
		if depth == consecutive {
			fmt.Println(distinct)
			break
		}
	}
}
