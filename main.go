package main

import (
	"api_go_no_framework/handler"
	"fmt"
	"net/http"
)

func main() {
	port := ":3000"

	http.HandleFunc("/", handler.ErrorNotFound)
	http.HandleFunc("/get/books", handler.GetAllBooks)
	http.HandleFunc("/post/books", handler.PostBook)
	http.HandleFunc("/put/books", handler.UpdateBook)
	http.HandleFunc("/get/book", handler.GetBookById)

    fmt.Printf("Start web server at http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
