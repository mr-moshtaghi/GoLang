package main

import (
	"fmt"
	"math"
)

type Squere struct {
	Length float64
}

func (s *Squere) Aria() float64 {
	return s.Length * s.Length
}

type Circle struct {
	Radius float64
}

func (c *Circle) Aria() float64 {
	// return math.Pi * math.Pow(c.Radius, 2)
	return math.Pi * c.Radius * c.Radius
}

type Shape interface {
	Aria() float64
}

func sumAria(shapes []Shape) float64 {
	total := 0.0

	for _, shape := range shapes {
		total += shape.Aria()
	}
	return float64(total)
}

func main() {
	s := Squere{
		Length: 20,
	}

	fmt.Println(s.Aria())

	c := Circle{
		Radius: 20,
	}

	fmt.Println(c.Aria())

	shapes := []Shape{&s, &c}

	sc := sumAria(shapes)

	fmt.Println(sc)
}
