package main

import (
	"fmt"

	"github.com/erikbryant/util-golang/figurate"
)

func main() {
	n := 1
	for {
		triangle := (n*n + n) / 2
		if figurate.IsPentagonal(triangle) && figurate.IsHexagonal(triangle) {
			fmt.Println(n, " (", triangle, ") is Tri + Pent + Hex")
			if triangle > 40755 {
				break
			}
		}
		n++
	}
}
