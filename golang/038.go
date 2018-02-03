package main

import (
	"fmt"
	"strconv"
	"strings"
)

func pandigital9(s string) bool {
	for i := 1; i <= 9; i++ {
		if !strings.Contains(s, string(i+'0')) {
			return false
		}
	}
	return true
}

func main() {
	i := 1
	max := 0
	for i < 987654322/2 {
		n := 1
		s := ""
		for len(s) < 9 {
			product := i * n
			s += strconv.Itoa(product)
			n++
		}
		if len(s) == 9 {
			num, _ := strconv.Atoi(s)
			if num > max && pandigital9(s) {
				max = num
				fmt.Println("i: ", i, " 1:", n, " s: ", s)
			}
		}
		i++
	}
	fmt.Println("max: ", max)
}
