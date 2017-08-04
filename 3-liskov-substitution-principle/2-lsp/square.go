// Example

package main

// Square ...
type Square struct {
	Rectangle
}

// SetWidth ...
func (s *Square) SetWidth(width rune) {
	s.width = width
	s.height = width
}

// SetHeight ...
func (s *Square) SetHeight(height rune) {
	s.height = height
	s.width = height
}

// GetHeight ...
func (s *Square) GetHeight() rune {
	return s.height
}

// GetWidth ...
func (s *Square) GetWidth() rune {
	return s.width
}
