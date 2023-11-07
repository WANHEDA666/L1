package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, item := range arr { //Запускам пять горунтин
		wg.Add(1) //Добавляем счетчик ожидания
		item := item
		go func() {
			fmt.Println(item * 2) //Горунтины выполняются асинхронно, кто первый выполнится - тот первый выведет значение и тд
			wg.Done()             //Убавляем счетчик ожидания
		}()
	}

	wg.Wait()
}
