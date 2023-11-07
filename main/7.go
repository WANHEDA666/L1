package main

import (
	"fmt"
	"sync"
)

func main() {
	newMap := make(map[int]bool)
	var wg sync.WaitGroup
	var lock sync.Mutex

	for i := 0; i < 3; i++ {
		i := i
		wg.Add(1)
		go func() {
			lock.Lock() //Блокируем доступ к словарю
			newMap[i] = true
			lock.Unlock() //Разблокируем доступ к словарю
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(newMap)
}
