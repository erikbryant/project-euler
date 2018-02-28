package main

import (
	"fmt"
	"math"
	"strconv"
)

func rotate(s string) string {
	return s[1:] + string(s[0])
}

func permute(p1, p2 string) {
	if len(p2) == 0 {
		if prime(p1) {
			fmt.Println(p1)
		}
		return
	}

	for i := 0; i < len(p2); i++ {
		permute(p1+string(p2[0]), p2[1:])
		p2 = rotate(p2)
	}
}

func prime(s string) bool {
	switch s[len(s)-1:] {
	case "0":
		return false
	case "2":
		return false
	case "4":
		return false
	case "5":
		return false
	case "6":
		return false
	case "8":
		return false
	}

	n, _ := strconv.Atoi(s)
	mid := int(math.Sqrt(float64(n)))
	for i := 3; i <= mid; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	s := "987654321"
	for len(s) > 0 {
		permute("", s)
		s = s[1:]
	}
}
