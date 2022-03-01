package pkg

type generationOperation func(int, int, int) int

func X1(a int, x0 int, c int, m int, operation generationOperation) int {
	return operation(a, x0, c) % m
}

func RectangularNumber(a int, x0 int, c int, m int, operation generationOperation) float64 {
	return float64(operation(a, x0, c)%m) / float64(m)
}
