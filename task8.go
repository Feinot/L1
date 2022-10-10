package main

import "fmt"

/*
Дана переменная int64.
Разработать программу которая устанавливает i-й бит в 1 или 0.
*/
func main() {
	var (
		val int64
		i   int
	)

	fmt.Scan(&val, &i)

	val |= 1 << i // установить бит в 1
	fmt.Println(val)
	val &= ^(1 << i) // сбросить бит
	fmt.Println(val)

}
