package main

import (
	"fmt"
	"modeling-and-simulation/pkg"
	"os"
)

func main() {
	m := readIntOrExit("Enter m: ", "invalid \"m\" value")
	expectedPeriod := pkg.ExpectedPeriod(m)
	fmt.Printf("Expected period: %v\n", expectedPeriod)
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
