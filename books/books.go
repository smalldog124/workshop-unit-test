package books

import (
	"math/rand"
)

type Books struct {
	ID    int
	Name  string
	Price int
	Stock int
}

func (b *Books) NewBook(name string, price int) {
	b.ID = rand.Int()
	b.Name = name
	b.Price = price
	b.Stock = 10
}
