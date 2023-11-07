package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	substrings := strings.Split(str, " ") //Создаем массив слов по разделению пробелами
	for i := 0; i < len(substrings)/2; i++ {
		substrings[i], substrings[len(substrings)-1-i] = substrings[len(substrings)-1-i], substrings[i] //Переставляем данные местами
	}
	concatenatedString := strings.Join(substrings, " ") //Пересобираем строку
	fmt.Println(concatenatedString)
}
