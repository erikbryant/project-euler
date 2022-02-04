package digits_413

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"runtime/pprof"
	"strconv"
)

//
// If len(n) > 9 then there are no single-digit sequences that
// are divisible by len(n) unless there is a zero in n.
//

// We say that a d-digit positive number (no leading zeros)
// is a one-child number if exactly one of its sub-strings
// is divisible by d.

// fastForwardZero() skips past combinations that are known to have multiple children.
func fastForwardZero(digits []int, length int) {
	// Increment past duplicate zeros.
	// E.G., 903,022 -> 903,100 -> 903,110 -> 903,111
	// There cannot be a zero in digits[0], so skip that.
	i := 1
	for ; i < length; i++ {
		if digits[i] == 0 {
			break
		}
	}
	i++
	for ; i < length; i++ {
		if digits[i] == 0 {
			for ; i < length; i++ {
				digits[i] = 1
			}
		}
	}
}

// fastForwardIndex() skips over combinations that are known to have multiple children.
func fastForwardIndex(digits []int, length, index int) bool {
	// Increment past a second child.
	// E.G.,    4435 -> 4500 -> 4510 -> 4511
	//  Index ---^
	for i := index + 1; i < length; i++ {
		digits[i] = 1
	}

	return increment(digits, index+1)
}

func increment(digits []int, length int) bool {
	l := length
	zero := false

	for l > 0 {
		l--
		if digits[l] == 9 {
			digits[l] = 0
			if l == 0 {
				// Integer overflow.
				return false
			}
			zero = true
			continue
		}
		digits[l]++
		break
	}

	if zero {
		fastForwardZero(digits, length)
	}

	return true
}

var Shifts = []int{
	0,                        // Substrings are 1-based, so skip zero.
	int(math.Pow(10.0, 0.0)), // substring len == 1
	int(math.Pow(10.0, 1.0)),
	int(math.Pow(10.0, 2.0)),
	int(math.Pow(10.0, 3.0)),
	int(math.Pow(10.0, 4.0)),
	int(math.Pow(10.0, 5.0)),
	int(math.Pow(10.0, 6.0)),
	int(math.Pow(10.0, 7.0)),
	int(math.Pow(10.0, 8.0)),
	int(math.Pow(10.0, 9.0)),
	int(math.Pow(10.0, 10.0)),
	int(math.Pow(10.0, 11.0)),
	int(math.Pow(10.0, 12.0)),
	int(math.Pow(10.0, 13.0)),
	int(math.Pow(10.0, 14.0)),
	int(math.Pow(10.0, 15.0)),
	int(math.Pow(10.0, 16.0)),
	int(math.Pow(10.0, 17.0)),
	int(math.Pow(10.0, 18.0)), // substring len == 19
}

// countChildren() counts the number of children for this slen substring length.
// Returns index==0 if 1 or fewer children are found.
// Returns index==X if a second child is found, where X is the end position of the second child.
func countChildrenOdd(digits []int, divisor, end, slen int) (count int, index int) {
	// Build the first string of digits 'slen' long.
	shift := Shifts[slen]
	substring := digits[0]
	for j := 1; j < slen; j++ {
		substring = substring*10 + digits[j]
	}

	if substring%divisor == 0 {
		count++
	}

	// Add/remove digits to walk the string to the end of 'digits'.
	for i := 1; i <= end-slen+1; i++ {
		substring = (substring-digits[i-1]*shift)*10 + digits[i+slen-1]
		if substring%divisor == 0 {
			count++
			if count > 1 {
				index = i + slen - 1
				break
			}
		}
	}

	return
}

// oneChildManual() returns the number of one-child numbers that have 'length' number of digits.
// It does it using brute force, except for the ones it knows how to compute.
func oneChildManualSingleDigit(length, divisor int) int {
	switch divisor {
	case 1:
		// Due to the loop unrolling below, we need to exclude 1-digit
		// numbers from the evaluation (they would get double counted).
		return 9
	case 2:
		return 4 * int(math.Pow(5.0, float64(length-1)))
	case 3:
		return 360
	case 4:
		return 2701
	case 5:
		return int(math.Pow(8.0, float64(length-1)))
	case 6:
		return 109466
	case 7:
		return 161022
	case 8:
		return 13068583
	case 9:
		return 2136960
	}

	// Start at 1x10^length.
	digits := make([]int, length)
	digits[0] = 1
	fastForwardZero(digits, length)

	count := 0
	next := false
	index := 0
	for {
		index = 0
		oneChildCount := 0

		// Enumerate all single-digit substrings.
		for j := 0; j < length; j++ {
			if digits[j]%divisor == 0 {
				oneChildCount++
				if oneChildCount > 1 {
					index = j
					next = true
					goto NEXT
				}
			}
		}

		// Enumerate all remaining substrings and check for divisbility.
		for slen := 2; slen <= length; slen++ {
			children := 0
			children, index = countChildrenOdd(digits, divisor, length-1, slen)
			oneChildCount += children
			if oneChildCount > 1 {
				next = true
				goto NEXT
			}
		}

		if oneChildCount == 1 {
			count++
		}

	NEXT:
		if next && index != 0 {
			if !fastForwardIndex(digits, length, index) {
				break
			}
			next = false
		} else {
			if !increment(digits, length) {
				break
			}
		}
	}

	return count
}

