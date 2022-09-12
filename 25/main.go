/*
Реализовать собственную функцию sleep
*/
package main

import (
	"context"
	"fmt"
	"time"
)

//sleep реализация слип через контекст с таймаутом
func sleep(duration time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	select {
	case <-ctx.Done():
		return
	}
}

//sleep2 реализация слип через таймафтер
func sleep2(duration time.Duration) {
	select {
	case <-time.After(duration):
		return
	}
}

func main() {
	start := time.Now()
	sleep(5 * time.Second)
	fmt.Printf("\n duration sleep %f", time.Since(start).Seconds())
	start = time.Now()
	sleep2(5 * time.Second)
	fmt.Printf("\n duration sleep %f", time.Since(start).Seconds())
}
