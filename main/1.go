package main

import "fmt"

func main() {
	action := new(Action)
	action.Print()
}

type Human struct {
}

func (human *Human) Print() { //Определяем метод Print() и выводим 2
	fmt.Println(2)
}

type Action struct {
	Human
}

func (action *Action) Print() { //Переопределяем метод Print() и выводим 4
	fmt.Println(4)
}
