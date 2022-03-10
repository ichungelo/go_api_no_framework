package errors

import (
	"api_go_no_framework/transport"
	"api_go_no_framework/usecase"
	"api_go_no_framework/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

//! ERROR BAD REQUEST
func ErrorBadRequest(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal(usecase.ResponseBook(false, "Bad Request"))
	log.Println(res)
	utils.SetJSONRes(w, res, http.StatusBadRequest)
}

//! ERROR NOT FOUND
func ErrorNotFound(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal(usecase.ResponseBook(false, "Not Found"))
	utils.SetJSONRes(w, res, http.StatusNotFound)
}

//!ERROR INPUT
func ErrorInput(w http.ResponseWriter, data transport.RequestBook, r *http.Request) {
	message := []string{}
	
	if data.Title == "" {
		message = append(message, "title")
	}

	if data.Year == 0 {
		message = append(message, "year")
	}

	if data.Author == "" {
		message = append(message, "author")
	}

	if data.Summary == "" {
		message = append(message, "summary")
	}

	result := strings.Join(message, ", ")

	res, _ := json.Marshal(usecase.ResponseBook(false, fmt.Sprintf("You need to input %s", result)))
	utils.SetJSONRes(w, res, http.StatusBadRequest)

}
