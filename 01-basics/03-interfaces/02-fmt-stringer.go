package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

/* fmt.Stringer interface implementation */
func (p Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %.2f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discountPercentage float32) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}

/* composition (aka inheritence) */
type PerishableProduct struct {
	Product
	Expiry string
}

/* fmt.Stringer interface implementation */
func (pp PerishableProduct) String() string { /* overriding the Product.Format() method */
	// return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.Expiry)
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %.2f, Expiry = %q", pp.Id, pp.Name, pp.Cost, pp.Expiry)
}

func main() {

	grapes := PerishableProduct{
		Product: Product{
			Id:   100,
			Name: "Grapes",
			Cost: 50,
		},
		Expiry: "2 Days",
	}
	// fmt.Println(grapes.Format())
	fmt.Println(grapes)
	fmt.Println("After applying 10% discount")
	grapes.ApplyDiscount(10)
	// fmt.Println(grapes.Format())
	fmt.Println(grapes)
}
