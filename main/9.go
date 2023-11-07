package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	firstChannel := make(chan int)
	secondChannel := make(chan int)

	go func() {
		for _, item := range arr {
			firstChannel <- item //Передаем данные в канал и останавливаем выполнение до их прочтения
		}
		close(firstChannel)
	}()

	go func() {
		for item := range firstChannel {
			secondChannel <- item * 2 //Чистаем данные из канала
		}
		close(secondChannel)
	}()

	for item := range secondChannel { //Читаем данные из канала пока он не закрыт
		fmt.Println(item)
	}
}
