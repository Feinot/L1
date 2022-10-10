package main

import (
	"context"
	"fmt"
	"time"
)

func work(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // если пришел сигнал об отмене контекста
			fmt.Println("Exit")
			return
		default:
			fmt.Println("working...")
		}
	}

}

func main() {
	ctx, cancel := context.WithCancel(context.Background()) // создаем экземляр контекста
	go work(ctx)
	time.Sleep(2 * time.Second)
	cancel() // отмена контекста
	time.Sleep(2 * time.Second)

}
