/*
Поменять местами два числа без создания временной переменной.
*/
package main

import "fmt"

func main() {
	a, b := 1, 3
	fmt.Printf("a=%d, b=%d\n", a, b)
	//сначала произойдёт подстановка значений справа, а потом присваивание
	a, b = b, a
	fmt.Printf("a=%d, b=%d\n", a, b)
}
