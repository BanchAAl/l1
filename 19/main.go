/*
Разработать программу, которая переворачивает подаваемую на вход строку (например: «главрыба — абырвалг»).
Символы могут быть unicode
*/
package main

import "fmt"

//flipString переворачивает строку. идёт по строке и пишет в результирующую с конца
func flipString(src string) (flip string) {
	fmt.Printf("\nsources string:\t%s", src)
	maxPos := len(src) - 1
	rslt := make([]rune, maxPos+1)

	for pos, ch := range src {
		rslt[maxPos-pos] = ch
	}

	flip = string(rslt)
	return
}

func main() {
	flip := flipString("Изподвыверта ворот видна хохлома. Пар - вода. шалаш")
	fmt.Printf("\nflip string:\t%s", flip)
}
