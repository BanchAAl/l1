/*
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться
*/
package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// startWriteWithContext запись в канал с ограничением времени работы с использованием контекста
func startWriteWithContext(dataChannel chan string, ctx context.Context) {
	fmt.Println("\nstart write with context")
	start := time.Now()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("\nclose timeout")
			fmt.Printf("\nStartWriteWithContext duration: %d", time.Since(start).Milliseconds())
			return
		default:
			data := "time context: Message namber " + strconv.Itoa(rand.Int())
			dataChannel <- data
			time.Sleep(200 * time.Millisecond)
		}
	}
}

// startWriteWithTimeSince запись в канал с ограничением времени работы с использование duration
func startWriteWithTimeSince(dataChannel chan string, duration int) {
	fmt.Printf("\nstart write with time since")
	start := time.Now()
	for {
		data := "time since: Message number " + strconv.Itoa(rand.Int())
		dataChannel <- data
		time.Sleep(200 * time.Millisecond)
		if d := time.Since(start); d.Seconds() > float64(duration) {
			fmt.Printf("\nSTOP")
			fmt.Printf("\nStartWriteWithTimeSince duration: %d", time.Since(start).Milliseconds())
			return
		}
	}

}

// startRead чтение из канала
func startRead(dataChannel chan string) {
	fmt.Printf("\nstart read")
	go func(data chan string) {
		for {
			select {
			case msg := <-data:
				fmt.Printf("\n" + msg)
			}
		}
	}(dataChannel)
}

func main() {
	var sec int

	fmt.Println("Введите время жизни задачи (в секундах)")
	_, err := fmt.Scan(&sec)
	if err != nil {
		fmt.Println("ошибка ввода")
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(sec)*time.Second)
	dataChannel := make(chan string)
	go startRead(dataChannel)
	startWriteWithContext(dataChannel, ctx)
	startWriteWithTimeSince(dataChannel, sec)
}
