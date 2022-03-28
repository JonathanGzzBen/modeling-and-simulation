package pkg

import (
	"fmt"
	"math"
)

func z0(average float64) float64 {
	numerator := (average - (0.5)) * math.Pow(10, 0.5)
	z0 := math.Abs(numerator / math.Pow(1.0/12.0, 0.5))
	// Set precision to 5 decimal places
	z0 = math.Round(z0*100000) / 100000
	fmt.Printf("|((%v - 1/2) * √10) / (√1/12)| = %v\n", average, z0)
	return z0
}

func zAlphaHalf(alpha float64) float64 {
	alphaReal := 100.0 - (100.0 * alpha)
	alphaHalf := (alphaReal / 100.0) / 2.0
	fmt.Printf("α real = 100 - %v = %v%%\n", alpha*100.0, alphaReal)
	fmt.Printf("Zα/2 = Z.%v/2 = Z%v\n", alphaReal, alphaHalf)
	return alphaHalf
}

func isAccepted(z0 float64, zAlphaHalf float64) bool {
	fmt.Println("Z0 < Zα/2")
	fmt.Printf("%v < %v\n", z0, zAlphaHalf)
	return z0 < zAlphaHalf
}

func AverageStatistical(average float64, alpha float64) bool {
	z0 := z0(average)
	zAlphaHalf := zAlphaHalf(alpha)
	var zAlphaHalfValue float64
	fmt.Printf("Ingrese valor de Z%v: ", zAlphaHalf)
	_, _ = fmt.Scanf("%f\n", &zAlphaHalfValue)
	fmt.Printf("Zα/2 = %v\n", zAlphaHalfValue)
	return isAccepted(z0, zAlphaHalfValue)
}
