package main

// Preformatting of the input file:
// sed 's/,/\n/g' 042_words.txt | sed 's/"//g' > 042_words.txt

import (
	"fmt"
	"os"
	"strings"

	"github.com/erikbryant/util-golang/figurate"
)

func triangleWord(s string) bool {
	if len(s) < 1 {
		return false
	}

	sum := 0
	for i := 0; i < len(s); i++ {
		sum += int(s[i]) - int('A') + 1
	}
	return figurate.IsTriangular(sum)
}

func countWords() {
	raw, _ := os.ReadFile("042_words.txt")
	lines := strings.Split(string(raw), "\n")
	count := 0
	for _, line := range lines {
		if triangleWord(line) {
			fmt.Println(line)
			count++
		}
	}
	fmt.Println("Count: ", count)
}

func main() {
	countWords()
}
