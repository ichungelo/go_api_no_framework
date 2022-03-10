package main

import (
	"api_go_no_framework/handler"
	"api_go_no_framework/errors"
	"fmt"
	"net/http"
)

func main() {
	port := ":3000"

	http.HandleFunc("/", errors.ErrorNotFound)
	http.HandleFunc("/get/books/all", handler.GetAllBooks)
	http.HandleFunc("/post/books", handler.PostBook)
	http.HandleFunc("/put/books", handler.UpdateBook)
	http.HandleFunc("/get/books", handler.GetBookById)
	http.HandleFunc("/delete/books", handler.DeleteBook)

    fmt.Printf("Start web server at http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
