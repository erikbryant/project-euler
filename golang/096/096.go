package main

// go fmt ./... && go vet ./... && go test && go run 096.go

import (
	"fmt"
	"os"
	"strings"
)

// load reads the puzzles file and publishes each board to a channel
func load(c chan Board) {
	defer close(c)

	raw, _ := os.ReadFile("p096_sudoku.txt")
	lines := strings.Split(string(raw), "\n")

	// There are 10 lines per each board (one name & nine data)
	for i := 0; i < len(lines); i += 10 {
		b := New(lines[i], lines[i+1:])
		c <- b
	}
}

func looper() {
	c := make(chan Board, 10)
	go load(c)

	solved := 0
	total := 0
	sum := 0

	for {
		// Read a board from the channel
		b, ok := <-c
		if !ok {
			break
		}

		total++

		fmt.Println(b.name)
		b.solve()
		if b.solved() {
			solved++
			fmt.Println("  Solved!!")
			sum += 100*b.grid[0][0] + 10*b.grid[1][0] + b.grid[2][0]
		} else {
			b.print()
		}
	}

	fmt.Println("Solved", solved, "of", total, "for sum", sum)
}

func main() {
	fmt.Printf("Welcome to 096\n\n")

	looper()
}
