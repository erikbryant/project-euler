package main

import (
	"fmt"

	"github.com/erikbryant/util-golang/util"
)

func main() {
	var distinct [4]int

	fmt.Println("Starting factor search ...")

	consecutive := 4
	i := 1
	depth := 0
	for {
		i++
		factors := util.Factors(i)
		if len(factors) != consecutive {
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
