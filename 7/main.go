/*
Реализовать конкурентную запись данных в map.
*/
package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// printSyncMap печать пар ключ-значение для синк мапы
func printSyncMap(key, value any) (ok bool) {
	k, okKey := key.(int)
	v, okVal := value.(string)
	ok = okKey && okVal
	if ok {
		fmt.Printf("{%d: %s}", k, v)
	}
	return
}

//writeToMap пишем в мапу. реализовано 2 варианта: мапа с мьютексом и синк мапа
func writeToMap(dataChannel chan string, ctx context.Context) {
	dataMap := make(map[int]string)
	var mu sync.Mutex
	var wg sync.WaitGroup
	var sMap sync.Map
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(dataChannel chan string, wg *sync.WaitGroup) {
			for {
				select {
				case <-ctx.Done():
					wg.Done()
					return
				case data := <-dataChannel:
					fmt.Printf("\n write to map %s", data)
					// 1 вариант - обычная мапа с использоанием мьютекса
					index := len(dataMap)
					mu.Lock()
					dataMap[index] = data
					mu.Unlock()

					//2 вариант - синк мапа
					sMap.Store(index, data)
				}
			}
		}(dataChannel, &wg)
	}
	wg.Wait()
	//смотрим, что данные в мапах. для компактности вывода в простой мапе печатаем только длину
	fmt.Printf("\nsimple map length %d", len(dataMap))
	fmt.Printf("\n")
	fmt.Printf("\nsync map data:\n")
	sMap.Range(printSyncMap)
}

//sendToMap формируем сообщения и отправляем в канал
func sendToMap(dataChannel chan string, ctx context.Context) {
	i := 0
	for {
		i++
		select {
		case <-ctx.Done():
			return
		default:
			data := fmt.Sprintf("message %d: %d", i, rand.Int())
			dataChannel <- data
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	dataChannel := make(chan string)
	go sendToMap(dataChannel, ctx)
	writeToMap(dataChannel, ctx)
}
