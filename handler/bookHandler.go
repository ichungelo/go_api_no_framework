package handler

import (
	"api_go_no_framework/repository"
	"api_go_no_framework/transport"
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

		data, err := repository.AddBook(req)
		if err != nil {
			res, _ := json.Marshal(data)
			message := []byte(res)
			utils.SetJSONRes(w, message, http.StatusBadRequest)
			return
		}

		res, _ := json.Marshal(data)
		message := []byte(res)
		utils.SetJSONRes(w, message, http.StatusCreated)
	}
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
