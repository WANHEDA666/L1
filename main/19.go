package main

import (
	"fmt"
)

func main() {
	str := "главрыба"
	arr := []rune(str) //Переносим строку в массив рун
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i] //Переставляем данные местами
	}
	result := string(arr)
	fmt.Println(result)
}
