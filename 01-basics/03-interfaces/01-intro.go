package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Height + r.Width)
}

/*
func PrintArea(x interface {
	Area() float32
}) {
	fmt.Println("Area : ", x.Area())
}
*/

type AreaFinder interface {
	Area() float32
}

func PrintArea(x AreaFinder) {
	fmt.Println("Area : ", x.Area())
}

func PrintPerimeter(x interface {
	Perimeter() float32
}) {
	fmt.Println("Perimeter :", x.Perimeter())
}

/* interface composition */

/*
func PrintShape(x interface {
	interface {
		Area() float32
	}
	interface {
		Perimeter() float32
	}
}) {
	PrintArea(x)
	PrintPerimeter(x)
}
*/

type PerimeterFinder interface {
	Perimeter() float32
}

type Shape interface {
	AreaFinder
	PerimeterFinder
}

func PrintShape(x Shape) {
	PrintArea(x)
	PrintPerimeter(x)
}

func main() {
	c := Circle{Radius: 12}
	// fmt.Println("Area : ", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintShape(c)

	r := Rectangle{Height: 10, Width: 12}
	// fmt.Println("Area : ", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShape(r)
}
