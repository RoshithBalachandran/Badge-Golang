package main

import "fmt"

func InsertEnd(arr []int, val int) []int {
	return append(arr, val)
}

func InsertPosition(arr []int, pos int, val int) []int {
	if pos < 0 || pos > len(arr) {
		return arr
	}
	arr = append(arr, 0)
	copy(arr[pos+1:], arr[:pos])
	arr[pos] = val
	return arr
}

func DeleteEnd(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	return arr[:len(arr)-1]
}

func DeleteAt(arr []int,pos int)[]int{
	if len(arr)<1 ||pos >=len(arr){
		return arr
	}
	copy(arr[pos:],arr[pos+1:])
	return arr[:len(arr)-1]
}
func main() {
	var arr = []int{5, 9, 7, 6, 4, 3}
	Result := InsertEnd(arr, 25)
	fmt.Println("Result InsertEnd :", Result)

	Result = InsertPosition(arr, 3, 8)
	fmt.Println("Result InsertPosition:", Result)

	Result = DeleteEnd(arr)
	fmt.Println("Result DeleteEnd:", Result)

	Result = DeleteAt(arr ,2)
	fmt.Println("Result DeleteAt:", Result)
}
