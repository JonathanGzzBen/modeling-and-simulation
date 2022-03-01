package pkg

func addition(a int, x0 int, c int) int {
	return (a * x0) + c
}

func x1(a int, x0 int, c int, m int) int {
	return addition(a, x0, c) % m
}

func rectangularNumber(a int, x0 int, c int, m int) float64 {
	return float64(addition(a, x0, c)%m) / float64(m)
}

func MixedCongruentialGeneration(a int, x0 int, c int, m int) (int, int, float64) {
	addition := addition(a, x0, c)
	x1 := x1(a, x0, c, m)
	rectangularNumber := rectangularNumber(a, x0, c, m)
	return addition, x1, rectangularNumber
}
