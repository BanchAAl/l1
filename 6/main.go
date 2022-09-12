/*
Реализовать все возможные способы остановки выполнения горутины
*/
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

//stopGoroutineWithWaitGroup остановка рутины с помощью sync.WaitGroup
func stopGoroutineWithWaitGroup() {
	wg := sync.WaitGroup{}
	numWork := 20
	wg.Add(numWork)

	go func(wg *sync.WaitGroup) {
		for i := 0; i < numWork; i++ {
			fmt.Printf("\nLoop work %d", i)
			time.Sleep(100 * time.Millisecond)
			wg.Done()
		}
	}(&wg)

	wg.Wait()

	fmt.Printf("\nStop gotoutine with wait group")
}

//stopGoroutineWithChannel остановка рутины через канал
func stopGoroutineWithChannel(stop chan bool) {
	for {
		select {
		case <-stop:
			fmt.Printf("\nstop go routine with channel")
			return
		default:
			fmt.Printf("\nany work channel stop")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

//stopGoroutineWithContext остановка рутины через контекст
func stopGoroutineWithContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("\nstop go routine with timeout")
			return
		default:
			fmt.Printf("\nany work context stop")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func main() {
	osNotify := make(chan os.Signal, 1)
	signal.Notify(osNotify, os.Interrupt)

	stopGoroutineWithWaitGroup()

	stop := make(chan bool)
	go stopGoroutineWithChannel(stop)
	// слипаем, чтобы дать немного поработать рутине, потом "неожиданно" приходят данные в канал завершения
	time.Sleep(2 * time.Second)
	stop <- true

	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	go stopGoroutineWithContext(ctx)

	<-osNotify
	fmt.Printf("\ngood luck")
}
