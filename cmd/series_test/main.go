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

	FEi := (float64(N) - 1) / math.Pow(float64(n), 2)
	fmt.Printf("FEi = N-1 / n^2 = %2d -1 /%2d^2 = %f\n", N, n, FEi)

	type RectangularNumbersPair struct {
		x float64
		y float64
	}

	rectangularNumbersX, rectangularNumbersY := pkg.GetRectangularNumbers(N-1), pkg.GetRectangularNumbers(N-1)
	var rectPairs []RectangularNumbersPair
	for i := 0; i < N-1; i++ {
		rectPairs = append(rectPairs, RectangularNumbersPair{x: rectangularNumbersX[i], y: rectangularNumbersY[i]})
	}
	var cells [][]int
	interval := 1.0 / float64(n)
	for xInterval := 0.0; xInterval < 1; xInterval += interval {
		var xValues []int
		for yInterval := 0.0; yInterval < 1; yInterval += interval {
			sum := 0
			for _, pair := range rectPairs {
				if xInterval <= pair.x && pair.x < (xInterval+interval) && yInterval <= pair.y && pair.y < (yInterval+interval) {
					sum++
				}
			}
			xValues = append(xValues, sum)
		}
		cells = append(cells, xValues)
	}

	for x := 0; x < n; x++ {
		fmt.Print("[")
		for y := 0; y < n; y++ {
			fmt.Printf("%-2d", cells[x][y])
		}
		fmt.Println("]")
	}

	x20 := math.Pow(float64(n), 2) / (float64(N) - 1.0)
	fmt.Println("x0^2 = n^2/N - 1")
	fmt.Printf("x0^2 = %d/%d ", n*n, N-1)
	sumToMultiply := 0.0
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			sumToMultiply += math.Pow(float64(cells[x][y])-FEi, 2)
			if x == 0 && y == 0 {
				fmt.Printf("(%d - %f)^2 + [", cells[x][y], FEi)
				continue
			}
			fmt.Printf("(%d - %f)^2", cells[x][y], FEi)
			if x+1 <= n && y+1 <= n {
				fmt.Print(" + ")
			}
		}
	}
	x20 *= sumToMultiply

	fmt.Println("]")
	fmt.Printf("x0^2 = %v\n", x20)

	var alpha float64
	fmt.Print("Ingrese alpha: ")
	fmt.Scanf("%f\n", &alpha)
	fmt.Printf("Xalpha^2, n^2 - 1 = %f, %f\n", alpha, float64((n*n)-1))

	var tableStatistical float64
	fmt.Print("Ingrese el estadistico de tabla: ")
	fmt.Scanf("%f\n", &tableStatistical)

	fmt.Printf("Xalpha^2, n^2 - 1 = %f, %f = %f\n", alpha, float64((n*n)-1), tableStatistical)

	fmt.Println("Xalpha^2, n^2 - 1 > X0^2")
	fmt.Printf("%f > %f\n", tableStatistical, x20)
	if tableStatistical > x20 {
		fmt.Println("Los numeros son aceptados")
	} else {
		fmt.Println("Los numeros son rechazados")
	}
}
