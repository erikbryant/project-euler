package main

import (
	"fmt"
)

// Given x, find the number of unique ways that numbers sum to x.
// For example, when x=5 there are 6 unique ways to sum.
// {5, [][]int{
// 	[]int{4, 1},
// 	[]int{3, 2},
// 	[]int{3, 1, 1},
// 	[]int{2, 2, 1},
// 	[]int{2, 1, 1, 1},
// 	[]int{1, 1, 1, 1, 1},
// }},

type sumList [][]int

func sums(x int, length int, nextNum int) sumList {
	if nextNum == 0 {
		nextNum = 1
	}

	list := sumList{}

	if x <= 1 {
		return list
	}

	if length <= 1 || length > x {
		return list
	}

	if length == 2 {
		for i := x - 1; i >= x-i; i-- {
			if x-i >= nextNum {
				sum := []int{i, x - i}
				list = append(list, sum)
			}
		}
	} else {

		for i := nextNum; length*i <= x; i++ {
			subLists := sums(x-i, length-1, i)
			for _, subList := range subLists {
				subList = append(subList, i)
				list = append(list, subList)
			}
		}
	}

	return list
}

func countSums(x int) (count uint64) {
	// There are x-1 different length sums (we do not
	// count {x} as a sum). Iterate through each length.
	for length := 2; length <= x; length++ {
		m := uint64(len(sums(x, length, 0)))
		count += m
		// fmt.Println(x, length, m, count)
	}

	return
}

var (
	sumsCache map[int]map[int]map[int]uint64
)

func initCache(max int) {
	sumsCache = make(map[int]map[int]map[int]uint64)

	for x := 0; x <= max; x++ {
		sumsCache[x] = make(map[int]map[int]uint64)
		for length := 2; length <= max; length++ {
			sumsCache[x][length] = make(map[int]uint64)
		}
	}
}

func sumsCounted(x int, length int, nextNum int) (count uint64) {
	if nextNum == 0 {
		nextNum = 1
	}

	if x <= 1 {
		return count
	}

	if length <= 1 || length > x {
		return count
	}

	cached, ok := sumsCache[x][length][nextNum]
	if ok {
		return cached
	}

	if length == 2 {
		for i := x - 1; i >= x-i; i-- {
			if x-i >= nextNum {
				count++
			}
		}
	} else {
		for i := nextNum; length*i <= x; i++ {
			count += sumsCounted(x-i, length-1, i)
		}
	}

	sumsCache[x][length][nextNum] = count
	return count
}

func countSumsFast(x int) (count uint64) {
	initCache(x)

	// There are x-1 different length sums (we do not
	// count {x} as a sum). Iterate through each length.
	for length := 2; length <= x; length++ {
		m := sumsCounted(x, length, 0)
		count += m
		// fmt.Println(x, length, m, count)
	}

	return
}

func main() {
	x := 100
	fmt.Println("x:", x, "sums:", countSumsFast(x))

	// for x := 3; x <= 8; x++ {
	// 	fmt.Println("---", x, "---")
	// 	for l := 2; l <= x; l++ {
	// 		s := sums(x, l, 0)
	// 		fmt.Println(l, s)
	// 	}
	// }
}
