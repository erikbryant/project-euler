package main

// go fmt ./... && go vet ./... && go test && go run 109.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"sort"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
)

// In the game of darts a player throws three darts at a target board which is
// split into twenty equal sized sections numbered one to twenty.
//
// The score of a dart is determined by the number of the region that the dart
// lands in. A dart landing outside the red/green outer ring scores zero. The
// black and cream regions inside this ring represent single scores. However,
// the red/green outer ring and middle ring score double and treble scores
// respectively.
//
// At the centre of the board are two concentric circles called the bull region,
// or bulls-eye. The outer bull is worth 25 points and the inner bull is a
// double, worth 50 points.
//
// There are many variations of rules but in the most popular game the players
// will begin with a score 301 or 501 and the first player to reduce their
// running total to zero is a winner. However, it is normal to play a "doubles
// out" system, which means that the player must land a double (including the
// double bulls-eye at the centre of the board) on their final dart to win; any
// other dart that would reduce their running total to one or lower means the
// score for that set of three darts is "bust".
//
// When a player is able to finish on their current score it is called a
// "checkout" and the highest checkout is 170: T20 T20 D25 (two treble 20s and
// double bull).
//
// There are exactly eleven distinct ways to checkout on a score of 6:
//
// D3
// D1	D2
// S2	D2
// D2	D1
// S4	D1
// S1	S1	D2
// S1	T1	D1
// S1	S3	D1
// D1	D1	D1
// D1	S2	D1
// S2	S2	D1
//
// D1 D2 is considered different to D2 D1 as they finish on different doubles.
// However, the combination S1 T1 D1 is considered the same as T1 S1 D1.
//
// In addition we shall not include misses in considering combinations; for
// example, D3 is the same as 0 D3 and 0 0 D3.
//
// Incredibly there are 42336 distinct ways of checking out in total.
//
// How many distinct ways can a player checkout with a score less than 100?

type region struct {
	name   string
	score  int
	double bool
}

func makeBoard() []region {
	board := []region{}

	// The numbered wedges
	for i := 1; i <= 20; i++ {
		name := fmt.Sprintf("S%d", i)
		board = append(board, region{name, i, false})
		name = fmt.Sprintf("D%d", i)
		board = append(board, region{name, i * 2, true})
		name = fmt.Sprintf("T%d", i)
		board = append(board, region{name, i * 3, false})
	}

	// The bullseyes
	board = append(board, region{"SB", 25, false})
	board = append(board, region{"DB", 50, true})

	sort.Slice(board, func(i, j int) bool {
		return board[i].score < board[j].score
	})

	return board
}

func checkoutInOne(board []region, score int) []region {
	checkouts := []region{}

	for _, r := range board {
		if r.score == score && r.double {
			checkouts = append(checkouts, r)
		}
	}

	return checkouts
}

func checkoutInTwo(board []region, score int) [][]region {
	checkouts := [][]region{}

	for _, r := range board {
		if r.score < score {
			for _, r2 := range board {
				if r.score+r2.score == score && r2.double {
					checkouts = append(checkouts, []region{r, r2})
				}
			}
		}
	}

	return checkouts
}

func checkoutInThree(board []region, score int) [][]region {
	checkouts := [][]region{}
	used := map[string]bool{}

	for _, r := range board {
		for _, r2 := range board {
			if r2.score > r.score {
				// The combination S1 T1 D1 is considered the same as T1 S1 D1
				break
			}
			for _, r3 := range board {
				if r.score+r2.score+r3.score == score && r3.double && !used[r.name+r2.name] {
					checkouts = append(checkouts, []region{r, r2, r3})
					used[r.name+r2.name] = true
					used[r2.name+r.name] = true
				}
			}
		}
	}

	return checkouts
}

func looper(maxScore int) int {
	count := 0

	board := makeBoard()

	for score := 1; score <= maxScore; score++ {
		// fmt.Println(score)

		check1 := checkoutInOne(board, score)
		count += len(check1)
		// fmt.Println(" 1:", check1)

		check2 := checkoutInTwo(board, score)
		count += len(check2)
		// fmt.Println(" 2:", check2)

		check3 := checkoutInThree(board, score)
		count += len(check3)
		// fmt.Println(" 3:", check3)
	}

	return count
}

func main() {
	fmt.Printf("Welcome to 109\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	maxScore := 100
	count := looper(maxScore - 1)
	fmt.Println("There are", count, "ways to checkout with a score <", maxScore)
}
