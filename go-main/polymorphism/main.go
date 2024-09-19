package main

import "fmt"

type IShape interface {
	area() float64
}

func printArea(shapes ...IShape) {
	for _, shape := range shapes {
		fmt.Println("Alan = ", shape.area())
	}
}

type Triangle struct {
	a float64
	h float64
}

func (t Triangle) area() float64 {
	return (t.a * t.h) / 2
}

type Square struct {
	a float64
}

func (s Square) area() float64 {
	return s.a * s.a
}

type Rectangle struct {
	a, b float64
}

func (r Rectangle) area() float64 {
	return r.a * r.b
}

func main() {

	t := Triangle{4, 10}
	s := Square{5}
	r := Rectangle{5, 10}

	printArea(t, s, r)

}
