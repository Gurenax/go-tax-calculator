package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
*	Function to convert string to float64
 */
func convertStringToFloat64(str string) float64 {
	if result, err := strconv.ParseFloat(str, 64); err == nil {
		return result
	}
	return 0
}

/*
* Function to determine the tax bracket
 */
func determineTaxBracke(salary float64) string {
	if salary <= 18200.0 {
		return "Tax Bracket 1"
	} else if salary > 18200.0 && salary <= 37000.0 {
		return "Tax Bracket 2"
	} else if salary > 37000.0 && salary <= 90000.0 {
		return "Tax Bracket 3"
	} else if salary > 90000.0 && salary <= 180000.0 {
		return "Tax Bracket 4"
	} else if salary > 180000.0 {
		return "Tax Bracket 5"
	} else {
		return "None"
	}
}

func main() {
	// Get command line arguments
	args := os.Args

	// Use first argument as annual salary and convert to float64
	annualSalary := convertStringToFloat64(args[1])

	// Output annual salary
	fmt.Println("The annual salary is: $" + fmt.Sprintf("%f", annualSalary))

	// Output tax bracket
	fmt.Println("The tax bracket is: " + determineTaxBracke(annualSalary))
}
