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

/*
* Function to compute salary
 */
func computeSalary(salary float64, num1 float64, num2 float64, num3 float64) float64 {
	return (salary-num1)*num2 + num3
}

/*
*	Function to calculate the annual tax
 */
func calculateAnnualTax(salary float64) float64 {
	switch determineTaxBracke(salary) {
	case "Tax Bracket 1":
		return 0.0
	case "Tax Bracket 2":
		return computeSalary(salary, 18200.0, 0.19, 0.0)
	case "Tax Bracket 3":
		return computeSalary(salary, 37000.0, 0.325, 3572.0)
	case "Tax Bracket 4":
		return computeSalary(salary, 90000.0, 0.37, 20797.0)
	case "Tax Bracket 5":
		return computeSalary(salary, 180000.0, 0.45, 54097.0)
	default:
		return 0.0
	}
}

/*
* Function to calculate annual taxed income
 */
func calculateAnnualTaxedIncome(annualSalary float64, annualTax float64) float64 {
	return annualSalary - annualTax
}

/*
* Function to calculate income less medical levy
 */
func calculateIncomeLessMedicalLevy(taxedSalary float64, annulSalary float64) float64 {
	medicalLevyPercentage := 0.02
	return taxedSalary - (annulSalary * medicalLevyPercentage)
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

	// Calculate the annual tax
	annualTax := calculateAnnualTax(annualSalary)
	fmt.Println("The annual tax is: $" + fmt.Sprintf("%f", annualTax))

	// Calculate the annual taxed income
	annualTaxedIncome := calculateAnnualTaxedIncome(annualSalary, annualTax)
	fmt.Println("The annual taxed income is: $" + fmt.Sprintf("%f", annualTaxedIncome))

	// Calculate income less medical levy
	fmt.Println("The annual taxed income less medical levy(2%) is: $" + fmt.Sprintf("%f", calculateIncomeLessMedicalLevy(annualTaxedIncome, annualSalary)))
}
