package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Second * 4
	channel := make(chan int)

	go func() {
		for true {
			channel <- 1 //Передаем данные в канал
			time.Sleep(time.Second)
		}
	}()

	go func() {
		time.Sleep(n) //По истечению n секунд закрываем канал
		close(channel)
	}()

	for item := range channel {
		fmt.Println(item) //Чистаем данные из канала пока он не закрыт
	}
}
