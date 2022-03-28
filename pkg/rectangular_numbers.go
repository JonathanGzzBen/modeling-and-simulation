package pkg

import (
	"fmt"
	"math/rand"
)

// GetRectangularNumbers returns n rectangular numbers
func GetRectangularNumbers(n int) []float64 {
	fmt.Print("Quiere generar los numeros rectangulares automaticamente? [Y/n]: ")
	var input string
	fmt.Scanf("%s\n", &input)
	rectangularNumbers := make([]float64, n)
	if input != "n" {
		for i := 0; i < n; i++ {
			rectangularNumber := rand.Float64()
			rectangularNumbers = append(rectangularNumbers, rectangularNumber)
			fmt.Printf("%.5f\n", rectangularNumber)
		}
	} else {
		for i := 0; i < n; i++ {
			fmt.Print("Ingrese numero: ")
			var in float64
			fmt.Scanf("%f\n", &in)
			rectangularNumbers = append(rectangularNumbers, in)
		}
	}
	return rectangularNumbers
}
