package main

import (
	"fmt"
	"modeling-and-simulation/pkg"
	"sort"
)

func main() {
	n := pkg.ReadIntOrExit("Ingrese cantidad de numeros: ", "valor de \"n\" no válido")
	rectangularNumbers := pkg.GetRectangularNumbers(n)
	sort.Float64s(rectangularNumbers)
	fmt.Println("Numeros rectangulares ordenados de forma ascendente.")
	for _, rn := range rectangularNumbers {
		fmt.Println(rn)
	}
	highestDn := -1.0
	fmt.Printf("%-4v%-10v%-12v%-v\n", "i", "xi", "F(xi)", "Dn = max |F(xi) - xi|")
	for i, rn := range rectangularNumbers {
		xi := rn
		fxi := float64(i+1) / float64(n)
		dn := fxi - xi
		if dn > highestDn {
			highestDn = dn
		}
		fmt.Printf("%-4v%-10.5f%2v/%v=%-6v%3v - %.5f = %-10.5f\n", i+1, xi, i+1, n, fxi, fxi, xi, dn)
	}
	fmt.Printf("Dn = %-.5f\n", highestDn)
	var tableStatistical float64
	fmt.Print("Ingrese el estadistico de tablas: ")
	_, _ = fmt.Scanf("%f\n", &tableStatistical)
	fmt.Println("Dn < x,N")
	fmt.Printf("%-.5f < %-.5f", highestDn, tableStatistical)
	if highestDn < tableStatistical {
		fmt.Println("Los números son aceptados")
	} else {
		fmt.Println("Los números son rechazados")
	}
}
