package moessner

// Moessner addition-only factorial
//
// http://www.luschny.de/math/factorial/FastFactorialFunctions.htm
// http://www.luschny.de/math/factorial/csharp/FactorialAdditiveMoessner.cs.html

// Factorial returns n!
func Factorial(n int) int {
	// Note that this function does not remove trailing zeroes.

	if n <= 1 {
		return 1
	}

	s := make([]int, n+1)
	s[0] = 1

	for m := 1; m <= n; m++ {
		s[m] = 0
		for k := m; k >= 1; k-- {
			for i := 1; i <= k; i++ {
				s[i] += s[i-1]
			}
		}
	}

	return s[n]
}
