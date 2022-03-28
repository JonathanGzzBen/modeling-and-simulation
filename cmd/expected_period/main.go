package main

import (
	"fmt"
	"modeling-and-simulation/pkg"
)

func main() {
	m := pkg.ReadIntOrExit("Enter m: ", "invalid \"m\" value")
	expectedPeriod := pkg.ExpectedPeriod(m)
	fmt.Printf("Expected period: %v\n", expectedPeriod)
}
