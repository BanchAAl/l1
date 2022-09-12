// Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в
// структуре Action от родительской структуры Human (аналог наследования).
package main

import "fmt"

type human struct {
	name string
	age  int
}

func (h human) Walk() {
	fmt.Printf("Hello. I'm walking!! I'm %s\n", h.name)
}

func (h human) Talk() {
	fmt.Printf("I'm talking!! my age - %d\n", h.age)
}

// добавляем анонимное поле типа human, чтобы в коде использовеать его данные методы "как свои" (как при
// наследованиии в плюсах), а не через имя поля
type action struct {
	human
}

func (a action) Move() {
	fmt.Printf("it's amazing. Name - %s, age - %d\n", a.name, a.age)
}

func main() {
	var act action
	act.name = "Ivan"
	act.age = 25
	act.Walk()
	act.Talk()
	act.Move()
}
