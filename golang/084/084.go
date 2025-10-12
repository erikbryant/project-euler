package main

//  go fmt ./... && go vet ./... && go test && go build 084.go && time ./084

import (
	"fmt"
	"log"
	"math"
	"slices"

	"github.com/erikbryant/util-golang/matrices"
)

//  In the game, Monopoly, the standard board is set up in the following way:
//
// https://projecteuler.net/resources/images/0084_monopoly_board.png?1678992052
//
// A player starts on the GO square and adds the scores on two 6-sided dice to determine
// the number of squares they advance in a clockwise direction. Without any further rules
// we would expect to visit each square with equal probability: 2.5%. However, landing on
// G2J (Go To Jail), CC (community chest), and CH (chance) changes this distribution.
//
// In addition to G2J, and one card from each of CC and CH, that orders the player to go
// directly to jail, if a player rolls three consecutive doubles, they do not advance the
// result of their 3rd roll. Instead, they proceed directly to jail.
//
// At the beginning of the game, the CC and CH cards are shuffled. When a player lands on
// CC or CH they take a card from the top of the respective pile and, after following the
// instructions, it is returned to the bottom of the pile. There are sixteen cards in each
// pile, but for the purpose of this problem we are only concerned with cards that order a
// movement; any instruction not concerned with movement will be ignored and the player will
// remain on the CC/CH square.
//
//   Community Chest (2/16 cards):
//   * Advance to GO
//   * Go to JAIL
//
//   chance (10/16 cards):
//   * Advance to GO
//   * Go to JAIL
//   * Go to C1
//   * Go to E3
//   * Go to H2
//   * Go to R1
//   * Go to next R (railway company)
//   * Go to next R
//   * Go to next U (utility company)
//   * Go back 3 squares
//
// The heart of this problem concerns the likelihood of visiting a particular square.
// That is, the probability of finishing at that square after a roll. For this reason
// it should be clear that, with the exception of G2J for which the probability of
// finishing on it is zero, the CH squares will have the lowest probabilities, as 5/8
// request a movement to another square, and it is the final square that the player
// finishes at on each roll that we are interested in. We shall make no distinction
// between "Just Visiting" and being sent to JAIL, and we shall also ignore the rule
// about requiring a double to "get out of jail", assuming that they pay to get out on
// their next turn.
//
// By starting at GO and numbering the squares sequentially from 00 to 39 we can
// concatenate these two-digit numbers to produce strings that correspond with sets of
// squares.
//
// Statistically it can be shown that the three most popular squares, in order, are
// JAIL (6.24%) = Square 10, E3 (3.18%) = Square 24, and GO (3.09%) = Square 00. So these
// three most popular squares can be listed with the six-digit modal string: 102400.
//
// If, instead of using two 6-sided dice, two 4-sided dice are used, find the six-digit
// modal string.

// Note:
// My results are slightly off from the given values. Close enough to make it all work,
// but still not exact. Not sure what the issue is. Anyone?

var (
	SquareNames = []string{
		"GO", "A1", "CC1", "A2", "T1", "R1", "B1", "CH1", "B2", "B3", "JL",
		"C1", "U1", "C2", "C3", "R2", "D1", "CC2", "D2", "D3",
		"FP", "E1", "CH2", "E2", "E3", "R3", "F1", "F2", "U2", "F3", "G2J",
		"G1", "G2", "CC3", "G3", "R4", "CH3", "H1", "T2", "H2",
	}
)

// name returns the textual name of that square
func name(i int) string {
	return SquareNames[i]
}

// index returns the index of the named square
func index(name string) int {
	for i, n := range SquareNames {
		if n == name {
			return i
		}
	}
	return -1
}

// boardNew returns a new, empty board (matrix) and an initial state matrix
func boardNew() (matrices.Matrix, matrices.Matrix) {
	rowsA := 40
	colsA := rowsA
	board := matrices.New(rowsA, colsA)
	start := matrices.New(1, colsA)
	start[0][0] = 1.0
	return board, start
}

// boardCheck aborts if a row does not sum to 1.0
func boardCheck(A matrices.Matrix, title string) bool {
	epsilon := 0.00001

	for row := 0; row < A.Rows(); row++ {
		p := 0.0
		for col := 0; col < A.Cols(); col++ {
			p += A[row][col]
		}
		if math.Abs(1.0-p) > epsilon {
			A.Print(title, name)
			log.Fatalf("ERROR! boardCheck failed for row %d, p=%f\n", row, p)
		}
	}
	return true
}

