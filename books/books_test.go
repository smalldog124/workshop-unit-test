package books_test

import (
	"potter/books"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewBook_Input_Harry_Potter_Part_1_And_Price_8_EUR(t *testing.T) {
	var book books.Books
	book.NewBook("Harry Potter Part 1", EUR(8))

	assert.Equal(t, "Harry Potter Part 1", book.Name)
	assert.Equal(t, EUR(8), book.Price)
}

func EUR(price int) int {
	return price * 100
}
