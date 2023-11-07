package main

import "fmt"

//Чтобы не хранить в памяти весь массив v, как в изначальном коде, мы делим получаемую строку на руны, а потом копируем в новый
//массив столько рун, сколько нам надо и забываем о первоначальном массиве

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	result := make([]rune, 10)
	copy(result, v[:10])
	justString = string(result)
	fmt.Println(justString)
}

func createHugeString(number int) []rune {
	return []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()")
}

func main() {
	someFunc()
}
