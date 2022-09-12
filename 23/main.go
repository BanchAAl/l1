/*
Удалить i-ый элемент из слайса.
*/
package main

import (
	"fmt"
	"time"
)

//deleteFromSlice удаление элемента слайса. сохраняется порядок расположения элементов.
func deleteFromSlice(data []int, pos int) (rslt []int, d int64) {
	start := time.Now()
	maxPos := len(data) - 1

	if pos > maxPos+1 {
		fmt.Println("\nout of range")
		rslt = data
		d = time.Since(start).Nanoseconds()
		return
	}

	if pos == maxPos {
		rslt = data[:maxPos]
		d = time.Since(start).Nanoseconds()
		return
	}

	rslt = append(rslt, data[:pos]...)
	rslt = append(rslt, data[pos+1:]...)
	d = time.Since(start).Nanoseconds()
	return
}

//deleteFromSliceFast удаление элемента слайса. без сохранения порядка расположения элементов, но быстрее.
func deleteFromSliceFast(data []int, pos int) (rslt []int, d int64) {
	start := time.Now()
	maxPos := len(data) - 1

	if pos > maxPos {
		fmt.Println("\nout of range")
		return
	}
	data[pos] = data[maxPos]
	rslt = data[:maxPos]
	d = time.Since(start).Nanoseconds()
	return
}

func main() {
	digits := []int{2, 4, 6, 8, 10}

	fmt.Printf("\n Слайс для удаления %v", digits)
	var pos int
	fmt.Printf("\n Введите позицию для удаления 0 - %d", len(digits)-1)
	fmt.Printf("\n")
	_, err := fmt.Scan(&pos)
	if err != nil || pos < 0 || pos > len(digits) {
		fmt.Printf("\n Ошибочный ввод. Завершаем.")
		return
	}

	fmt.Printf("\nsource data\t%v", digits)

	digitsNew, d := deleteFromSlice(digits, pos)
	fmt.Printf("\nduration %d nsec", d)
	fmt.Printf("\nmodifed data\t%v", digitsNew)

	digitsNew, d = deleteFromSliceFast(digits, pos)
	fmt.Printf("\nmodifed data\t%v", digitsNew)
	fmt.Printf("\nduration %d nsec", d)
}
