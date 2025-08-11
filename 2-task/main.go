package main

import "fmt"

func Calculator(a, b float64, op string) float64 {

	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "/":
		if b != 0 {
			return a / b
		}
		fmt.Println("Cannot devide by 0")
		return 0
	case "*":
		return a * b
	default:
		fmt.Println("invalid charactor")
		return 0

	}

}

func main() {

	fmt.Println("Whelcome to calculator ")
	var a, b float64
	var operator string
	fmt.Println("Enter first number")
	fmt.Scan(&a)
	fmt.Println("Enter the operator(+ ,- ,/ , *)")
	fmt.Scan(&operator)
	fmt.Println("Enter second number")
	fmt.Scan(&b)
	result := Calculator(a, b, operator)
	fmt.Println("Result :", result)
}
