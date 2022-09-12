/*
Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной
типа interface{}.
*/
package main

import (
	"fmt"
)

//typeOf определяет тип переменной.
func typeOf(val interface{}) {
	switch v := val.(type) {
	case int:
		fmt.Printf("\nType INT")
	case string:
		fmt.Printf("\nType STRING")
	case bool:
		fmt.Printf("\nType BOOL")
	case chan int:
		fmt.Printf("\nType CHAN INT")
	case chan string:
		fmt.Printf("\nType CHAN STRING")
	case chan bool:
		fmt.Printf("\nType CHAN BOOL")
	default:
		//тоже покажет тип, но просто покажет.
		fmt.Printf("\nI don't know about type %T!\n", v)
	}
}

func main() {
	fmt.Printf("\nвыберите тип. \nint - 0\nstring - 1\nbool - 2\nchan int - 3\nchan string - 4\nchan bool - 5\n")
	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil || choice < 0 || choice > 5 {
		fmt.Printf("\nнекорректный ввод. завершаем.")
		return
	}
	var vInterface interface{}
	switch choice {
	case 0:
		var a int
		vInterface = a
	case 1:
		var a string
		vInterface = a
	case 2:
		var a bool
		vInterface = a
	case 3:
		var a chan int
		vInterface = a
	case 4:
		var a chan string
		vInterface = a
	case 5:
		var a chan bool
		vInterface = a
	}
	typeOf(vInterface)
}
