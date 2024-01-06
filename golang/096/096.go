package main

// go fmt ./... && go vet ./... && go test && go run 096.go

import (
	"fmt"
	"os"
	"strings"

	"github.com/erikbryant/sudoku"
)

// load reads the puzzles file and publishes each board to a channel
func load(c chan sudoku.Board) {
	defer close(c)

	raw, _ := os.ReadFile("p096_sudoku.txt")
	lines := strings.Split(string(raw), "\n")

	// There are 10 lines per each board (one name & nine data)
	for i := 0; i < len(lines); i += 10 {
		b := sudoku.New(lines[i], lines[i+1:])
		c <- b
	}
}

func looper() {
	c := make(chan sudoku.Board, 10)
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

		fmt.Println(b.Name())
		b.Solve()
		if b.Solved() {
			solved++
			fmt.Println("  Solved!!")
			sum += 100*b.Grid(0, 0) + 10*b.Grid(1, 0) + b.Grid(2, 0)
		} else {
			b.Print()
		}
	}

	fmt.Println("Solved", solved, "of", total, "for sum", sum)
}

func main() {
	fmt.Printf("Welcome to 096\n\n")

	looper()
}
