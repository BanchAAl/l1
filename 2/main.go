// Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из
// массива (2,4,6,8,10) и выведет их квадраты в stdout
package main

import (
	"fmt"
	"sync"
)

func squaring(digits ...int) {
	var wg sync.WaitGroup
	for _, digit := range digits {
		wg.Add(1)
		go func(wg *sync.WaitGroup, digit int) {
			fmt.Printf("squaring for digit %d : %d \n", digit, digit*digit)
			wg.Done()
		}(&wg, digit)
	}
	wg.Wait()
}

func main() {
	digits := []int{2, 4, 6, 8, 10}

	squaring(digits...)
}
