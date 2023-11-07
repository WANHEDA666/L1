package main

import "fmt"

func main() {
	showType(67)
	showType("dfdsf")
	showType([]int{1, 2, 3})
}

func showType(some interface{}) {
	fmt.Printf("%T", some) //выводим полученный тип через fmt
}
