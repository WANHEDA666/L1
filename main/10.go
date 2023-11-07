package main

import (
	"math"
	"strconv"
	"strings"
)

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	result := make(map[string][]float64)
	for _, item := range arr {
		stringItem := strconv.FormatFloat(item, 'f', 2, 32) //Переводим float в string
		firstNumber := ""
		if item < 0 {
			firstNumber = stringItem[0:2] + calculateZeros(item) //Определяем число для ключа
		} else {
			firstNumber = stringItem[0:1] + calculateZeros(item) //Определяем число для ключа
		}
		result[firstNumber] = append(result[firstNumber], item)
	}
}

func calculateZeros(number float64) string {
	intModule := int(math.Round(math.Abs(number)))      //Отбрасываем дровную часть и делаем число положительным
	intString := strconv.Itoa(intModule)                //Переаодим int в string
	return buildZerosString(len([]rune(intString)) - 1) //Определяем длину строки и убираем один элемент
}

func buildZerosString(number int) string {
	var builder strings.Builder
	for i := 0; i < number; i++ {
		builder.WriteString("0") //Собираем количество нулей после первого элемента
	}
	return builder.String()
}
