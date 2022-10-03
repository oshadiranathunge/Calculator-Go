package main

import "fmt"

func main() {
	var num1, number2 float64
	var operator string

	fmt.Print("Enter number:")
	fmt.Scanln(&num1)

	fmt.Print("Enter number:")
	fmt.Scanln(&number2)

	fmt.Print("Enter operator (+ - / * ):")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Printf("%f %s %f = %.3f", num1, number1, operator, number1+number2)
	case "-":
		fmt.Printf("%f %s %f = %.3f", num1, number1, operator, number1-number2)
	case "*":
		fmt.Printf("%f %s %f = %.3f", num1, number1, operator, number1*number2)

	case "/":
		if number2 == 0.0 {
			fmt.Println("Divide by zero situation")
		} else {
			fmt.Printf("%f %s %f = %.3f", num1, number1, operator, number1/number2)

		}
	default:
		fmt.Println("Please enter a valid operation")
	}
}
