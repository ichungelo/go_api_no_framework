package handler

import (
	"api_go_no_framework/entity"
	"api_go_no_framework/repository"
	"api_go_no_framework/transport"
	"api_go_no_framework/usecase"
	"api_go_no_framework/utils"
	"encoding/json"
	"net/http"
)

//? GET ALL BOOKS
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data := repository.GetBooks()
		res, err := json.Marshal(data)

		if err != nil {
			message := []byte(err.Error())
			utils.SetJSONRes(w, message, http.StatusInternalServerError)
			return
		}

		utils.SetJSONRes(w, res, http.StatusOK)
		return
	}
	errorBadRequest(w, r)
}

//? POST BOOK
func PostBook(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		req := transport.RequestBook{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			message := []byte(err.Error())
			utils.SetJSONRes(w, message, http.StatusInternalServerError)
			return
		}

		book := usecase.AddBookUsecase(req)

		if !(book.Title != "" && book.Year != 0 && book.Author != "" && book.Summary != "") {//handler
			res, _ := json.Marshal(usecase.ResponseBook(false, "Bad Request"))
			message := []byte(res)
			utils.SetJSONRes(w, message, http.StatusBadRequest)
			return
		}

		repository.AddBook(book)
		res, _ := json.Marshal(usecase.ResponseBook(true, "Book Added"))
		message := []byte(res)
		utils.SetJSONRes(w, message, http.StatusCreated)
		return
	}
	errorBadRequest(w, r)
}

//? UPDATE BOOK BY ID
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		req := entity.BookEntity{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			message := []byte(err.Error())
			utils.SetJSONRes(w, message, http.StatusInternalServerError)
			return
		}

		if !(req.BookId != "" && req.Title != "" && req.Year != 0 && req.Author != "" && req.Summary != "") {
			res, _ := json.Marshal(usecase.ResponseBook(false, "Bad Request"))
			message := []byte(res)
			utils.SetJSONRes(w, message, http.StatusBadRequest)
			return
		}
		repository.UpdateBook(req)

		res, _ := json.Marshal(usecase.ResponseBook(true, "Book Updated"))
		message := []byte(res)
		utils.SetJSONRes(w, message, http.StatusAccepted)
		return
	}
	errorBadRequest(w, r)
}

//! ERROR BAD REQUEST
func errorBadRequest(w http.ResponseWriter, r *http.Request) {
	data := transport.ResponseBook{
		Success: false,
		Message: "Bad Request",
	}
	res, err := json.Marshal(data)
	if err != nil {
		message := []byte(err.Error())
		utils.SetJSONRes(w, message, http.StatusInternalServerError)
	}
	message := []byte(res)

	utils.SetJSONRes(w, message, http.StatusBadRequest)
}

//! ERROR NOT FOUND
func ErrorNotFound(w http.ResponseWriter, r *http.Request) {
	data := transport.ResponseBook{
		Success: false,
		Message: "Not Found",
	}
	res, err := json.Marshal(data)
	if err != nil {
		message := []byte(err.Error())
		utils.SetJSONRes(w, message, http.StatusInternalServerError)
	}
	message := []byte(res)

	utils.SetJSONRes(w, message, http.StatusNotFound)
}
