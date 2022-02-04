package main

import (
	"../primes"
	"fmt"
	"strconv"
)

func truncate(p string) bool {
	if len(p) < 2 {
		return false
	}
	s := p
	for len(s) > 0 {
		n, _ := strconv.Atoi(s)
		if !primes.Prime(n) {
			return false
		}
		s = s[1:]
	}

	s = p
	for len(s) > 0 {
		n, _ := strconv.Atoi(s)
		if !primes.Prime(n) {
			return false
		}
		s = s[:len(s)-1]
	}

	return true
}

func main() {
	primes.Init(false)

	count := 0
	sum := 0

	for i, prime := range primes.Primes {
		if !prime {
			continue
		}
		s := strconv.Itoa(i)
		if truncate(s) {
			fmt.Println(s)
			count++
			n, _ := strconv.Atoi(s)
			sum += n
		}
	}
	fmt.Println("Count: ", count)
	fmt.Println("Sum  : ", sum)
}
