package handler

import (
	"api_go_no_framework/entity"
	"api_go_no_framework/repository"
	"api_go_no_framework/transport"
	"api_go_no_framework/usecase"
	"api_go_no_framework/utils"
	"api_go_no_framework/errors"
	"encoding/json"
	"log"
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
	errors.ErrorNotFound(w, r)
}

//? GET BOOK BY ID
func GetBookById(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		req := transport.RequestBookId{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			message := []byte(err.Error())
			utils.SetJSONRes(w, message, http.StatusInternalServerError)
			return
		}

		id := req.BookId

		if id == "" {
			res, _ := json.Marshal(usecase.ResponseBook(false, "Bad Request"))
			utils.SetJSONRes(w, res, http.StatusBadRequest)
			return
		}

		data, err := repository.GetBookById(id)

		if err != nil {
			errors.ErrorNotFound(w, r)
			return
		}

		res, _ := json.Marshal(data)
		utils.SetJSONRes(w, res, http.StatusOK)
		return
	}
	errors.ErrorNotFound(w, r)
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
		
		if !(req.Title != "" && req.Year != 0 && req.Author != "" && req.Summary != "") { //handler
			errors.ErrorInput(w, req, r)
			return
		}

		book := usecase.AddBookUsecase(req)
		
		repository.AddBook(book)
		res, _ := json.Marshal(usecase.ResponseBook(true, "Book Added"))
		utils.SetJSONRes(w, res, http.StatusCreated)
		return
	}
	errors.ErrorNotFound(w, r)
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
			log.Println("Error")
			errors.ErrorBadRequest(w, r)
			return
		}
		repository.UpdateBook(req)

		res, _ := json.Marshal(usecase.ResponseBook(true, "Book Updated"))
		utils.SetJSONRes(w, res, http.StatusAccepted)
		return
	}
	errors.ErrorNotFound(w, r)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		req := transport.RequestBookId{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			message := []byte(err.Error())
			utils.SetJSONRes(w, message, http.StatusInternalServerError)
			return
		}
		
		id := req.BookId

		if id == "" {
			res, _ := json.Marshal(usecase.ResponseBook(false, "Bad Request"))
			utils.SetJSONRes(w, res, http.StatusBadRequest)
			return
		}

		repository.DeleteBook(id)

		res, _ := json.Marshal(usecase.ResponseBook(true, "Book Deleted"))
		utils.SetJSONRes(w, res, http.StatusOK)
		return
	}
	errors.ErrorNotFound(w, r)
}
