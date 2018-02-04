package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	start := int(math.Sqrt(1020304050607080900.0))

	// The target number ends in '0', so its root
	// must also end in '0'. Only try squaring numbers
	// that are multiples of ten.
	start = start - start%10

	i := start
	for {
		n := i * i
		s := strconv.Itoa(n)
		if s[0] == '1' && s[2] == '2' && s[4] == '3' && s[6] == '4' && s[8] == '5' && s[10] == '6' && s[12] == '7' && s[14] == '8' && s[16] == '9' && s[18] == '0' {
			fmt.Println(i, s)
			break
		}

		i += 10
	}
}
