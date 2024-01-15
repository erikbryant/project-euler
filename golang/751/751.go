package main

// go fmt ./... && go vet ./... && go test && go run 751.go

// A non-decreasing sequence of integers a(n) can be generated from any positive
// real value θ by the following procedure:
//
//   b1 = θ
//   bn = ⌊bn-1⌋ * (bn-1 - ⌊bn-1⌋ + 1) for all n >= 2
//   an = ⌊bn⌋
//
// Where ⌊.⌋ is the floor function.
//
// For example, θ = 2.956938891377988... generates the Fibonacci sequence:
// 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
//
// The concatenation of a sequence of positive integers a(n) is a real value denoted 𝛕
// constructed by concatenating the elements of the sequence after the decimal point,
// starting at a1: a1.a2a3a4...
//
// For example, the Fibonacci sequence constructed from θ = 2.956938891377988...
// yields the concatenation 𝛕 = 2.3581321345589... Clearly, 𝛕 != θ for this value of θ.
//
// Find the only value of θ for which the generated sequence starts at a1 = 2 and the
// concatenation of the generated sequence equals the original value: 𝛕 = θ.
// Give your answer rounded to 24 places after the decimal point.

import (
	"fmt"
	"math/big"
)

// floor returns the whole portion of the given big.Float (equivalent to trunc)
func floor(f *big.Float) *big.Float {
	whole := big.NewInt(0)
	f.Int(whole)
	return big.NewFloat(float64(whole.Int64()))
}

// stringify returns the string representation of the non-fractional portion of the given big.Float
func stringify(f *big.Float) string {
	whole := big.NewInt(0)
	f.Int(whole)
	return fmt.Sprintf("%d", whole.Int64())
}

// calc returns b[n] = a[n-1] * (b[n-1] - a[n-1] + 1) when calld with a==a[n-1] and b==b[n-1]
func calc(a, b *big.Float) *big.Float {
	tempA := big.NewFloat(0)
	tempB := big.NewFloat(0)

	one := big.NewFloat(1.0)

	tempA.Copy(a)
	tempB.Copy(b)

	tempB.Sub(tempB, tempA)
	tempB.Add(tempB, one)
	tempB.Mul(tempB, tempA)

	return tempB
}

// tauBigFloat64 returns the tau generated from the given theta
func tauBigFloat64(theta *big.Float) string {
	limit := 30
	a := make([]*big.Float, limit+1)
	b := make([]*big.Float, limit+1)

	n := 1
	b[n] = theta
	a[n] = floor(b[n])
	n++

	tauStr := stringify(a[1]) + "."

	for len(tauStr) < limit {
		// b[n] = a[n-1] * (b[n-1] - a[n-1] + 1)
		b[n] = calc(a[n-1], b[n-1])
		a[n] = floor(b[n])
		tauStr += stringify(a[n])
		n++
	}

	return tauStr
}

// convergeBigFloat64 finds the theta that generats an equal tau
func convergeBigFloat64(t float64) {
	var ok bool
	theta := big.NewFloat(t)
	theta.SetPrec(300)

	for {
		tauStr := tauBigFloat64(theta)
		thetaStr := fmt.Sprintf("%0.30f", theta)
		fmt.Printf("theta: %s\n  tau: %s\n", thetaStr[:27], tauStr[:27])
		if tauStr[:26] == thetaStr[:26] {
			fmt.Println("Found the answer! ----^")
			break
		}
		theta, ok = theta.SetString(tauStr)
		if !ok {
			panic("Import of string failed")
		}
	}
}

func main() {
	fmt.Printf("Welcome to 751\n\n")

	theta := 2.0
	convergeBigFloat64(theta)
}
