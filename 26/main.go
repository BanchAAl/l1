/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.
Например:
	abcd — true
	abCdefAaf — false
	aabcd — false
*/
package main

import (
	"fmt"
	"strings"
)

//checkUniq проверка уникальности символов в строке.
func checkUniq(src string) bool {
	fmt.Printf("\nsource string: %s", src)

	freq := make(map[int32]int)

	src = strings.ToLower(src)

	for _, v := range src {
		if _, ok := freq[v]; ok {
			return false
		}
		freq[v] = 1
	}

	return true
}

func main() {
	check := checkUniq("abcd")
	fmt.Printf("\nabcd: uniq - %t", check)

	check = checkUniq("abCdefAaf")
	fmt.Printf("\nabCdefAaf: uniq - %t", check)

	check = checkUniq("aabcd")
	fmt.Printf("\naabcd: uniq - %t", check)

	check = checkUniq("ilubhbuUhuibiophiounioIUyb")
	fmt.Printf("\nilubhbuUhuibiophiounioIUyb: uniq - %t", check)
}
