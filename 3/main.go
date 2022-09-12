/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

// интерфейс для сумматоров
type summator interface {
	addTo(int)
	getSum() int
}

// потокобезопасный сумматор
type syncSummator struct {
	sum int
	mu  sync.Mutex
}

// плюсуем значение к сумматору
func (s *syncSummator) addTo(value int) {
	s.mu.Lock()
	s.sum += value
	s.mu.Unlock()
}

// получаем значение сумматора
func (s *syncSummator) getSum() (value int) {
	s.mu.Lock()
	value = s.sum
	s.mu.Unlock()
	return
}

// сумматор через канал
type channelSummator struct {
	sum  int
	ch   chan int
	stop chan bool
}

// плюсуем значение к сумматору
func (c *channelSummator) addTo(value int) {
	c.ch <- value
}

// получаем значение сумматора
func (c *channelSummator) getSum() (value int) {
	return c.sum
}

// запуск прослушивания канала для суммирования
func (c *channelSummator) startSummator() {
	go func() {
		for {
			select {
			case digit := <-c.ch:
				c.sum += digit
			case v := <-c.stop:
				if v {
					break
				}
			}
		}
	}()
}

// останавливаем суммирование
func (c *channelSummator) close() {
	c.stop <- true
}

// конструктор сумматора через канал
func newChannelSummator() *channelSummator {
	sum := channelSummator{0, make(chan int), make(chan bool)}
	return &sum
}

// считаем квадрат и добавляем значение к сумме
func calc(wg *sync.WaitGroup, sum summator, valueForAdd int) {
	defer wg.Done()
	square := valueForAdd * valueForAdd
	switch t := sum.(type) {
	case *syncSummator:
		sum.addTo(square)
	case *channelSummator:
		sum.addTo(square)
	default:
		fmt.Printf("%T summator??\n", t)
	}
}

// считаем сумму квадратов
func calcSum(digits []int, sum summator) {
	now := time.Now()
	var wg sync.WaitGroup
	for _, digit := range digits {
		wg.Add(1)
		go calc(&wg, sum, digit)
	}
	wg.Wait()
	duration := time.Since(now)

	fmt.Printf("Duration: %d nanoseconds\n", duration.Nanoseconds())
	return
}

func main() {
	digits := []int{2, 4, 6, 8, 10}

	// считаем сумму через синхронизацию
	var sSum syncSummator
	calcSum(digits, &sSum)
	fmt.Printf("sum square without sync for %v - %d\n", digits, sSum.getSum())

	// считаем сумму через канал
	chSum := newChannelSummator()
	chSum.startSummator()
	calcSum(digits, chSum)
	chSum.close()
	fmt.Printf("sum square with channel for %v - %d\n", digits, chSum.getSum())
}
