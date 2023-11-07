package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

//Используем context.WithCancel т.к. им легко управлять из главной горунтины

func main() {
	workers := 4
	channel := make(chan int) //Канал, куда будем записывать данные
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background()) //Контекст для завершения горунтины
	wg.Add(1)
	go func() {
		for true {
			channel <- rand.Int() //Передаем данные в канал
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < workers; i++ { //Запускаем читателей
		go channelReader(channel, ctx)
	}

	c := make(chan os.Signal, 1) //Канал проверки не нажали ли мы Ctrl+C
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			cancel() //При нажатии Ctrl+C все останавливаем
			wg.Done()
			close(channel)
			close(c)
		}
	}()

	wg.Wait()
}

func channelReader(channel <-chan int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(<-channel)
		}
	}
}
