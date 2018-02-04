package main

import (
	"fmt"
	"strconv"
)

func terminus(n int) int {
	for {
		s := strconv.Itoa(n)
		sum := 0
		for _, c := range s {
			d := c - '0'
			sum += int(d) * int(d)
		}
		if sum == n || sum == 89 {
			return sum
		}
		n = sum
	}
}

func main() {
	one := 0
	eightynine := 0
	for i := 1; i < 1000*1000*10; i++ {
		t := terminus(i)
		switch t {
		case 1:
			one++
		case 89:
			eightynine++
		default:
			fmt.Println("ERROR: i: ", i, " terminus: ", t)
		}
	}
	fmt.Println("1: ", one, " 89: ", eightynine)
}
