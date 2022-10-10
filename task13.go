package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/
func main() {
	a, b := 3, 8
	fmt.Println(a, b) // 3 8
	b, a = a, b
	fmt.Println(a, b) // 8 3
	/*
		a = a * b
		b = a / b
		a = a / b
	*/
}
