/*
Реализовать пересечение двух неупорядоченных множеств
*/
package main

import (
	"fmt"
	"time"
)

//intersection ищем пересечение переданных множеств
func intersection(a, b []int) {
	start := time.Now()

	fmt.Printf("\na %v", a)
	fmt.Printf("\nb %v", b)

	intersectionMap := make(map[int]int)
	rslt := make([]int, 0)
	var maxLen int

	aLen := len(a)
	bLen := len(b)

	if aLen >= bLen {
		maxLen = aLen
	} else {
		maxLen = bLen
	}

	for i := 0; i < maxLen; i++ {
		if i < aLen {
			aVal := a[i]
			if freq, ok := intersectionMap[aVal]; ok {
				if freq == 1 {
					rslt = append(rslt, aVal)
					intersectionMap[aVal] = 2
				}
			} else {
				intersectionMap[aVal] = 1
			}
		}
		if i < bLen {
			bVal := b[i]
			if freq, ok := intersectionMap[bVal]; ok {
				if freq == 1 {
					rslt = append(rslt, bVal)
					intersectionMap[bVal] = 2
				}
			} else {
				intersectionMap[bVal] = 1
			}
		}
	}

	fmt.Printf("\nintersection %v\n intersection duration %d", rslt, time.Since(start).Nanoseconds())
}

//intersectionSlow ищем пересечение переданных множеств.работает медленнее, т.к. больше проходов
func intersectionSlow(a, b []int) {
	start := time.Now()

	fmt.Printf("\na %v", a)
	fmt.Printf("\nb %v", b)

	intersectionMap := make(map[int]int)
	rslt := make([]int, 0)

	for _, val := range a {
		if _, ok := intersectionMap[val]; !ok {
			intersectionMap[val] = 1
		}
	}
	for _, val := range b {
		if _, ok := intersectionMap[val]; ok {
			intersectionMap[val] = 2
		}
	}
	for k, v := range intersectionMap {
		if v == 2 {
			rslt = append(rslt, k)
		}
	}

	fmt.Printf("\nintersection %v\n intersectionSlow duration %d", rslt, time.Since(start).Nanoseconds())
}

func main() {
	a := []int{2, 3, 6, 12, 2, 67, 57, 75, 8, 5, 42}
	b := []int{5, 2, 8, 54, 8, 23, 44, 42, 1, 3, 78, 9, 45}

	intersectionSlow(a, b)
	fmt.Printf("\n")
	intersection(a, b)
}
