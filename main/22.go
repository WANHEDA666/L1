package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	x := big.NewInt(int64(math.Pow(2, 40))) //Используем библиотеку big для работы с большими числами
	y := big.NewInt(int64(math.Pow(2, 38)))
	fmt.Println(multiply(x, y))
	fmt.Println(divide(x, y))
	fmt.Println(add(x, y))
	fmt.Println(subtract(x, y))
}

func multiply(x *big.Int, y *big.Int) *big.Int { //Умножение
	return x.Mul(x, y)
}

func divide(x *big.Int, y *big.Int) *big.Int { //Деление
	return x.Div(x, y)
}

func add(x *big.Int, y *big.Int) *big.Int { //Сложение
	return x.Add(x, y)
}

func subtract(x *big.Int, y *big.Int) *big.Int { //Вычитание
	return x.Sub(x, y)
}
