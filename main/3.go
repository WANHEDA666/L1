package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	result := 0
	firstChannel := make(chan int)

	var wg sync.WaitGroup
	var lock sync.Mutex

	for _, item := range arr {
		wg.Add(1)
		item := item
		go func() {
			lock.Lock() //Блокируем запись в канал
			firstChannel <- item
			wg.Done()
			lock.Unlock() //Разблокируем запись в канал
		}()
	}

	go func() { //Асинхронно ждем выполнения всех горунтин и потом закрываем канал
		wg.Wait()
		close(firstChannel)
	}()

	for item := range firstChannel { //Читаем данные из канала пока он не закрыт
		result += item * item
	}

	fmt.Println(result)
}
