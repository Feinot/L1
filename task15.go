package main

import "fmt"

//Проблема: Символы латиницы и некоторые специальные символы (ASCII) занимают 1 байт на 1 символ, а
//русские буквы уже занимают 2 байта на символ.
//Решение: конвертировать всю строку в rune

var justString string

func someFunc() {
	v := createHugeString(1 << 10) // размер строки в 2 раза больше ожидаемого
	r := []rune(v)                 // теперь мы будем работать с символами, а не с байтами
	justString = string(r[:100])
	fmt.Println(justString)
}
func createHugeString(size int) string {
	var v string
	for i := 0; i < size; i++ {
		v += "Ы" // добавляем к строке символ,занимающий 2 байта
	}

	return v
}

func main() {
	someFunc()
}
