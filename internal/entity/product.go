package entity

type Product struct {
	name  string
	price int
}

func NewProduct(name string, price int) *Product {
	return &Product{
		name:  name,
		price: price,
	}
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) Price() int {
	return p.price
}
