package repository

import (
	"api_go_no_framework/books"
	"api_go_no_framework/entity"
	"api_go_no_framework/transport"
	"api_go_no_framework/utils"
	"errors"
	"log"
)

type BookRepository interface {
	GetBooks() []entity.BookEntity
	AddBook(req transport.RequestBook) (transport.ResponseBook, error)
}

func GetBooks() []entity.BookEntity{
	return books.Bookshelf
}

func AddBook(req transport.RequestBook) (transport.ResponseBook, error) {
	id := utils.IdGEnerator(16)
	if !(req.Title != "" && req.Year != 0 && req.Author != "" && req.Summary != ""){
		res := transport.ResponseBook{
			Success: false,
			Message: "Bad Request",
		}
		return res, errors.New("invalid input")
	}
	
	data := entity.BookEntity{
		BookId:  id,
		Title:   req.Title,
		Year:    req.Year,
		Author:  req.Author,
		Summary: req.Summary,
	}

	books.Bookshelf = append(books.Bookshelf, data)
	log.Println(books.Bookshelf)

	res := transport.ResponseBook{
		Success: true,
		Message: "Book Added",
	}
	return res, nil
}
