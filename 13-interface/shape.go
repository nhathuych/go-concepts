package main

import "math"

type Shape interface {
	Area() float64
	// CalculatePerimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
