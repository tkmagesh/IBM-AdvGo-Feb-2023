package main

import "fmt"

func main() {
	/*
		pen := struct {
			Id   int
			Name string
			Cost float32
		}{
			Id:   100,
			Name: "Pen",
			Cost: 9.99,
		}
	*/
	type Product struct {
		Id   int
		Name string
		Cost float32
	}

	var pen Product
	pen = Product{
		Id:   100,
		Name: "Pen",
		Cost: 9.99,
	}
	fmt.Println(pen)
}
