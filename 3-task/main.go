package main

import "fmt"

func Fibinocii(n int) {

	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()

}

func main() {
	var n int
	fmt.Println("Enter how many fibbinoci number to print")
	fmt.Scan(&n)
	Fibinocii(n)
}
