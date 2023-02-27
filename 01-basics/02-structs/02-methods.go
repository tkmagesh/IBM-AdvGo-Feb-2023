package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %.2f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discountPercentage float32) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}

func main() {

	var pen Product
	pen = Product{
		Id:   100,
		Name: "Pen",
		Cost: 9.99,
	}
	fmt.Println(pen.Format())

	pencil := &Product{
		Id:   101,
		Name: "Pencil",
		Cost: 4.99,
	}
	fmt.Println("Accessing the members of a struct pointer")
	fmt.Println(pencil.Id, pencil.Name, pencil.Cost)
	fmt.Println(pencil.Format())
	fmt.Println("After applying 10% discount")
	pencil.ApplyDiscount(10)
	fmt.Println(pencil.Format())
}
