package main

// go fmt ./... && go vet ./... && go test && go build heegner.go && time ./heegner

// A secret bonus problem! :)
//
// https://projecteuler.net/problem=heegner
//
// Among all non-square integers n with absolute value not exceeding 10^3,
// find the value of n such that cos(π*sqrt(n)) is closest to an integer.

import (
	"fmt"
	"math"
	"math/big"
)

const (
	Precision = 2000
	Pi        = "3.141592653589793238462643383279502884197169399375105820974944592307816406286208998628034825342117067982148086513282306647093844609550582231725359408128481117450284102701938521105559644622948954930381964428810975665933446128475648233786783165271201909145648566923460348610454326648213393607260249141273724587006606315588174881520920962829254091715364367892590360011330530548820466521384146951941511609433057270365759591953092186117381932611793105118548074462379962749567351885752724891227938183011949129833673362440656643086021394946395224737190702179860943702770539217176293176752384674818467669405132000568127145263560827785771342757789609173637178721468440901224953430146549585371050792279689258923542019956112129021960864034418159813629774771309960518707211349999998372978049951059731732816096318595024459455346908302642522308253344685035261931188171010003137838752886587533208381420617177669147303598253490428755468731159562863882353787593751957781857780532171226806613001927876611195909216420199"
)

func coshFract(n int) float64 {
	//           ∞   x^(2k)
	// cosh(x) = Σ   ------
	//          k=0  (2k)!
	//
	//           ∞                        x^(2k)
	// cosh(x) = Σ   T(k)   where T(k) =  ------
	//          k=0                       (2k)!
	//
	//                                x^2
	// T(0) = 1  and  T(k+1) = T(k) --------
	//                              2k(2k-1)

	x := new(big.Float).SetPrec(Precision).SetFloat64(float64(n))
	x.Sqrt(x)
	pi := new(big.Float).SetPrec(Precision).SetFloat64(0.0)
	pi.SetString(Pi)
	x.Mul(x, pi)
	tk := new(big.Float).SetPrec(Precision).SetFloat64(1.0)
	numerator := new(big.Float).SetPrec(Precision).SetFloat64(1.0)
	numerator.Mul(x, x)
	cosh := new(big.Float).SetPrec(Precision).SetFloat64(1.0)
	for k := 1; tk.MantExp(nil) > -1000; k++ {
		denominator := new(big.Float).SetPrec(Precision).SetFloat64(float64(2 * k * ((2 * k) - 1)))
		tk.Mul(tk, numerator)
		tk.Quo(tk, denominator)
		cosh.Add(cosh, tk)
	}

	// Remove the integer portion of cosh, leaving just the decimal portion
	i := big.NewInt(0)
	cosh.Int(i)
	str := i.Text(10)
	decimal := new(big.Float).SetPrec(Precision).SetFloat64(0.0)
	decimal.SetString(str)
	cosh.Sub(cosh, decimal)

	// Convert the decimal portion to a float64 and return it
	f64, _ := cosh.Float64()

	return f64
}

// closestNPos returns n where 0 <= n < 10^3 and cos(π*sqrt(n)) is closest to an integer
func closestNPos() (int, float64) {
	// No need to calculate the cosine. Just find values of sqrt(n)
	// that are integers or integers+0.5. We don't care about the
	// decimal portion, just the fractional. So, look for the
	// fractional portion of sqrt(n) that is closest to 0.0 or 0.5.

	n := 0
	closestZero := 0
	closestHalf := 0
	distZero := 999.9
	distHalf := 999.9

	for n = 0; n < 1000; n++ {
		sqrt := math.Sqrt(float64(n))
		if sqrt == math.Trunc(sqrt) {
			// Skip square values of n
			continue
		}
		fractional := sqrt - math.Trunc(sqrt)
		if fractional < distZero {
			distZero = fractional
			closestZero = n
		}
		if fractional < distHalf {
			distHalf = fractional
			closestHalf = n
		}
	}

	if distZero < distHalf {
		return closestZero, distZero
	}
	return closestHalf, distHalf
}

func closestNNeg() (int, float64) {
	minFrac := 999.0
	minN := 999
	maxFrac := -1.0
	maxN := -1

	for n := 1; n < 1000; n++ {
		cf := coshFract(n)
		if cf < minFrac {
			minFrac = cf
			minN = n
		}
		if cf > maxFrac {
			maxFrac = cf
			maxN = n
		}
	}

	// Is maxFrac closer to 1.0 or is minFrac closer to 0.0?
	if 1.0-maxFrac < minFrac {
		return -maxN, maxFrac
	}
	return -minN, minFrac
}

func main() {
	fmt.Printf("Welcome to heegner\n\n")

	nPos, nPosDist := closestNPos()
	fmt.Printf("Closest +n: %4d  at: %.15f\n", nPos, nPosDist)

	nNeg, nNegDist := closestNNeg()
	fmt.Printf("Closest -n: %4d  at: %.15f\n", nNeg, nNegDist)
}