func router(length, divisor int) int {
	if divisor < 10 {
		return oneChildManualSingleDigit(length, divisor)
	}

	switch divisor {
	// case 11:
	// 	return 71101800
	// case 12:
	// 	return 55121700430
	// case 13:
	// 	return 1057516028
	default:
		// zeroP = z(length-1,divisor)
		// zero = zeroP * 10      // previous zero-childs now appended with [0..9]

		// zero -= Y              // zero-childs that are now multi-child
		// zero -= X              // zero-childs that are now single-child

		// countP = f(length-1,divisor)
		// count = countP * 9    // previous one-childs now appended with [1..9]

		// count -= Z             // one-childs that are now multi-child
		// count += X             // zero-childs that are now single-child
	}

	if divisor%2 == 0 {
		return oneChildManualEven(length, divisor)
	}
	return oneChildManualOdd(length, divisor)
}

// oneChildManual() returns the number of one-child numbers that have 'length' number of digits.
// It does it using brute force, except for the ones it knows how to compute.
func oneChildManualEven(length, divisor int) int {

	// Start at 1x10^length.
	digits := make([]int, length)
	digits[0] = 1
	fastForwardZero(digits, length)

	count := 0
	for {
	TOP:
		oneChildCount := 0

		// Only look at substrings that end with an even digit.
		for i := 0; i < length; i++ {
			if digits[i]&0x01 != 0 {
				continue
			}
			sum := 0
			shift := 1
			for j := i; j >= 0; j-- {
				// Cannot have a multi-digit substring with a leading zero.
				if digits[j] == 0 && j != i {
					shift *= 10
					continue
				}
				sum += digits[j] * shift
				if sum%divisor == 0 {
					oneChildCount++
					if oneChildCount > 1 {
						// TODO: Figure out why this is broken.
						// if !fastForwardIndex(digits, length, i) {
						if !increment(digits, length) {
							goto DONE
						}
						goto TOP
					}
				}
				shift *= 10
			}
		}

		if oneChildCount == 1 {
			count++
		}

		if !increment(digits, length) {
			break
		}
	}

DONE:
	return count
}

// oneChildManual() returns the number of one-child numbers that have 'length' number of digits.
// It does it using brute force, except for the ones it knows how to compute.
func oneChildManualOdd(length, divisor int) int {

	// Start at 1x10^length.
	digits := make([]int, length)
	digits[0] = 1
	fastForwardZero(digits, length)

	count := 0
	for {
	TOP:
		oneChildCount := 0

		// Single-digit substrings. Since divisor > 9 we only care about zero.
		// There is at most one zero. The first digit cannot be a zero.
		for j := 1; j < length; j++ {
			if digits[j] == 0 {
				oneChildCount++
				break
			}
		}

		// Enumerate all remaining substrings and check for divisbility.
		for slen := 2; slen <= length; slen++ {
			children, index := countChildrenOdd(digits, divisor, length-1, slen)
			oneChildCount += children
			if oneChildCount > 1 {
				if index > 0 {
					if !fastForwardIndex(digits, length, index) {
						goto DONE
					}
					goto TOP
				}
				break
			}
		}

		if oneChildCount == 1 {
			count++
		}

		if !increment(digits, length) {
			break
		}
	}

DONE:
	return count
}

// F(N) is the number of the one-child numbers less than N.
func F(n int) int {
	count := 0
	exponent := 1

	for n > 1 {
		children := router(exponent, exponent)
		count += children
		exponent++
		n = n / 10
	}

	return count
}

func strIncrement(s string) (string, bool) {
	l := len(s)
	zero := false

	for l > 0 {
		l--
		if s[l] == 9+48 {
			s = s[:l] + "0" + s[l+1:]
			if l == 0 {
				// Integer overflow.
				return "", false
			}
			zero = true
			continue
		}
		s = s[:l] + string(s[l]+1) + s[l+1:]
		break
	}

	if zero {
		s = strFastForwardZero(s)
	}

	return s, true
}

// fastForwardIndex() skips over combinations that are known to have multiple children.
func strFastForwardIndex(s string, index int) (string, bool) {
	// Increment past a second child.
	// E.G.,    4435 -> 4500 -> 4510 -> 4511
	//  Index ---^
	ok := true
	length := len(s)
	s = s[:index+1]
	s, ok = strIncrement(s)
	for i := index + 1; i < length; i++ {
		s += "0"
	}

	return s, ok
}

// fastForwardZero() skips over combinations that are known to have multiple children.
func strFastForwardZero(s string) string {
	// Increment past duplicate zeros.
	// E.G., 903,022 -> 903,100 -> 903,110 -> 903,111
	// There cannot be a zero in s[0], so skip that.
	length := len(s)
	i := 1
	for ; i < length; i++ {
		if s[i] == 0+48 {
			break
		}
	}
	i++
	for ; i < length; i++ {
		if s[i] == 0+48 {
			s = s[:i]
			for ; i < length; i++ {
				s += "1"
			}
		}
	}

	return s
}

