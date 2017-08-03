// An example that does not respect the Liskov substitution principle
// go run main.go square.go

package main

import "fmt"

// FigureInterface ...
type FigureInterface interface {
	SetWidth(width rune)
	SetHeight(height rune)
	GetWidth() rune
	GetHeight() rune
}

// Rectangle ...
type Rectangle struct {
	height rune
	width  rune
}

// SetWidth ...
func (r *Rectangle) SetWidth(width rune) {
	r.width = width
}

// SetHeight ...
func (r *Rectangle) SetHeight(height rune) {
	r.height = height
}

// GetHeight ...
func (r *Rectangle) GetHeight() rune {
	return r.height
}

// GetWidth ...
func (r *Rectangle) GetWidth() rune {
	return r.width
}

// AreaOfTheFigure ...
func AreaOfTheFigure(figure FigureInterface, height, width rune) rune {
	figure.SetHeight(height)
	figure.SetWidth(width)

	return figure.GetHeight() * figure.GetWidth()
}

func main() {
	rect := &Rectangle{}
	square := &Square{}

	fmt.Println(AreaOfTheFigure(square, 10, 5))
	fmt.Println(AreaOfTheFigure(rect, 10, 5))
}
