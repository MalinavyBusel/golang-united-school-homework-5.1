package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() (p Point) {
	return Point{
		x: s.start.x + int(s.a),
		y: s.start.y + int(s.a),
	}
}

func (s Square) Perimeter() (perimeter uint) {
	return s.a * 4
}

func (s Square) Area() (area uint) {
	return s.a * s.a
}
