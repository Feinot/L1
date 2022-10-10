package main

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/
func main() {
	arr := []int{1, 3, 5, 6}
	i := 1
	copy(arr[i:], arr[i+1:]) // copy принимает два параметра: куда будем копировать и откуда
	fmt.Println(arr)         // [1 5 6 6]
	arr[len(arr)-1] = 0
	arr = arr[:len(arr)-1]
	fmt.Println(arr) // [1 5 6]

}
