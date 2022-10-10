package main

import (
	"fmt"
	"sync"
)

//Каналы имеют встроенный механизм синхронизации — операции вставки и извлечения выполняются последовательно.

//Обеспечив отправку "задач" через канал с единственным получателем, мы "естественным" образом выполним увеличение счетчика последовательно.

func main() {
	counter := 0
	n := 1000
	ch := make(chan struct{}, n)
	chanWg := sync.WaitGroup{}
	chanWg.Add(1)
	go func() {
		for range ch {
			counter++
		}
		chanWg.Done()
	}()

	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			ch <- struct{}{}
			wg.Done()
		}(i)
	}
	wg.Wait()
	close(ch)
	chanWg.Wait()

	fmt.Println(counter) //1000
}
