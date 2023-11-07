package main

import (
	"fmt"
	"time"
)

func main() {
	waitForSeconds(5)
	fmt.Println("I'm free")
}

func waitForSeconds(seconds int) {
	start := time.Now()
	end := start.Add(time.Duration(seconds * 1000000000))
	for time.Now().Before(end) { //Сравниваем стартовое значение с финальным
	}
}
