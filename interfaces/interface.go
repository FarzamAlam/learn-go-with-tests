package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}
type Triangle struct {
	Base   float64
	Height float64
}
type obj struct {
	r Rectangle
	c Circle
	s Shape
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Shape interface {
	Area() float64
}

func main() {
	r := Rectangle{12.0, 10.0}
	c := Circle{10.0}
	fmt.Println("Rectangle : ", r.Area())
	fmt.Println("Circle : ", c.Area())

	fmt.Println("------Interface-------")
	var s Shape
	s = Rectangle{10, 20}
	s = Circle{10}
	fmt.Println(s.Area())
	s = Triangle{10, 12}
	fmt.Println(s.Area())
}
