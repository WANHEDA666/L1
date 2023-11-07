package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "abCdefAaf"
	fmt.Println(checkForNotDuplicate(str))
}

func checkForNotDuplicate(str string) bool {
	lowercaseString := strings.ToLower(str) //Понижаем весь регистр до нижнего
	checkMap := make(map[string]bool)
	for _, item := range lowercaseString {
		if _, ok := checkMap[string(item)]; !ok { //Проаеряем есть ли такой элемент в словаре
			checkMap[string(item)] = true
		} else {
			return false //Если уже есть в словаре, значит дубликат, возвращаем false
		}
	}
	return true
}
