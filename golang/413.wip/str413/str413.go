package str413

import (
	"fmt"
	"math"
	"strconv"
)

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
	shortcut := Shortcuts(divisor)
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

// Shortcuts returns the maximum # of multiples that need to be
// calculated for a given digit length (divisor). It turns out that
// for most divisors only the first few multiples matter in identifying
// whether a number is one-child or not.
// Some divisors have more complicated patterns than others. For these,
// an upper bounds approximation is returned.
func Shortcuts(d int) int {
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
