package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	sum := 0

	wg := new(sync.WaitGroup) // определяем группу горутин
	mu := new(sync.Mutex)     // определяем мьютекс

	for _, val := range arr {
		wg.Add(1) // увеличиваем счетчик горутин в группе
		go func(num int) {
			defer wg.Done() // сообщаем, что элемент группы завершил работу
			mu.Lock()
			sum += num * num
			mu.Unlock()
		}(val)
	}
	wg.Wait() // ожидаем завершения всех горутин в группе

	fmt.Println(sum)
}
