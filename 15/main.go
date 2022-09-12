/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный
пример реализации
var justString string
func someFunc() {
v := createHugeString(1 << 10)
justString = v[:100]
}

func main() {
someFunc()
*/
package main

var justString string

func createHugeString(len int) (str string) {
	// out of range -негативные последствия
	//buf := make([]byte, 0, len)
	// correct
	buf := make([]byte, len)
	return string(buf)
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}
