/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b,
значение которых > 2^20
*/
package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1 << 60)
	b := big.NewInt(1 << 50)
	c := big.NewInt(0).Div(a, b)
	fmt.Printf("\n %d/%d=%d", a, b, c)

	k := big.NewFloat(1 << 62)
	m := big.NewFloat(1 << 59)
	n := big.NewFloat(0).Mul(k, m)
	fmt.Printf("\n %f*%f=%f", k, m, n)

	x := big.NewFloat(1 << 62)
	y := big.NewFloat(1 << 59)
	z := big.NewFloat(0).Add(x, y)
	fmt.Printf("\n %f+%f=%f", x, y, z)

	g := big.NewFloat(1 << 40)
	o := big.NewFloat(1 << 30)
	d := big.NewFloat(0).Sub(g, o)
	fmt.Printf("\n %f-%f=%f", g, o, d)

}
