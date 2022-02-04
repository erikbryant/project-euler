package ints_413

import (
	"fmt"
)

func OneChild() int {
	total := 0
	count := 0

	// 1-digit
	divisor := 1
	for i := 0; i <= 9; i++ {
		if i%1 == 0 {
			count++
		}
	}
	total += count
	fmt.Println(1, count, total)

	// 5-digit
	divisor = 5
	for i := 10000; i <= 99999; i++ {
		if i%divisor == 0 {
			count++
		}
	}
	total += count
	fmt.Println(1, count, total)

	return total
}
