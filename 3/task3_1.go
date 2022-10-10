package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	sum := 0

	wg := new(sync.WaitGroup)
	ch := make(chan int, len(arr)) // буферизованный канал для синхронизации данных

	for _, num := range arr {
		wg.Add(1)
		go func(val int) {
			ch <- val * val
			wg.Done()
		}(num)
	}
	wg.Wait()
	close(ch)

	for val := range ch {
		sum += val
	}
	fmt.Println(sum)

}
