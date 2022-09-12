/*
Реализовать бинарный поиск встроенными методами языка
*/
package main

import (
	"17/pkg/sort"
	"fmt"
)

//binaryFind реализация бинарного поиска. Поиск осуществляется по отсортированному массиву.
func binaryFind(arr []int, val int) bool {
	var mid, tmp int

	first := 0
	last := len(arr) - 1

	for first <= last {
		mid = (first + last) / 2
		tmp = mid
		if arr[tmp] == val {
			return true
		}
		//массив отсортирован, "середину" уже проверили
		if arr[tmp] > val {
			//ищем в левой части
			last = mid - 1
		} else {
			//ищем в правой части
			first = mid + 1
		}
	}
	return false
}

func main() {
	arr := []int{2, 65, 12, 654, 11, 45}
	sort.QSort(arr)
	fmt.Printf("\n введите число для поиска в массиве\n")
	var val int
	_, err := fmt.Scan(&val)
	if err != nil {
		fmt.Printf("\nerror input. exit...")
	}
	ok := binaryFind(arr, val)
	fmt.Printf("\n Value %d exist in array %v? - %t", val, arr, ok)
}
