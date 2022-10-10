package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает,
вычитает две числовых переменных a,b, значение которых > 2^20.
*/

func BigAdd(a int64, b int64) big.Int {
	var z big.Int
	a1 := big.NewInt(a)
	b1 := big.NewInt(b)
	z.Add(a1, b1)
	return z
}

func BigSub(a int64, b int64) big.Int {
	var z big.Int
	a1 := big.NewInt(a)
	b1 := big.NewInt(b)
	z.Sub(a1, b1)
	return z
}

func BigMul(a int64, b int64) big.Int {
	var z big.Int
	a1 := big.NewInt(a)
	b1 := big.NewInt(b)
	z.Mul(a1, b1)
	return z
}

func BigDiv(a int64, b int64) big.Int {
	var z big.Int
	a1 := big.NewInt(a)
	b1 := big.NewInt(b)
	z.Div(a1, b1)
	return z
}

func main() {
	var a, b int64 = 53556149123455552, 3226799943225626

	fmt.Println(BigAdd(a, b))
	fmt.Println(BigSub(a, b))
	fmt.Println(BigMul(a, b))
	fmt.Println(BigDiv(a, b))

}
