package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println(count) //1000
}
