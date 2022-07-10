package main

type Point struct {
	x, y int
}

type Square struct {
	start Point
	side  uint
}

func (s Square) End() (p Point) {
	return Point{
		x: s.start.x + int(s.side),
		y: s.start.y + int(s.side),
	}
}

func (s Square) Perimeter() (perimeter uint) {
	return s.side * 4
}

func (s Square) Area() (area uint) {
	return s.side * s.side
}
