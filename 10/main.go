/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Объединить данные
значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/
package main

import "fmt"

/*
categorization раскладываем значения температур по группам.
берём очередное значение, получаем ключ - "ровный десяток", кидаем в мапу слайсов по ключу, если такого нет, делаем
новый слайс.
Особый случай - ноль. По условию, в категорию пишутся значения "текщего десятка", но, в этом случае, к 0 относятся
значения от -9.99 до 9.99.
*/
func categorization(data []float32) {
	categories := make(map[int][]float32)
	for _, val := range data {
		key := int(val/10) * 10
		_, ok := categories[key]
		if ok {
			categories[key] = append(categories[key], val)
		} else {
			categories[key] = []float32{val}
		}
	}
	fmt.Println(categories)
}

func main() {
	data := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, -0.2, 0.45, -0.23, 0.67, -8, 7, 8.7, -5.5, 24.5, -21.0, 32.5, 123.3, 125, 131, 134.2}
	categorization(data)
}