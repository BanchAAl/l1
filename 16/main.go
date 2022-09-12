/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка
*/
package main

import "fmt"

//qSort реализация быстрой сортировки
func qSort(ar []int) {
	//дальше нечего разбивать
	if len(ar) <= 1 {
		return
	}

	split := partition(ar)

	// рубим массив на две части
	qSort(ar[:split])
	qSort(ar[split:])
}

//partition делим массив на части. опорный элемент выбираем посередине.
func partition(ar []int) int {
	pivot := ar[len(ar)/2]

	left := 0
	right := len(ar) - 1

	for {
		for ; ar[left] < pivot; left++ {
		}

		for ; ar[right] > pivot; right-- {
		}

		if left >= right {
			return right
		}

		//нашли левый элемент больше опорного и правый меньше опорного - надо поменять их местами
		swap(ar, left, right)
	}
}

//swap меняем элементы местами.
func swap(ar []int, i, j int) {
	tmp := ar[i]
	ar[i] = ar[j]
	ar[j] = tmp
}

func main() {
	f := []int{2, 67, 4, 23, 3, 5, 12, 98, 101, 34, 25}
	fmt.Printf("\n src: %v", f)
	qSort(f)
	fmt.Printf("\n sort: %v", f)
}
