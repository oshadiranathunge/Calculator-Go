package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter number:")
	fmt.Scanln(&num1)

	fmt.Print("Enter number:")
	fmt.Scanln(&num2)

	fmt.Print("Enter operator (+ - / * ):")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Printf("%f %s %f = %.3f", num1, number1, operator, number1+num2)
	case "-":
		fmt.Printf("%f %s %f = %.3f", num1, number1, operator, number1-num2)
	case "*":
		fmt.Printf("%f %s %f = %.3f", num1, number1, operator, number1*num2)

	case "/":
		if num2 == 0.0 {
			fmt.Println("Cannot divide by Zero")
		} else {
			fmt.Printf("%f %s %f = %.3f", num1, number1, operator, number1/num2)

		}
	default:
		fmt.Println("Please enter a valid operation")
	}
}
