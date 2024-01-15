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
	"math/big"
	"strings"

	"github.com/erikbryant/util-golang/algebra"
	"github.com/erikbryant/util-golang/romanNumerals"
)

// possibles returns a map of all valid subsequent chars for the given prefix
func possibles(prefix string, romans []string) map[string]bool {
	nextChars := map[string]bool{}
	i := len(prefix)

	for _, roman := range romans {
		if len(roman) > i && strings.HasPrefix(roman, prefix) {
			nextChars[string(roman[i])] = true
		}
	}

	return nextChars
}

// possiblesToString returns a string (in sorted order) of the roman numerals in possibles
func possiblesToString(possibles map[string]bool) string {
	s := ""
	for _, r := range []string{"I", "V", "X", "L", "C", "D", "M", "#"} {
		if possibles[r] {
			s += r
		}
	}
	return s
}

func probabilityBigInt(roman string, romans []string) (*big.Int, *big.Int) {
	pNumer := big.NewInt(1)
	pDenom := big.NewInt(1)

	for i := 0; i < len(roman); i++ {
		prefix := roman[0:i]
		poss := possibles(prefix, romans)
		terminal := 0
		l := len(poss)
		myWeight := 14
		if roman[i] == '#' {
			myWeight = 2
		}
		if poss["#"] {
			terminal = 2
			l -= 1
		}
		w := big.NewInt(int64(myWeight))
		t := big.NewInt(int64(14*l + terminal))
		pNumer, pDenom = algebra.MulFractionBigInt(pNumer, pDenom, w, t)
		// fmt.Printf("rn: %10s char: %c possibles: %7s weight: %d / %d ret: %d/%d\n", roman, roman[i], possiblesToString(poss), myWeight, 14*l+terminal, pNumer, pDenom)
	}

	return pNumer, pDenom
}

func firstThousandBigInt() (*big.Int, *big.Int) {
	romans := []string{}
	evCumN := big.NewInt(0)
	evCumD := big.NewInt(1)

	for i := 0; i <= 999; i++ {
		r := romanNumerals.Roman(i) + "#"
		romans = append(romans, r)
	}

	for i := 0; i <= 999; i++ {
		n, evD := probabilityBigInt(romans[i], romans)
		temp := big.NewInt(int64(i))
		evN := temp.Mul(n, temp)
		evCumN, evCumD = algebra.SumFractionBigInt(evCumN, evCumD, evN, evD) // Running total of expected value
		// fmt.Printf("rn: %14s  this: %18d/%-18d   EV: %18d/%-18d   cum: %20d/%-20d\n", romans[i], n, d, evN, evD, evCumN, evCumD)
	}

	return evCumN, evCumD
}

func main() {
	fmt.Printf("Welcome to 610\n\n")

	// Calculate the expected value of 0-999
	n, d := firstThousandBigInt()
	evWhole := new(big.Int)
	evFractional := new(big.Int)
	evWhole.QuoRem(n, d, evFractional)
	n2, d2 := algebra.ReduceFractionBigInt(evFractional, d)
	fmt.Printf("EV: %d / %d == %d r %d == %d + %d / %d\n", n, d, evWhole, evFractional, evWhole, n2, d2)

	// Convert the EV we found to a BigFloat. Patch the
	// decimal and fractional portions together in a
	// string as the initializer. Only way I could find
	// to preserve this much precision.
	floatAsString := fmt.Sprintf("%d", evWhole)
	evFractional.Mul(evFractional, big.NewInt(1000000000000000000))
	evWhole.QuoRem(evFractional, d, evFractional)
	floatAsString = fmt.Sprintf("%s.%d", floatAsString, evWhole)
	ev0to999 := new(big.Float)
	ev0to999.SetString(floatAsString)
	fmt.Printf("%d + %d / %d == %0.15f\n", evWhole, n2, d2, ev0to999)

	// Now we add in the M numbers. Keep adding until nothing
	// up to and including the 9th decimal place changes.
	oneMoreM := big.NewFloat(1000.0)
	mProb := big.NewFloat(0.14)

	fmt.Printf("\nSearching for stability beyond the 9th decimal place...\n\n")
	for i := 1; i < 25; i++ {
		oneMoreM.Mul(oneMoreM, mProb)
		ev0to999.Add(ev0to999, oneMoreM)
		fmt.Printf("M's: %3d EV: %0.15f\n", i, ev0to999)
	}
}
