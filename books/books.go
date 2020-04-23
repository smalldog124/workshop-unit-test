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

type Basket struct {
	Books         SelectedBook
	TotalPrice    int
	DiscountPrice int
	NetPrice      int
}

type SelectedBook struct {
	ID    int
	Name  string
	Price int
	Stock int
}
