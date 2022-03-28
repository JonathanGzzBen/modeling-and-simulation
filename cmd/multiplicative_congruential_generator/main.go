package main

import (
	"fmt"
	"modeling-and-simulation/pkg"
)

func main() {
	a := pkg.ReadIntOrExit("Enter a: ", "invalid \"a\" value")
	x0 := pkg.ReadIntOrExit("Enter x0: ", "invalid \"x0\" value")
	m := pkg.ReadIntOrExit("Enter m: ", "invalid \"m\" value")

	expectedPeriod := pkg.ExpectedPeriod(m)
	fmt.Printf("Periodo esperado: %v\n", expectedPeriod)
	x0Aux := x0
	fmt.Println("\n n\tx0\t(ax0)modm\txn+1\tNúmeros rectangulares")
	var n int
	for n = 1; n <= expectedPeriod; n++ {
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
