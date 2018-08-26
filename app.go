package main

import (
	"fmt"
	"os"
	
)

func main() {
	args := os.Args
	annualSalary := args[1]

	fmt.Println("Hello"+annualSalary)
}