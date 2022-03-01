package pkg

func Multiplication(a int, x0 int, _ int) int {
	return a * x0
}

// MultiplicativeCongruentialGeneration generates a new rectangular number.
func MultiplicativeCongruentialGeneration(a int, x0 int, c int, m int) (addition int, x1 int, rectangularNumber float64) {
	return Multiplication(a, x0, 0), X1(a, x0, c, m, Multiplication), RectangularNumber(a, x0, c, m, Multiplication)
}
