package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками,
 которые представлены в виде структуры
Point с инкапсулированными параметрами x,y и конструктором.
*/

type Point struct {
	x, y int
}

func NewPoint(X, Y int) *Point {
	return &Point{
		x: X,
		y: Y,
	}
}

func Distance(a *Point, b *Point) float64 {

	return math.Sqrt(math.Pow(float64(b.x-a.x), 2) + math.Pow(float64(b.y-a.y), 2))

}

func main() {
	a := NewPoint(3, 4)
	b := NewPoint(5, 8)
	res := Distance(a, b)
	fmt.Println(res)
}
