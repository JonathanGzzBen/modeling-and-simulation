package pkg

func Addition(a int, x0 int, c int) int {
	return (a * x0) + c
}

// MixedCongruentialGeneration generates a new rectangular number.
func MixedCongruentialGeneration(a int, x0 int, c int, m int) (addition int, x1 int, rectangularNumber float64) {
	return Addition(a, x0, c), X1(a, x0, c, m, Addition), RectangularNumber(a, x0, c, m, Addition)
}
