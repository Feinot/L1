package main

import (
	"fmt"
)

/*Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской
структуры Human(аналог наследования).
*/
type Human struct {
	Name string
	Age  int
}

func (h Human) Print() { // метод "класса" Human
	fmt.Println(h.Age)
}

type Action struct {
	Human // встраивание структуры
}

func main() {
	a := Action{Human{Name: "alex", Age: 21}}
	fmt.Println(a.Name, a.Age)

}
