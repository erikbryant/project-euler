package main

// go fmt ./... && go vet ./... && go test && go run 610.go

// A random generator produces a sequence of symbols drawn from the set
// {I, V, X, L, C, D, M, #}. Each item in the sequence is determined by
// selecting one of these symbols at random, independently of the other
// items in the sequence. At each step, the seven letters are equally
// likely to be selected, with probability 14% each, but the # symbol
// only has a 2% chance of selection.
//
// We write down the sequence of letters from left to right as they are
// generated, and we stop at the first occurrence of the # symbol (without
// writing it). However, we stipulate that what we have written down must
// always (when non-empty) be a valid Roman numeral representation in
// minimal form. If appending the next letter would contravene this then
// we simply skip it and try again with the next symbol generated.
//
// Please take careful note of About... Roman Numerals for the definitive
// rules for this problem on what constitutes a "valid Roman numeral
// representation" and "minimal form". For example, the (only) sequence
// that represents 49 is XLIX. The subtractive combination IL is invalid
// because of rule (ii), while XXXXIX is valid but not minimal. The rules
// do not place any restriction on the number of occurrences of M, so all
// positive integers have a valid representation. These are the same rules
// as were used in Problem 89, and members are invited to solve that problem
// first.
//
// Find the expected value of the number represented by what we have written
// down when we stop. (If nothing is written down then count that as zero.)
// Give your answer rounded to 8 places after the decimal point.

//
// There is a pattern, of course, to the roman numerals. They run from
// '' (zero) to 'CMXCIX' (999) in one pattern. They then move to a highly
// repetitive pattern for numbers above that.
//
// For numbers 1000-1999 simply prepend an 'M' to each of the 0-999 values.
// For numbers 2000-2999 prepend 'MM'
// For numbers 3000-3999 prepend 'MMM'
//   ...
//
// This becomes an infinite sum of the expected values. It is the expected
// value of all roman numerals from 0-999 plus the expected value of all
// roman numerals from 1000-1999 plus those of 2000-2999 + ...
//
// Assume E(r) is the expected value of a given roman numeral r.
//
//    ∞                   999
//    ∑ P('M')^n (n*1000 + ∑ E(k))
//   n=0                  k=0
//
// Refactoring:
//
//         999
//    EV =  ∑ E(k)
//         k=0
//    ∞
//    ∑ P('M')^n (n*1000 + EV)
//   n=0
//
// The weighted values are 14% for each of 'IVXLCDM' and 2% for '#'.
// When generating the roman numerals from 0-999 the first character
// can be anything but 'M'. Subsequent characters depend highly on
// previous characters. The set must be enumerated to calculate the
// expected value.
//
// Roman numerals 1000-1999 must begin with 'M'. A 14% chance.
// Roman numerals 2000-2999 must begin with 'MM'. A 14%*14% chance.
//   ...
//

import (
	"fmt"
	"strings"

	"github.com/erikbryant/util-golang/algebra"
	"github.com/erikbryant/util-golang/romanNumerals"
)

// possibles returns a map of all valid subsequent chars for the given prefix
func possibles(prefix string, romans []string) map[byte]bool {
	nextChars := map[byte]bool{}
	i := len(prefix)

	for _, roman := range romans {
		if len(roman) > i && strings.HasPrefix(roman, prefix) {
			nextChars[roman[i]] = true
		}
	}

	return nextChars
}

func probability(roman string, romans []string) (int, int) {
	if roman == "" {
		// Out of (I, V, X, L, C, D, and #) we only care
		// about #. It is 2 units, the others are 14 units.
		poss := possibles("", romans)
		l := len(poss)
		PNumer := 2
		PDenom := 14*l + 2
		return PNumer, PDenom
	}

	// Probability roman starts with this letter
	poss := possibles("", romans)
	l := len(poss)
	PNumer := 14 * l
	PDenom := 14*l + 2

	prefix := ""
	for i := 0; i < len(roman)-1; i++ {
		// Probability of each middle letter
		prefix += string(roman[i])
		poss = possibles(prefix, romans)
		l = len(poss)
		PNumer, PDenom = algebra.MulFraction(PNumer, PDenom, 14*l, 14*l+2)
	}

	// Probability that roman ended at that final letter
	if len(roman) > 0 {
		poss = possibles(roman, romans)
		l = len(poss)
		PNumer, PDenom = algebra.MulFraction(PNumer, PDenom, 2, 14*l+2)
	}

	return PNumer, PDenom
}

func firstThousand() float64 {
	romans := []string{}
	expected := 0.0

	maxLen := 0
	for i := 0; i <= 999; i++ {
		r := romanNumerals.Roman(i)
		if len(r) > maxLen {
			maxLen = len(r)
		}
		romans = append(romans, r)
	}

	for i := 0; i <= 999; i++ {
		a, b := probability(romans[i], romans)
		fmt.Println("Finish writing this!", a, b)
	}

	return expected
}

func main() {
	fmt.Printf("Welcome to 610\n\n")
	expected := firstThousand()
	fmt.Println("Expected value of 0-999 =", expected)
}
