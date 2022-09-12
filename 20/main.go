/*
Разработать программу, которая переворачивает слова в строке. Пример: «snow dog sun — sun dog snow»
*/
package main

import "fmt"

//flipStringByWords переворачивает предложение. Читает строку, выбирая целые слова, полученные слова пишет задом
//наперёд в результирующую строку
func flipStringByWords(src string) (flip string) {
	fmt.Printf("\nsource string: %s", src)
	var word []rune
	for _, ch := range src {
		//отсекаем слово при первом попавшемся пробеле
		if ch == ' ' {
			if len(word) > 0 {
				flip = string(word) + " " + flip
				word = word[:0]
			}
			//игнорируем повторяющиеся пробелы
			continue
		}

		word = append(word, ch)
	}
	if len(word) > 0 {
		flip = string(word) + " " + flip
	}

	return
}

func main() {
	flip := flipStringByWords("Йода Мастер - без труда       не вытащишь рыбку из пруда")
	fmt.Printf("\nresult string: %s", flip)
}
