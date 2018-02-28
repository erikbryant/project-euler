package main

import (
	"fmt"
	"strconv"
)

func palindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	sum := 0
	for i := 1; i < 1000*1000; i++ {
		s := strconv.Itoa(i)
		if palindrome(s) {
			s = strconv.FormatInt(int64(i), 2)
			if palindrome(s) {
				fmt.Println(i, " is a base-10 and base-2 palindrome ", s)
				sum += i
			}
		}
	}
	fmt.Println("Sum: ", sum)
}
