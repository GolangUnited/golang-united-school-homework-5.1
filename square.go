package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s *Square) End() Point {
	// implement me

	var endPoint Point
	endPoint.x = s.start.x + int(s.a)
	endPoint.y = s.start.y + int(s.a)

	return endPoint

}

func (s *Square) Area() uint {
	// implement me
	area := s.a * s.a
	return area
}

func (s *Square) Perimeter() uint {
	// implement me

	perimeter := 4 * s.a

	return perimeter

}
