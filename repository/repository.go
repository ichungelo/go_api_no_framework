package repository

import (
	"api_go_no_framework/books"
	"api_go_no_framework/entity"
	"log"
)

func GetBooks() []entity.BookEntity {
	return books.Bookshelf
}

func AddBook(book entity.BookEntity) {
	books.Bookshelf = append(books.Bookshelf, book)
	log.Println(books.Bookshelf)
}

func UpdateBook(book entity.BookEntity) {
	data := GetBooks()

	for i := 0; i < len(data); i++ {
		if data[i].BookId == book.BookId {
			data[i] = book
		}
	}
}
