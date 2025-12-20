package interfaces

import (
	"fmt"
	"math"
)

type Square struct {
	Length, Area int
}
type Circle struct {
	Radius float64
	Area   float64
}

func (s *Square) calcArea() {
	s.Area = s.Length * s.Length
}
func (c *Circle) calcArea() {
	c.Area = math.Pi * c.Radius * c.Radius
}

func NewSquare(len int) *Square {
	newS := &Square{Length: len}
	newS.calcArea()
	return newS
}
func NewCircle(rad float64) *Circle {
	newC := &Circle{Radius: rad}
	newC.calcArea()
	return newC
}

type Shape interface {
	calcArea()
}

func PrintShape(s Shape) {
	fmt.Printf("-------\n")
	fmt.Printf("Shape: %T\n", s)
	fmt.Printf("Values: %+v\n", s)
	fmt.Printf("-------\n\n")
}
