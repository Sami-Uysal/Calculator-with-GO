package main

import (
	"fmt"
)

func main() {
	//Calculator
	var a, b int
	var operator string
	fmt.Println("Enter the first number: ")
	fmt.Scan(&a)
	fmt.Println("Enter the second number: ")
	fmt.Scan(&b)
	fmt.Println("Enter the operator: ")
	fmt.Scan(&operator)
	switch operator {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		fmt.Println(a / b)
	default:
		fmt.Println("Invalid operator")

	}

}
