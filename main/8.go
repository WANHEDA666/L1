package main

import "fmt"

func main() {
	i := 18
	arr := int64ToBits(int64(16))
	arr[i] = 1 //меняем i элемент
	fmt.Println(bitsToInt64(arr))
}

func int64ToBits(number int64) []int { //Переводим int64 в набор битов
	var bits []int
	for i := 0; i < 64; i++ {
		bit := (number >> i) & 1
		bits = append(bits, int(bit))
	}
	return bits
}

func bitsToInt64(bits []int) int64 { //Переводим набор битов в int64
	var result int64
	for i, bit := range bits {
		if bit == 1 {
			result |= 1 << i
		}
	}
	return result
}
