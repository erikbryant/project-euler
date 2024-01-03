package main

import (
	"fmt"
	"strconv"

	"github.com/erikbryant/util-golang/util"
)

func main() {
	sum := 0
	for i := 1; i < 1000*1000; i++ {
		s := strconv.Itoa(i)
		if util.IsPalindromeString(s) {
			s = strconv.FormatInt(int64(i), 2)
			if util.IsPalindromeString(s) {
				fmt.Println(i, " is a base-10 and base-2 palindrome ", s)
				sum += i
			}
		}
	}
	fmt.Println("Sum: ", sum)
}
