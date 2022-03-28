package main

import (
	"fmt"
	"modeling-and-simulation/pkg"
	"os"
)

func main() {
	n := readIntOrExit("Ingrese cantidad de numeros: ", "valor de \"n\" no válido")
	sum := 0.0

	rectangularNumbers := pkg.GetRectangularNumbers(n)
	for _, rn := range rectangularNumbers {
		sum += rn
	}

	avg := sum / float64(n)
	fmt.Print("Ingrese α: ")
	var alpha float64
	_, err := fmt.Scanf("%f\n", &alpha)
	if err != nil {
		_ = fmt.Errorf("α invalido\n")
		return
	}
	if pkg.AverageStatistical(avg, alpha) {
		fmt.Println("Los numeros son aceptados")
	} else {
		fmt.Println("Los numeros son rechazados")
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
