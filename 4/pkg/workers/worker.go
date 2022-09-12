// Package workers
/*
Создаётся пул воркеров для удобства обращения с ними.
Запуск остановки воркеров извне происходит централизовано через пул, что позволяет остановить воркеры так, чтобы
можно было выполнить некую работу (в данном случае элементарную) перед остановкой воркера.
sync.WaitGroup позволяет дождаться остановки работы всех воркеров.
*/
package workers

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

// Pool Пул воркеров
type Pool struct {
	workers  []*worker
	osNotify chan os.Signal
	wg       sync.WaitGroup
}

// CreateWorkers создание нужного количества воркеров
func (wp *Pool) CreateWorkers(num int, data chan string) {
	for i := 0; i < num; i++ {
		wrk := newWorker(i, data, &wp.wg)
		wp.workers = append(wp.workers, wrk)
	}
}

//StartWorkers запуск чтения из канала
func (wp *Pool) StartWorkers() {
	wp.osNotify = make(chan os.Signal, 1)
	signal.Notify(wp.osNotify, os.Interrupt)
	for _, wrk := range wp.workers {
		wrk.start()
	}
	go func() {
		select {
		case <-wp.osNotify:
			for _, wrk := range wp.workers {
				wrk.stop()
			}
			return
		}
	}()

	wp.wg.Wait()
}

//NewPoolWorkers конструктор пула воркеров
func NewPoolWorkers() *Pool {
	return &Pool{}
}

// Воркер, читающий из канала
type worker struct {
	dataChannel chan string
	exit        chan bool
	numWorker   int
	wg          *sync.WaitGroup
}

// запуск воркера
func (w *worker) start() {
	go func() {
		fmt.Printf("\nWorker with nubmer %d - started", w.numWorker)
		for {
			select {
			case data := <-w.dataChannel:
				fmt.Printf("\nWorker with number %d; receive data: %s\n", w.numWorker, data)
			case <-w.exit:
				fmt.Printf("\nWorker with number %d - stopped", w.numWorker)
				w.wg.Done()
				return
			}
		}
	}()
}

// остановка воркера
func (w *worker) stop() {
	w.exit <- true
}

// конструктор воркера
func newWorker(number int, dataChannel chan string, wg *sync.WaitGroup) *worker {
	wg.Add(1)
	return &worker{dataChannel, make(chan bool), number, wg}
}
