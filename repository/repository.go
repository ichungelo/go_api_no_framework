package repository

import (
	"api_go_no_framework/books"
	"api_go_no_framework/entity"
	"errors"
	"log"
)

func GetBooks() []entity.BookEntity {
	return books.Bookshelf
}

func GetBookById(id string) (*entity.BookEntity, error) {
	data := GetBooks()
	result := entity.BookEntity{}
	count := 0
	for i := 0; i < len(data); i++ {
		if data[i].BookId == id {
			result = data[i]
			break
		}
		count++
	}

	if count == len(data) {
		return nil, errors.New("not found")
	}

	return &result, nil
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

func DeleteBook(id string)  {
	data :=GetBooks()
	newData := []entity.BookEntity{}
	for i := 0; i < len(data); i++ {
		if data[i].BookId == id {
			continue
		}
		newData = append(newData, data[i])
	}

	books.Bookshelf = newData
}
