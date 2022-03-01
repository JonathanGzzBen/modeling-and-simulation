package main

import (
	"fmt"
	"math"
	"modeling-and-simulation/pkg"
	"os"
)

func main() {
	a := readIntOrExit("Enter a: ", "invalid \"a\" value")
	x0 := readIntOrExit("Enter x0: ", "invalid \"x0\" value")
	m := readIntOrExit("Enter m: ", "invalid \"m\" value")

	var expectedPeriod int
	// If number is base 2
	if math.Log2(float64(m)) == math.Floor(float64(m)) {
		expectedPeriod = m / 4
	} else if math.Log10(float64(m)) == math.Floor(float64(m)) {
		expectedPeriod = m
	}
	x0Aux := x0
	fmt.Println("\n n\tx0\t(ax0)modm\txn+1\tNúmeros rectangulares")
	var n int
	for n = 1; n <= m; n++ {
		multiplication, x1, rectangularNumber := pkg.MultiplicativeCongruentialGeneration(a, x0Aux, 0, m)
		intPart := multiplication / m
		residueNumerator := multiplication - (intPart * m)
		fmt.Printf("%2d\t%2d\t%2d + %2d/%-2d\t%3d\t%10.5f\n", n, x0Aux, intPart, residueNumerator, m, x1, rectangularNumber)
		if x1 == x0 {
			break
		}
		x0Aux = x1
	}
	if n == expectedPeriod {
		fmt.Println("Generador multiplicativo confiable")
	} else {
		fmt.Println("Generador multiplicativo no confiable")
	}

}
func readIntOrExit(prompt string, errorMessage string) int {
	fmt.Print(prompt)
	var value int
	_, err := fmt.Scanf("%d\n", &value)
	if err != nil {
		_ = fmt.Errorf(errorMessage)
		os.Exit(1)
	}
	return value
}
