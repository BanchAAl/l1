/*
Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с
инкапсулированными параметрами x,y и конструктором
*/
package main

import (
	"fmt"

	"24/pkg/point"
)

func main() {
	mPoint := point.NewZeroPoint()
	nPoint := point.NewZeroPoint()

	mPoint.SetX(12)
	mPoint.SetY(10)
	nPoint.SetX(23)
	nPoint.SetY(15)

	dist := mPoint.Distance(nPoint)

	fmt.Printf("distance between %+v and %+v: %f", mPoint, nPoint, dist)
}
