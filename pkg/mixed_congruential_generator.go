package pkg

func Addition(a int, x0 int, c int) int {
	return (a * x0) + c
}

func X1(a int, x0 int, c int, m int) int {
	return Addition(a, x0, c) % m
}

func RectangularNumber(a int, x0 int, c int, m int) float64 {
	return float64(Addition(a, x0, c)%m) / float64(m)
}

// MixedCongruentialGeneration generates a new rectangular number.
func MixedCongruentialGeneration(a int, x0 int, c int, m int) (addition int, x1 int, rectangularNumber float64) {
	return Addition(a, x0, c), X1(a, x0, c, m), RectangularNumber(a, x0, c, m)
}
