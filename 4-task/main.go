package main

import (
	"errors"
	"fmt"
)

func Add(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot deviceible by zero")
	}
	return a + b, nil
}

func main() {
	var a, b int
	fmt.Println("Enter 2 interger numbers ")
	fmt.Scan(&a, &b)
	result, err := Add(a, b)
	if err != nil {
		fmt.Println("Error :", err)
	}
	fmt.Println("Result :", result)
}