// roll2Dice fills in [weighted] movement probabilities for a row
func roll2Dice(A matrices.Matrix, row, dieSides int) {
	n := float64(dieSides)
	C1 := index("C1")
	CC1 := index("CC1")
	CC2 := index("CC2")
	CC3 := index("CC3")
	CH1 := index("CH1")
	CH2 := index("CH2")
	CH3 := index("CH3")
	D3 := index("D3")
	E3 := index("E3")
	G2J := index("G2J")
	GO := index("GO")
	H2 := index("H2")
	JL := index("JL")
	R1 := index("R1")
	R2 := index("R2")
	R3 := index("R3")
	T1 := index("T1")
	U1 := index("U1")
	U2 := index("U2")

	numDoubles := n
	p1Double := numDoubles / (n * n)
	p3rdDouble := p1Double * p1Double * p1Double

	// Single roll of 2 dice
	for d1 := 1; d1 <= dieSides; d1++ {
		for d2 := 1; d2 <= dieSides; d2++ {
			sum := d1 + d2
			p := 1.0 / (n * n)
			if d1 == d2 {
				// Adjust for probability of two previous doubles
				p -= p3rdDouble / numDoubles
				A[row][JL] += p3rdDouble / numDoubles
			}
			nextSquare := (row + sum) % A.Cols()
			switch nextSquare {
			case G2J:
				A[row][JL] += p
			case CC1:
				A[row][CC1] += p * 14.0 / 16.0
				A[row][GO] += p * 1.0 / 16.0
				A[row][JL] += p * 1.0 / 16.0
			case CC2:
				A[row][CC2] += p * 14.0 / 16.0
				A[row][GO] += p * 1.0 / 16.0
				A[row][JL] += p * 1.0 / 16.0
			case CC3:
				A[row][CC3] += p * 14.0 / 16.0
				A[row][GO] += p * 1.0 / 16.0
				A[row][JL] += p * 1.0 / 16.0
			case CH1:
				A[row][CH1] += p * 6.0 / 16.0
				A[row][GO] += p * 1.0 / 16.0
				A[row][JL] += p * 1.0 / 16.0
				A[row][C1] += p * 1.0 / 16.0
				A[row][E3] += p * 1.0 / 16.0
				A[row][H2] += p * 1.0 / 16.0
				A[row][R1] += p * 1.0 / 16.0
				A[row][R2] += p * 1.0 / 16.0 // next railway
				A[row][R2] += p * 1.0 / 16.0 // next railway
				A[row][U1] += p * 1.0 / 16.0 // next utility
				A[row][T1] += p * 1.0 / 16.0 // back 3 squares
			case CH2:
				A[row][CH2] += p * 6.0 / 16.0
				A[row][GO] += p * 1.0 / 16.0
				A[row][JL] += p * 1.0 / 16.0
				A[row][C1] += p * 1.0 / 16.0
				A[row][E3] += p * 1.0 / 16.0
				A[row][H2] += p * 1.0 / 16.0
				A[row][R1] += p * 1.0 / 16.0
				A[row][R3] += p * 1.0 / 16.0 // next railway
				A[row][R3] += p * 1.0 / 16.0 // next railway
				A[row][U2] += p * 1.0 / 16.0 // next utility
				A[row][D3] += p * 1.0 / 16.0 // back 3 squares
			case CH3:
				A[row][CH3] += p * 6.0 / 16.0
				A[row][GO] += p * 1.0 / 16.0
				A[row][JL] += p * 1.0 / 16.0
				A[row][C1] += p * 1.0 / 16.0
				A[row][E3] += p * 1.0 / 16.0
				A[row][H2] += p * 1.0 / 16.0
				A[row][R1] += p * 1.0 / 16.0
				A[row][R1] += p * 1.0 / 16.0 // next railway
				A[row][R1] += p * 1.0 / 16.0 // next railway
				A[row][U1] += p * 1.0 / 16.0 // next utility
				//A[row][CC3] += p * 1.0 / 16.0 // back 3 squares
				p2 := p * 1.0 / 16.0
				A[row][CC3] += p2 * 14.0 / 16.0
				A[row][GO] += p2 * 1.0 / 16.0
				A[row][JL] += p2 * 1.0 / 16.0
			default:
				A[row][nextSquare] += p
			}
		}
	}
}

// goToJail fills in the G2J row
func goToJail(A matrices.Matrix) {
	G2J := index("G2J")
	JL := index("JL")
	A.SetRow(G2J, 0.0)
	A[G2J][JL] = 1.0
}

// boardInit sets the transition probabilities for each square
func boardInit(A matrices.Matrix, dieSides int) {
	// Initialize all rows with the default transition
	for row := 0; row < A.Cols(); row++ {
		roll2Dice(A, row, dieSides)
	}

	goToJail(A)

	boardCheck(A, "Board after init")
}

// transition applies the transition matrix steps times
func transition(board, state matrices.Matrix, steps int) {
	nextState := state.Copy()

	for move := 1; move <= steps; move++ {
		state.Mul(board, nextState)
		state, nextState = nextState, state
		boardCheck(state, "Latest position probability state")
	}
}

// leaders ranks each square by how frequently a turn ended on it
func leaders(state matrices.Matrix) {
	values := []float64{}

	for col := 0; col < state.Cols(); col++ {
		values = append(values, state[0][col])
	}

	slices.Sort(values)
	slices.Reverse(values)

	fmt.Printf("\nRankings:\n")
	for _, value := range values {
		for col := 0; col < state.Cols(); col++ {
			if state[0][col] == value {
				fmt.Printf("%02d   %3s  %6.2f%%\n", col, name(col), 100.0*value)
			}
		}
	}
}

func main() {
	fmt.Printf("Welcome to 084\n\n")

	dieSides := 4
	steps := 1000

	board, state := boardNew()
	boardInit(board, dieSides)
	board.Print("Board Initialized (transition matrix)", name)
	state.Print("Start State", name)
	transition(board, state, steps)
	msg := fmt.Sprintf("State After %d step(s)", steps)
	state.Print(msg, name)

	leaders(state)
}
