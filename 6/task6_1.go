package main

import (
	"fmt"
	"time"
)

func main() {
	quit := make(chan bool)
	go func() {
		fmt.Println("Start")
		for {
			select {
			case <-quit: // если пришли данные, завершаем работу
				fmt.Println("Stop")
				return
			}
		}
	}()
	quit <- true // отправляем данные в канал
	time.Sleep(1 * time.Second)
}
