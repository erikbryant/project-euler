package main

// Preformatting of the input file:
// sed 's/,/\n/g' 042_words.txt | sed 's/"//g' > 042_words.txt

import (
	"./triangles"
	"fmt"
	"io/ioutil"
	"strings"
)

func triangleWord(s string) bool {
	if len(s) < 1 {
		return false
	}

	sum := 0
	for i := 0; i < len(s); i++ {
		sum += int(s[i]) - int('A') + 1
	}
	return triangles.Triangle(sum)
}

func countWords() {
	raw, _ := ioutil.ReadFile("042_words.txt")
	lines := strings.Split(string(raw), string(10))
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
	triangles.Init()
	countWords()
}
