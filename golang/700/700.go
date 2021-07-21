package main

import (
	"fmt"
)

// Consider the sequence 1504170715041707*n mod 4503599627370517.
//
// An element of this sequence is defined to be an Eulercoin if it
// is strictly smaller than all previously found Eulercoins.
//
// For example, the first term is 1504170715041707 which is the first Eulercoin.
// The second term is 3008341430083414 which is greater than 1504170715041707 so
// is not an Eulercoin. However, the third term is 8912517754604 which is small
// enough to be a new Eulercoin.
//
// The sum of the first 2 Eulercoins is therefore 1513083232796311.
//
// Find the sum of all Eulercoins.
func eulerCoins(coin, mod int64) int64 {
	var sum int64
	var smallest int64
	var accumulator int64
	var i int64
	var lasti int64
	var biggestdelta int64

	sum = 0
	smallest = coin + 1
	accumulator = coin
	i = 0

	fmt.Println("            delta(i)                 skip                  sum             smallest")
	for {
		if accumulator < smallest {
			sum += accumulator
			smallest = accumulator

			if accumulator == 0 {
				break
			}

			// Each eulerCoin is at least as far away as the last. Fast forward
			// one short of what it took to find this coin.
			if i > 0 {
				deltai := (i - lasti) - 1
        lasti = i

				if deltai > biggestdelta {
					biggestdelta = deltai
				}

				fmt.Printf("%20d %20d %20d\n", deltai, sum, smallest)

				var ff int64
				for ff = 0; ff < biggestdelta/6000; ff++ {
					accumulator += coin * 6000
					accumulator %= mod
				}
				accumulator += coin * (biggestdelta % 6000)
				accumulator %= mod

				i += 6000*(biggestdelta/6000) + (biggestdelta % 6000)
			}
		}

		accumulator += coin

		if accumulator > mod {
			accumulator -= mod
		}

		i++
	}

	return sum
}

func main() {
	fmt.Println("Welcome to 700")

	fmt.Println("eulerCoins(1504170715041707, 4503599627370517) = ", eulerCoins(1504170715041707, 4503599627370517))
}
