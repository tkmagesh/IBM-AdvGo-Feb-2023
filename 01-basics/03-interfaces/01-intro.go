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

type Rectangle struct {
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
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

func main() {
	c := Circle{Radius: 12}
	// fmt.Println("Area : ", c.Area())
	PrintArea(c)
	PrintPerimeter(c)

	r := Rectangle{Height: 10, Width: 12}
	// fmt.Println("Area : ", r.Area())
	PrintArea(r)
	PrintPerimeter(r)
}
