package main

import "modeling-and-simulation/pkg"

func main() {
	n := pkg.ReadIntOrExit("Ingrese cantidad de numeros: ", "valor de \"n\" no válido")
	sum := 0.0

	rectangularNumbers := pkg.GetRectangularNumbers(n)
	for _, rn := range rectangularNumbers {
		sum += rn
	}

}