func strCompare(n string, pos int, t string) bool {
	j := 1

	for ; j < len(t) && n[pos+j] == t[j]; j++ {
	}

	return j == len(t)
}

func findSubstringByRank(n string, multiples [10][20][]string) (oneChildCount int, index int) {
	maxLength := len(n)
	for i := 0; i < len(n); i++ {
		firstDigit := n[i] - 48

		// For each multiple that begins with 'firstDigit', see if it is in 'n'.
		// Iterate through each length that would fit.
		for l := 1; l <= maxLength; l++ {
			for _, target := range multiples[firstDigit][l] {
				if strCompare(n, i, target) {
					idx := i + l - 1
					if idx > index {
						index = idx
					}
					oneChildCount++
					if oneChildCount > 1 {
						return
					}
				}
			}
		}

		maxLength--
	}

	return
}

func strMultiplesByRank(divisor int) [10][20][]string {
	// For very fast lookup we need to be able to rank the multiples
	// by what their first digit is and how many digits are in the
	// multiple. Create a 3D array indexed by first digit and length.
	// There are no zero-length multiples, so those slices will be empty.
	// [firstDigit][length][]string
	var multiples [10][20][]string

	i := 0
	sum := 0
	dMax := int(math.Pow(10.0, float64(divisor))) - 1
	shortcut := shortcuts(divisor)
	for sum <= dMax && i < shortcut {
		s := strconv.Itoa(sum)
		firstDigit := s[0] - 48
		length := len(s)
		multiples[firstDigit][length] = append(multiples[firstDigit][length], s)
		sum += divisor
		i++
	}

	return multiples
}

func strChildrenByRank(divisor int, length int) int {
	min := int(math.Pow(10.0, float64(length-1)))

	multiples := strMultiplesByRank(divisor)

	s := strconv.Itoa(min)
	s = strFastForwardZero(s)

	count := 0
	ok := true
	for {
		found, index := findSubstringByRank(s, multiples)
		if found > 1 && index < len(s)-1 {
			s, ok = strFastForwardIndex(s, index)
			if !ok {
				break
			}
		} else {
			if found == 1 {
				count++
			}
			s, ok = strIncrement(s)
			if !ok {
				break
			}
		}
	}

	return count
}

// shortcuts() returns the maximum # of multiples that need to be
// calculated for a given digit length (divisor). It turns out that
// for most divisors only the first few multiples matter in identifying
// whether a number is one-child or not.
// Some divisors have more complicated patterns than others. For these,
// an upper bounds approximation is returned.
func shortcuts(d int) int {
	shortcut := 0

	switch d {
	case 1:
		shortcut = 10
	case 2:
		shortcut = 50
	case 3:
		shortcut = 270
	case 4:
		shortcut = 250
	case 5:
		shortcut = 20
	case 6:
		// 165@3 1650@4 16500@5 165000@6
		shortcut = 165 * int(math.Pow(10.0, float64(d-3)))
	case 7:
		// 143@3 1429@4 14286@5  ????
		shortcut = 143 * int(math.Pow(10.0, float64(d-3)))
	case 8:
		shortcut = 1250
	case 9:
		// 99@3 999@4 9988@5 99876@6  ????
		shortcut = 998 * int(math.Pow(10.0, float64(d-3)))
	case 10:
		shortcut = 10
	case 11:
		// 90@3 900@4 9000@5
		shortcut = 9 * int(math.Pow(10.0, float64(d-2)))
	case 12:
		// 84@3 825@4 8250@5
		shortcut = 825 * int(math.Pow(10.0, float64(d-4)))
	case 13:
		// 77@3 770@4 7693@5  ????
		shortcut = 77 * int(math.Pow(10.0, float64(d-3)))
	case 14:
		// 72@3 714@4 7135@5  ????
		shortcut = 72 * int(math.Pow(10.0, float64(d-3)))
	case 15:
		// 66@3 660@4 6600@5
		shortcut = 66 * int(math.Pow(10.0, float64(d-3)))
	case 16:
		// 63@3 625@4 6250@5
		shortcut = 625 * int(math.Pow(10.0, float64(d-4)))
	case 17:
		// 59@3 589@4 5883@5  ????
		shortcut = 59 * int(math.Pow(10.0, float64(d-3)))
	case 18:
		// 55@3 550@4 5500@5
		shortcut = 55 * int(math.Pow(10.0, float64(d-3)))
	case 19:
		// 53@3 527@4 5264@5  ????
		shortcut = 53 * int(math.Pow(10.0, float64(d-3)))
	default:
		fmt.Println("ERROR: Unsupported value passed to shortcuts()", d)
	}

	return shortcut
}

func strF(f int) {
	count := 0
	for i := 1; i <= f; i++ {
		children := strChildrenByRank(i, i)
		count += children
		fmt.Println(i, children, count)
	}
	fmt.Println("F(", f, ") =", count)
}
