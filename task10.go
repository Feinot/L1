package main

import "fmt"

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножноствах не важна.
*/
func main() {
	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	values := make(map[int][]float64)

	for _, tmp := range temp {
		key := int(tmp) / 10 * 10 // округляем до ближайшего десятка
		values[key] = append(values[key], tmp)

	}
	fmt.Println(values)
}
