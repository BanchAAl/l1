/*
 	Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают
	произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
	Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/
package main

import (
	"bufio"
	"fmt"
	"os"

	"4/pkg/workers"
)

// запуск записи в канал
func startWrite(dataChannel chan string) {
	for {
		stdin := bufio.NewReader(os.Stdin)
		fmt.Printf("\nВведите сообщение\n")

		var message string

		_, err := fmt.Fscan(stdin, &message)
		if err != nil {
			stdin.ReadString('\n')
			fmt.Printf("\nЧто-то пошло не так.")
			continue
		}
		dataChannel <- message
	}
}

func main() {
	var numWorkers int

	fmt.Println("Введите нужное количество воркеров (число строго большее нуля)")
	_, err := fmt.Scan(&numWorkers)
	if err != nil {
		fmt.Println("ошибка ввода")
		return
	}

	dataChannel := make(chan string)
	pool := workers.NewPoolWorkers()
	pool.CreateWorkers(numWorkers, dataChannel)
	go startWrite(dataChannel)
	pool.StartWorkers()
}
