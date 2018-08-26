package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
*	Convert string to float64
*/
func convertStringToFloat64(str string) float64 {
	if result, err := strconv.ParseFloat(str, 64); err == nil {
		return result
	}
	return 0
}

func main() {
	// Get command line arguments
	args := os.Args

	// Use first argument as annual salary and convert to float64
	annualSalary := convertStringToFloat64(args[1])

	// Output annual salary
	fmt.Println("The annual salary is: $"+fmt.Sprintf("%f", annualSalary))
}