package main

// go fmt ./... && go vet ./... && go test && go run 098.go -cpuprofile cpu.prof && echo top | go tool pprof cpu.prof

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"

	"github.com/erikbryant/util-golang/util"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
)

// findAnagrams returns any anagram pairs of the given words, bucketed by word
// length
func findAnagrams(wordsByLen [][]string) [][][]string {
	anagrams := make([][][]string, len(wordsByLen))

	for l, words := range wordsByLen {
		for i, word := range words {
			for j := i + 1; j < len(words); j++ {
				if util.IsAnagram(word, words[j]) {
					anagrams[l] = append(anagrams[l], []string{word, words[j]})
				}
			}
		}
	}

	return anagrams
}

// loadWords reads all of the words from the dictionary in and returns them
// bucketed by word length
func loadWords() [][]string {
	raw, _ := ioutil.ReadFile("p098_words.txt")
	csv := strings.ReplaceAll(string(raw), "\"", "")
	words := strings.Split(csv, ",")

	wordsByLen := [][]string{}

	for _, word := range words {
		l := len(word)
		for l >= len(wordsByLen) {
			wordsByLen = append(wordsByLen, []string{})
		}
		wordsByLen[l] = append(wordsByLen[l], word)
	}

	return wordsByLen
}

// crypt returns the string that is made when substitutions are applied to the given word
func crypt(word string, substitutions map[byte]byte) string {
	word2 := ""

	for i := 0; i < len(word); i++ {
		word2 += string(substitutions[word[i]])
	}

	return word2
}

// cryptable returns the second square (if it exists) of an anagram pair
func cryptable(square int, w1, w2 string) (int, bool) {
	sSquare := fmt.Sprintf("%d", square)

	// Does the given square pattern match to the first word?
	substitutions, ok := util.Cryptoquip(sSquare, w1)
	if ok {
		// If so, does it convert into the second word?
		sSquare2 := crypt(w2, substitutions)
		if sSquare2[0] == '0' {
			// Leading zeroes are not allowed
			return 0, false
		}
		square2, _ := strconv.Atoi(sSquare2)
		return square2, util.IsSquare(square2)
	}

	return 0, false
}

// match returns the max square (if any) where the given square matches any
// of the anagrams
func match(square int, anagrams [][][]string) int {
	sSquare := fmt.Sprintf("%d", square)
	maxSquare := 0

	for _, words := range anagrams[len(sSquare)] {

		if square2, ok := cryptable(square, words[0], words[1]); ok {
			fmt.Println("Match!", square, square2, words)
			if square > maxSquare {
				maxSquare = square
			}
			if square2 > maxSquare {
				maxSquare = square2
			}
		}

		if square2, ok := cryptable(square, words[1], words[0]); ok {
			fmt.Println("Match!", square, square2, words)
			if square > maxSquare {
				maxSquare = square
			}
			if square2 > maxSquare {
				maxSquare = square2
			}
		}
	}

	return maxSquare
}

// generateSquares generates each square number, checking to see if they satisfy
// the problem
func generateSquares(anagrams [][][]string) int {
	maxSquare := 0

	for i := 2; ; i++ {
		square := i * i
		sSquare := fmt.Sprintf("%d", square)
		if len(sSquare) >= len(anagrams) {
			break
		}
		max := match(square, anagrams)
		if max > maxSquare {
			maxSquare = max
		}
	}

	return maxSquare
}

func main() {
	fmt.Printf("Welcome to 098\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	wordsByLen := loadWords()
	anagrams := findAnagrams(wordsByLen)
	maxInt := generateSquares(anagrams)
	fmt.Println("The maximum is:", maxInt)
}
