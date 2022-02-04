package main

import (
	"fmt"
	"math"
)

var (
	squares       map[int]struct{}
	squaresSorted []int
)

func init() {
	squares = make(map[int]struct{})
	root := int(math.Sqrt(float64(1000 * 1000 * 1000 * 1000)))
	for i := 1; i <= root+1; i++ {
		square := i * i
		squares[square] = struct{}{}
		squaresSorted = append(squaresSorted, square)
	}
}

func printMatrix(N int) {
	root := int(math.Sqrt(float64(N)))

	closed := make(map[int]bool)

	for a := 0; a < root; a++ {
		for b := a + 1; b <= root; b++ {
			action := squaresSorted[a] + squaresSorted[b]
			closed[action] = !closed[action]
		}
	}

	fmt.Printf("       ")
	for b := 2; b <= root; b++ {
		fmt.Printf("%3d ", b)
	}
	fmt.Printf("\n")

	for a := 0; a < root; a++ {
		fmt.Printf("a:%3d  ", a+1)
		for b := a + 1; b <= root; b++ {
			action := squaresSorted[a] + squaresSorted[b]
			if action > N {
				break
			}
			a := ""
			if closed[action] {
				a = ""
			}
			fmt.Printf("%3d%s ", action, a)
		}
		fmt.Printf("\n")
	}
}

func walkAllSum(N int, c chan int) {
	defer close(c)

	sum := 3

	for {
		for a := (sum - 1) >> 1; a >= 1; a-- {
			// b := sum - a
			// action := a*a + b*b
			action := squaresSorted[a-1] + squaresSorted[sum-a-1]
			if action > N {
				break
			}
			c <- action
		}
		if squaresSorted[sum] > N<<1 {
			// TODO: This may terminate
			// too early or too late
			break
		}
		sum++
	}
}

func walkAllLoop(N int, c chan int) {
	defer close(c)

	root := int(math.Sqrt(float64(N)))

	for a := 0; a < root; a++ {
		for b := a + 1; b <= root; b++ {
			action := squaresSorted[a] + squaresSorted[b]
			if action > N {
				break
			}
			c <- action
		}
	}
}

func walkCollapsed(N int) int {
	all := make(chan int, 1000)
	go walkAllSum(N, all)

	doors := make(map[int]struct{})

	max := 0
	sorted := true
	count := 0
	iter := 0

	for {
		action, ok := <-all
		if !ok {
			break
		}
		if action <= max {
			_, ok = doors[action]
			if ok {
				delete(doors, action)
			} else {
				doors[action] = struct{}{}
				iter++
			}
			if sorted && iter > 1000*100 {
				// Prune the ones that are smaller than
				// we will see again
				for key := range doors {
					if key < action {
						delete(doors, key)
						count++
					}
				}
				iter = 0
			}
			sorted = false
		} else {
			max = action
			doors[action] = struct{}{}
			sorted = true
		}
	}

	return count + len(doors)
}

func walk(N int) int {
	return walkCollapsed(N)
}

func main() {
	// printMatrix(500)

	doors := 1000 * 1000 * 1000 // open: 68496000
	open := walk(doors)
	fmt.Println("Doors =", doors, "Open =", open)
}
