/*
Add a new method to the Shape interface called perimeter which calculates the perimeter of a shape.
Implement the method for Circle and Rectangle.
*/
package main

import (
	"math"
	"fmt"
)

func main() {
	c := Circle{x: 0, y: 0, r: 5}
	r := Rectangle{x1:1, x2:2, y1:4, y2:8}
	fmt.Println("Circle area :", c.area())
	fmt.Println("Circle perimeter :", c.perimeter())
	fmt.Println("Rectangle area :", r.area())
	fmt.Println("Rectangle perimeter :", r.perimeter())
	fmt.Println("Total area : ",totalArea(&c, &r))
	fmt.Println("Total perimeter : ",totalPerimeter(&c, &r))
}

type Shape interface {
	area() float64
	perimeter() float64
}
// create a custom type definition, just like creating a class in java
type Circle struct {
	x, y, r float64
}
/*
create a function that will be bound to a struct (in go this is called method).
in java this is like method definition for class, unlike java go define the method outside of struct and bind it
*/
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return math.Pi * 2 * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}
func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * l + 2 * w
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func totalPerimeter(shapes ...Shape) float64 {
	var perimeter float64
	for _, s := range shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}