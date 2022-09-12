package point

import "math"

//Point математическая точка
type Point struct {
	x, y float32
}

//SetX установить зачение координаты Х для точки
func (p *Point) SetX(x float32) {
	p.x = x
}

//SetY установить зачение координаты Y для точки
func (p *Point) SetY(y float32) {
	p.y = y
}

//Distance расчёт дистанции до указанной точки
func (p *Point) Distance(point *Point) (dist float64) {
	dist = math.Sqrt(math.Pow(float64(point.x-p.x), 2) + math.Pow(float64(point.y-p.y), 2))
	return
}

//NewPoint конструктор точки с установкой координат
func NewPoint(x, y float32) *Point {
	return &Point{x, y}
}

//NewZeroPoint конструктор точки с нулевыми координатами
func NewZeroPoint() *Point {
	return &Point{}
}
