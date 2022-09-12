/*
Реализовать паттерн «адаптер» на любом примере
*/
package main

import "fmt"

// Target новый интерфейс с которым должен работать, написанный ранее код
type Target interface {
	Request() string
}

// Adaptee адаптируемый тип, который не может работать с новым интерфейсом
type Adaptee struct {
}

// NewAdapter конструктор адаптера
func NewAdapter(adaptee *Adaptee) Target {
	return &Adapter{adaptee}
}

// SpecificRequest "старая" реализация
func (a *Adaptee) SpecificRequest() string {
	return "Request"
}

// Adapter тип адаптера, позволяющий работать с новыминтерфейсом.
type Adapter struct {
	*Adaptee
}

// Request реализация метода адаптации
func (a Adapter) Request() string {
	return a.SpecificRequest()
}

func checkAdapter(t Target) {
	fmt.Printf("\n adapted: %s", t.Request())
}

func main() {
	var atee Adaptee
	ader := Adapter{&atee}
	checkAdapter(ader)

}
