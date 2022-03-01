package pkg

import "math"

func ExpectedPeriod(m int) int {
	var expectedPeriod int
	if (m & (m - 1)) == 0 {
		// If number is base 2
		expectedPeriod = m / 4
	} else if math.Log10(float64(m)) == math.Floor(float64(m)) {
		// If number is base 10
		exponent := math.Log10(float64(m))
		if exponent >= 5 {
			expectedPeriod = int(5 * math.Pow(10, exponent-2))
		} else {
			expectedPeriod = int((math.Pow(5, exponent-1)) * 4)
		}
		expectedPeriod = m
	}
	return expectedPeriod
}

func Multiplication(a int, x0 int, _ int) int {
	return a * x0
}

// MultiplicativeCongruentialGeneration generates a new rectangular number.
func MultiplicativeCongruentialGeneration(a int, x0 int, c int, m int) (multiplication int, x1 int, rectangularNumber float64) {
	return Multiplication(a, x0, 0), X1(a, x0, c, m, Multiplication), RectangularNumber(a, x0, c, m, Multiplication)
}
