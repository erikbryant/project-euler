package main

import (
	"fmt"
	"os"
	"strings"
)

// Card represents a single playing card
type Card struct {
	face int
	suit rune
}

// Hand represents a hand of cards
type Hand [5]Card

var (
	intToFace = map[int]string{2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "T", 11: "J", 12: "Q", 13: "K", 14: "A"}
	faceToInt = map[string]int{"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "T": 10, "J": 11, "Q": 12, "K": 13, "A": 14}
)

// high() returns the nth highest card.
// ASSUMES: cards are sorted in descending order.
func (hand *Hand) highCard(n int) int {
	if n < 1 || n > len(hand) {
		fmt.Printf("ERROR: Expected 1 <= n <= %d, got %d\n", len(hand)-1, n)
		panic("out of bounds")
	}

	// Duplicates are ignored, so in [9,8,8,5,2]
	// the 5 is the 3rd highest card.
	for i := 0; i < len(hand); i++ {
		if n == 1 {
			return hand[i].face
		}
		if i < len(hand)-1 && hand[i].face != hand[i+1].face {
			n--
		}
	}

	return 0
}

func (hand *Hand) minDuplicates(min int) (bool, int) {
	count := make(map[int]int)

	for i := 0; i < len(hand); i++ {
		count[hand[i].face]++
	}

	for key, val := range count {
		if val >= min {
			return true, key
		}
	}

	return false, 0
}

func (hand *Hand) onePair() (bool, int) {
	return hand.minDuplicates(2)
}

func (hand *Hand) twoPairs() (bool, int, int) {
	count := make(map[int]int)
	two := false
	twoFace := 0
	twoAgain := false
	twoAgainFace := 0

	for i := 0; i < len(hand); i++ {
		count[hand[i].face]++
	}

	for key, val := range count {
		if val >= 2 {
			if two {
				twoAgain = true
				twoAgainFace = key
			} else {
				two = true
				twoFace = key
			}
		}
	}

	if twoFace < twoAgainFace {
		twoFace, twoAgainFace = twoAgainFace, twoFace
	}
	return two && twoAgain, twoFace, twoAgainFace
}

func (hand *Hand) threeOfAKind() (bool, int) {
	return hand.minDuplicates(3)
}

// ASSUMES: cards are sorted in descending order.
func (hand *Hand) straight() bool {
	for i := 0; i < len(hand)-1; i++ {
		if hand[i].face != hand[i+1].face+1 {
			return false
		}
	}
	return true
}

// flush() checks whether the hand has a flush.
// A flush is all cards are of the same suit.
func (hand *Hand) flush() bool {
	for _, card := range hand {
		if card.suit != hand[0].suit {
			return false
		}
	}

	return true
}

func (hand *Hand) fullHouse() bool {
	count := make(map[int]int)
	three := false
	two := false

	for i := 0; i < len(hand); i++ {
		count[hand[i].face]++
	}

	for _, val := range count {
		if val == 3 {
			three = true
		}
		if val == 2 {
			two = true
		}
	}

	return three && two
}

func (hand *Hand) fourOfAKind() (bool, int) {
	return hand.minDuplicates(4)
}

// straightFlush() checks whether the hand has a straight flush.
// A straight flush is all cards are consecutive values of same suit.
// ASSUMES: cards are sorted in descending order.
func (hand *Hand) straightFlush() bool {
	if !hand.flush() {
		return false
	}
	for i := 0; i < len(hand)-1; i++ {
		if hand[i].face != hand[i+1].face+1 {
			return false
		}
	}
	return true
}

// royalFlush() checks whether the hand has a straight flush.
// A royal flush is a straight flush with Ace high.
func (hand *Hand) royalFlush() bool {
	return hand.straightFlush() && hand.highCard(1) == faceToInt["A"]
}

func winnerHighCard(h1, h2 Hand) int {
	i := 1
	for i = 1; i <= 5; i++ {
		if h1.highCard(i) != h2.highCard(i) {
			break
		}
	}
	if h1.highCard(i) > h2.highCard(i) {
		return 1
	}
	return 2
}

func resolve(oneHas, twoHas bool, high1, high2 int, h1, h2 Hand) int {
	if oneHas && twoHas {
		if high1 == high2 {
			return winnerHighCard(h1, h2)
		}
		if high1 > high2 {
			return 1
		}
		return 2
	}
	if oneHas {
		return 1
	}
	if twoHas {
		return 2
	}

	fmt.Println("ERROR!")
	return 0
}

// winner() compares two hands are returns which one won (1 or 2). It assumes there are no ties.
func winner(h1, h2 Hand) int {
	// Royal flush
	if h1.royalFlush() {
		return 1
	}
	if h2.royalFlush() {
		return 2
	}

	// Straight flush
	if h1.straightFlush() && h2.straightFlush() {
		return winnerHighCard(h1, h2)
	}
	if h1.straightFlush() {
		return 1
	}
	if h2.straightFlush() {
		return 2
	}

	// Four of a kind
	oneHas, high1 := h1.fourOfAKind()
	twoHas, high2 := h2.fourOfAKind()
	if oneHas || twoHas {
		return resolve(oneHas, twoHas, high1, high2, h1, h2)
	}

	// Full house
	if h1.fullHouse() && h2.fullHouse() {
		return winnerHighCard(h1, h2)
	}
	if h1.fullHouse() {
		return 1
	}
	if h2.fullHouse() {
		return 2
	}

	// Flush
	if h1.flush() && h2.flush() {
		return winnerHighCard(h1, h2)
	}
	if h1.flush() {
		return 1
	}
	if h2.flush() {
		return 2
	}

	// Straight
	if h1.straight() && h2.straight() {
		return winnerHighCard(h1, h2)
	}
	if h1.straight() {
		return 1
	}
	if h2.straight() {
		return 2
	}

	// Three of a kind
	oneHas, high1 = h1.threeOfAKind()
	twoHas, high2 = h2.threeOfAKind()
	if oneHas || twoHas {
		return resolve(oneHas, twoHas, high1, high2, h1, h2)
	}

	// Two pairs
	low1 := 0
	low2 := 0
	oneHas, high1, low1 = h1.twoPairs()
	twoHas, high2, low2 = h2.twoPairs()
	if oneHas || twoHas {
		if oneHas && twoHas {
			if high1 > high2 {
				return 1
			}
			if high2 > high1 {
				return 2
			}
			if low1 > low2 {
				return 1
			}
			if low2 > low1 {
				return 2
			}
			return winnerHighCard(h1, h2)
		}
		if oneHas {
			return 1
		}
		return 2
	}

	// One pair
	oneHas, high1 = h1.onePair()
	twoHas, high2 = h2.onePair()
	if oneHas || twoHas {
		return resolve(oneHas, twoHas, high1, high2, h1, h2)
	}

	// High card
	return winnerHighCard(h1, h2)
}

func sort(h Hand) Hand {
	var hSorted Hand
	var index int

	for j := 0; j < len(h); j++ {
		for i := 0; i < len(h); i++ {
			if h[i].face > hSorted[j].face {
				hSorted[j] = h[i]
				index = i
			}
		}
		h[index].face = 0
	}

	return hSorted
}

func load() <-chan [2]Hand {
	c := make(chan [2]Hand)
	go func() {
		defer close(c)

		raw, _ := os.ReadFile("p054_poker.txt")
		lines := strings.Split(string(raw), "\n")

		for _, line := range lines {
			cards := strings.Split(line, " ")
			if len(cards) <= 1 {
				continue
			}
			var pair [2]Hand
			for i, card := range cards {
				pair[i/5][i%5].suit = rune(card[len(card)-1:][0])
				pair[i/5][i%5].face = faceToInt[card[:len(card)-1]]
			}
			pair[0] = sort(pair[0])
			pair[1] = sort(pair[1])
			c <- pair
		}
	}()

	return c
}

func main() {
	player1 := 0
	player2 := 0

	for hands := range load() {
		switch winner(hands[0], hands[1]) {
		case 1:
			player1++
		case 2:
			player2++
		default:
			fmt.Println("ERROR: unknown player")
		}
	}

	fmt.Printf("Player1: %d\nPlayer2: %d\n", player1, player2)
}
