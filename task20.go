package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/
func main() {
	s := "snow dog sun"
	arr := strings.Split(s, " ")

	for i := 0; i < len(arr)/2; i++ {

		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}

	fmt.Println(strings.Join(arr, " "))

}
