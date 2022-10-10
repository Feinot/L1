package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные
из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.
*/

func printData(ch chan int, wg *sync.WaitGroup) {
	for val := range ch {
		_, err := fmt.Fprintln(os.Stdout, val) // выводим данные из канала
		if err != nil {
			return
		}
	}
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	var N int
	fmt.Scan(&N)
	sigCh := make(chan os.Signal, 1) // канал,из которого будем читать отправленные сигналы
	ch := make(chan int)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM) // отправляем данные в канал
	// SIGINT - отправляется при введении Ctrl+c
	// SIGTERM - общий сигнал, используемый для завершения программы

	for i := 0; i < N; i++ { // запускам пул воркеров
		wg.Add(1)
		go printData(ch, &wg)
	}

	func() {
		for {
			select {
			case <-sigCh: // если пришел сигнал о завершении программы
				fmt.Println("Exit")
				return
			default:
				ch <- rand.Int() // отправляем рандомное значение в канал

			}
		}
	}()
	close(sigCh)
	close(ch)
	wg.Wait()

}
