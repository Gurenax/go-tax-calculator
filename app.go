package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
* Function to convert string to float64
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
func determineTaxBracket(salary float64) string {
	if salary <= 0.0 {
		return "None"
	} else if salary <= 18200.0 {
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
* Function to compute tax
 */
func computeTax(salary float64, num1 float64, num2 float64, num3 float64) float64 {
	return (salary-num1)*num2 + num3
}

/*
* Function to calculate the annual tax
 */
func calculateAnnualTax(salary float64) float64 {
	switch determineTaxBracket(salary) {
	case "Tax Bracket 1":
		return 0.0
	case "Tax Bracket 2":
		return computeTax(salary, 18201.0, 0.19, 0.0)
	case "Tax Bracket 3":
		return computeTax(salary, 45001.0, 0.325, 5092.0)
	case "Tax Bracket 4":
		return computeTax(salary, 120001.0, 0.37, 29467.0)
	case "Tax Bracket 5":
		return computeTax(salary, 180001.0, 0.45, 51667.0)
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

/*
* Function to calculate monthly net earnings
 */
func calculateMonthlyNetEarnings(taxedSalary float64) float64 {
	return taxedSalary / 12.0
}

/*
* Convert float to string
 */
func floatToString(value float64) string {
	return fmt.Sprintf("%f", value)
}

/*
* Execute App
 */
func main() {
	// Get command line arguments
	args := os.Args

	// Use first argument as annual salary and convert to float64
	annualSalary := convertStringToFloat64(args[1])

	// Output annual salary
	fmt.Println("The annual salary is: $" + floatToString(annualSalary))

	// Output tax bracket
	fmt.Println("The tax bracket is: " + determineTaxBracket(annualSalary))

	// Calculate the annual tax
	annualTax := calculateAnnualTax(annualSalary)
	fmt.Println("The annual tax is: $" + floatToString(annualTax))

	// Calculate the annual taxed income
	annualTaxedIncome := calculateAnnualTaxedIncome(annualSalary, annualTax)
	fmt.Println("The annual taxed income is: $" + floatToString(annualTaxedIncome))

	// Calculate income less medical levy
	incomeLessMedicalLevy := calculateIncomeLessMedicalLevy(annualTaxedIncome, annualSalary)
	fmt.Println("The annual taxed income less medical levy(2%) is: $" + floatToString(incomeLessMedicalLevy))

	// Calculate monthly net earnings
	monthlyNetEarnings := calculateMonthlyNetEarnings(incomeLessMedicalLevy)
	fmt.Println("The monthly net earnings is: $" + floatToString(monthlyNetEarnings))
}
