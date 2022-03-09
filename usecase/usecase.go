package usecase

import (
	"api_go_no_framework/entity"
	"api_go_no_framework/transport"
	"api_go_no_framework/utils"
)

func AddBookUsecase (data transport.RequestBook) (entity.BookEntity) {
	id := utils.IdGEnerator(16)
	result := entity.BookEntity{
		BookId: id,
		Title: data.Title,
		Year: data.Year,
		Author: data.Author,
		Summary: data.Summary,
	}

	return result
}

func ResponseBook (success bool, message string) (transport.ResponseBook) {
	result := transport.ResponseBook{
		Success: success,
		Message: message,
	}

	return result
}