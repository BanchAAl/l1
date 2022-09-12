/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/
package main

import (
	"fmt"
)

// setBit установка указанного значения бита в указанной позиции
func setBit(val int64, position int, set int) int64 {
	var newVal int64
	var mask int64
	switch set {
	case 0:
		mask = ^(int64(1) << position)
		newVal = val & mask
	case 1:
		mask = int64(1) << position
		newVal = val | mask
	default:
		newVal = val
	}

	return newVal
}

func main() {
	fmt.Println("\nВведите число")

	var val int64
	var position int
	var bitValueSet int

	_, err := fmt.Scan(&val)
	if err != nil {
		fmt.Println("\nОшибка ввода. Завершаем. ")
		return
	}

	fmt.Println("\nВведите позицию бита 0 - 63")
	_, err = fmt.Scan(&position)
	if err != nil || position < 0 || position > 63 {
		fmt.Println("\nОшибка ввода. Завершаем. ")
		return
	}
	fmt.Println("\nВведите значение бита 0, 1")
	_, err = fmt.Scan(&bitValueSet)
	if err != nil && bitValueSet < 0 || bitValueSet > 1 {
		fmt.Println("\nОшибка ввода. Завершаем. ")
		return
	}

	fmt.Printf("\nstart\t%d\t-\t%b", val, val)
	newVal := setBit(val, position, bitValueSet)
	fmt.Printf("\nexit\t%d\t-\t%b", newVal, newVal)
}
