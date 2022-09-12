/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/
package main

import (
	"errors"
	"fmt"
)

//createSubset создаём собственное подмножество. Оно должно принадлежать исходному множеству, но не должно быть равно
//ему и не должно быть пустым. Выбираем уникальные значения, если это не удовлетворит условию, просто выкинем последний
//элемент
func createSubset(set []string) (err error) {
	if len(set) == 0 {
		err = errors.New("исходное множество пустое. в пустом множестве невозможно найти собственное подмножество")
		return
	}
	frequency := make(map[string]int)
	subset := make([]string, 0)
	for _, v := range set {
		if _, ok := frequency[v]; !ok {
			frequency[v] = 1
			subset = append(subset, v)
		}
	}

	if len(set) == len(subset) {
		subset = subset[:len(subset)-1]
	}

	fmt.Printf("\nsubset %v", subset)

	return
}

func main() {
	set := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Printf("\nsource set %v", set)
	//1й вариант построения собственного подмножества
	err := createSubset(set)
	if err != nil {
		fmt.Println(err.Error())
	}

	//2й вариант (элементарный) построения подмножества
	if len(set) != 0 {
		subset := set[0]
		fmt.Printf("\nsimple subset [%v]", subset)
	}
}
