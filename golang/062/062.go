package main

import (
	"fmt"
	"sort"
	"strconv"
)

// Taken from https://siongui.github.io/2017/05/07/go-sort-string-slice-of-rune/
func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

// Taken from https://siongui.github.io/2017/05/07/go-sort-string-slice-of-rune/
func SortStringByCharacter(s string) string {
	r := StringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func main() {
	cubes := make(map[string][]int)
	i := 0
	for {
		val := i * i * i
		s := strconv.Itoa(val)
		s = SortStringByCharacter(s)
		cubes[s] = append(cubes[s], i)
		if len(cubes[s]) == 5 {
			val := cubes[s][0]
			val = val * val * val
			fmt.Println(val)
			break
		}
		i++
	}
}
