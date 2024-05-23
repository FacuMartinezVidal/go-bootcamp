package product

import (
	"errors"
	"fmt"
)

type Product struct {
	name        string
	price       float64
	description string
}

func New(name string, price float64, description string) (*Product, error) {
	if name == "" || price < 0 || description == "" {
		return nil, errors.New("Invalid input")
	}
	return &Product{
		name:        name,
		price:       price,
		description: description,
	}, nil
}

func (p *Product) String() string {
	return fmt.Sprintf("Name: %s, Price: %f, Description: %s", p.name, p.price, p.description)
}
