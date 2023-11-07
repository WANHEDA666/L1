package main

import (
	"fmt"
	"math"
)

func main() {
	firstPoint := NewPoint(1, 2)
	secondPoint := NewPoint(3, 4)
	xLen := math.Abs(firstPoint.x - secondPoint.x) //Вычисляем расстояние между X
	yLen := math.Abs(firstPoint.y - secondPoint.y) //Вычисляем расстояние между Y
	fmt.Println(xLen, yLen)
}

type Point struct { //Структура
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point { //Конструктор структуры
	return &Point{
		x: x,
		y: y,
	}
}
