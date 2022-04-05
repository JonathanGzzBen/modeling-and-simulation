package main

import (
	"fmt"
	"math"
	"modeling-and-simulation/pkg"
)

func main() {
	var N int
	fmt.Print("Ingrese N:")
	fmt.Scanf("%d\n", &N)
	var n int
	fmt.Print("Ingrese n:")
	fmt.Scanf("%d\n", &n)

	FEi := float64(N) / float64(n)
	interval := 1.0 / float64(n)
	fmt.Printf("FEi = N/n = %2d/%-2d = %f\n", N, n, FEi)

	rectangularNumbers := pkg.GetRectangularNumbers(N)
	sum := 0.0
	fmt.Printf("X20 = 1 / %f[", FEi)
	for i := 0.0; i < 1; i += interval {
		count := 0
		for _, rn := range rectangularNumbers {
			if i <= rn && rn < (i+interval) {
				count++
			}
		}
		fmt.Printf("(%d - %f)^2 +", count, FEi)
		sum += math.Pow(float64(count)-FEi, 2)
	}
	fmt.Printf("\b\b]\n")
	x20 := (1.0 / FEi) * sum
	fmt.Printf("X20 = %f\n", x20)

	fmt.Printf("Ingrese el valor de X^2 alpha, n - 1: ")
	var xalphan float64
	fmt.Scanf("%f\n", &xalphan)
	fmt.Println("X20 < X2 alpha, n -1")
	fmt.Printf("%f < %f\n", x20, xalphan)
	if x20 < xalphan {
		fmt.Println("Los numeros son aceptados")
	} else {
		fmt.Println("Los numeros son rechazados")
	}
}
