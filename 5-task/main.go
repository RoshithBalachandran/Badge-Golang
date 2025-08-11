package main

import (
	"fmt"
	"math"

)

func ISprime(n int) bool {
	if n<=1{
		return false
	}
	limit:=int(math.Sqrt(float64(n)))
	for i:=2;i<limit;i++{
		if n%i==0{
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Println("Enter number to find is it ptime or not")
	fmt.Scan(&n)
	ISprime(n)

	if ISprime(n){
		fmt.Println("Prime ",n)
	}else{
		fmt.Println("Not prime ",n)
	}
}