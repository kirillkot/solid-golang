// go run main.go
// Example open/closed principle
// strategy pattern

package main

import "fmt"

// Calculer ...
type Calculer interface {
	Execute(rune, rune) rune
}

// Add ...
type Add struct{}

// Minus ...
type Minus struct{}

// Execute ...
func (a Add) Execute(numA, numB rune) rune {
	return numA + numB
}

// Execute ...
func (m Minus) Execute(numA, numB rune) rune {
	return numA - numB
}

// Calcul ...
type Calcul struct {
	c Calculer
}

// Calculate ...
func (c Calcul) Calculate(numA, numB rune) rune {
	return c.c.Execute(numA, numB)
}

func main() {
	a := Calcul{c: Add{}}
	b := Calcul{c: Minus{}}

	fmt.Println(a.Calculate(10, 5))
	fmt.Println(b.Calculate(10, 5))
}
