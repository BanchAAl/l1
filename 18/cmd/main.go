/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна
выводить итоговое значение счетчика
*/
package main

import (
	"fmt"
	"sync"

	"18/pkg/counters"
)

//checkCounter проверка конкурентного счётчика
func checkCounter(counter *counters.SyncCounter) {
	var c counters.SyncCounter
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, c *counters.SyncCounter) {
			for j := 0; j < 3; j++ {
				fmt.Printf("\n\tCounter value: %d", counter.Value())
				counter.Incr()
			}
			fmt.Printf("\n\tCounter value: %d", counter.Value())
			counter.Decr()
			wg.Done()
		}(&wg, &c)
	}

	wg.Wait()
	fmt.Printf("\n\tCounter total: %d", counter.Value())
}

func main() {
	fmt.Printf("\n Sync Counter")
	var sCounter counters.SyncCounter
	checkCounter(&sCounter)
}
