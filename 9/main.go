/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	digits := []int{2, 4, 6, 8, 10}

	simpleChan := make(chan int)
	doubleChan := make(chan int)
	var wg sync.WaitGroup

	go func() {
		for {
			select {
			case digit := <-simpleChan:
				dbl := 2 * digit
				doubleChan <- dbl
			}
		}
	}()

	go func(wg *sync.WaitGroup) {
		for {
			select {
			case double := <-doubleChan:
				fmt.Printf("\n%d", double)
				wg.Done()
			}
		}
	}(&wg)

	fmt.Printf("\n%v", digits)
	for _, digit := range digits {
		wg.Add(1)
		simpleChan <- digit
	}

	wg.Wait()
}
