package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	//Первый способ через context.WithCancel
	ctx, cancel := context.WithCancel(context.Background())
	go first(ctx)
	time.Sleep(time.Second * 4)
	cancel() //Посылаем сигнал об остановке

	//Второй способ через канал
	channel := make(chan bool)
	go second(channel)
	fmt.Println(<-channel)

	//Третий способ через runtime.Goexit()
	go third()
	time.Sleep(time.Second)
}

func first(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("I'm working")
		}
		time.Sleep(time.Second)
	}
}

func second(channel chan bool) {
	fmt.Println("I'm working")
	channel <- true
}

func third() {
	fmt.Println("I'm working")
	runtime.Goexit()
}
