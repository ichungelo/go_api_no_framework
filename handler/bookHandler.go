package handler

import (
	"api_go_no_framework/books"
	"api_go_no_framework/transport"
	"api_go_no_framework/utils"
	"encoding/json"
	"net/http"
)

//? GET ALL BOOKS
func GetAllBooks(w http.ResponseWriter, r *http.Request)  {
	if r.Method == http.MethodGet {
		res, err := json.Marshal(books.Bookshelf)

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
