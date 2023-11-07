package main

import (
	"fmt"
	"sort"
)

func main() {
	willingToFind := 78
	arr := []int{1, 56, 32, 78, 5, 9, 678, 123, 4, 8, 88, 93}
	sort.Ints(arr) //Сортируем массив
	fmt.Println(findId(arr, 0, len(arr)-1, willingToFind))
}

func findId(arr []int, left int, right int, toFind int) int {
	if left == right { //Если остался один элемент, значит это искомый
		return left
	}
	center := left + (right-left)/2 //Выбираем середину
	if arr[center] == toFind {      //Смотрим не попали ли мы случайно на искомый элемент
		return center
	}
	if arr[center] > toFind { //Если середина больше искомого, то ищем искомое в левом массиве
		return findId(arr, left, center-1, toFind)
	} else { //Если середина меньше искомого, то ищем искомое в правом массиве
		return findId(arr, center+1, right, toFind)
	}
}
