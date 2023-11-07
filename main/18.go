package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := new(Counter)

	var wg sync.WaitGroup
	var lock sync.Mutex

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			lock.Lock() //Блокируем запись в канал
			counter.Increment()
			wg.Done()
			lock.Unlock() //Разблокируем запись в канал
		}()
	}

	wg.Wait()

	fmt.Println(counter.i)
}

type Counter struct {
	i int
}

func (counter *Counter) Increment() {
	counter.i++
}
