package main

import (
	"fmt"
)

func triangle(n int) {
	triangle := 0

	i := 1
	for {
		fmt.Println(triangle)
		if triangle+i > n {
			break
		}
		triangle += i
		i++
	}
	fmt.Println(triangle)
}

func main() {
	fmt.Println("Welcome to 621")
	triangle(17526 * 1000 * 1000 * 1000)
}
