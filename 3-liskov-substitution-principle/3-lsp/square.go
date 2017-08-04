//Example

package main

// Square ...
type Square struct {
	Rectangle
	height rune
	width  rune
	size   rune
}

// SetSize ...
func (s *Square) SetSize(size rune) {
	s.size = size
}

// GetSize ...
func (s *Square) GetSize() rune {
	return s.size
}
