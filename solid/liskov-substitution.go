package solid

import "fmt"

// Liskov Substitutions Principle
// named after barbara liskov principle
// :=  Suppose we have an api that take base class & work correctly it should also worked in derived class

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetHeight() * sized.GetWidth()

	fmt.Println("Expected area: ", expectedArea, "\n ", ", but got: ", actualArea)

}

// a problem here ... the principle broke here
type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := &Square{}
	sq.width = size
	sq.height = size
	return sq
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.height = height
	s.width = height
}

// frist approach to solve
type Square2 struct {
	size int // width, height
}

func (s *Square2) Rectangle() *Rectangle {
	return &Rectangle{
		width:  s.size,
		height: s.size,
	}
}

func (s *Square2) SetWidth(width int) {
	s.size = width
	s.size = width
}

func (s *Square2) SetHeight(height int) {
	s.size = height
	s.size = height
}
