package main

import (
	"api_go_no_framework/handler"
	"api_go_no_framework/utils"
	"fmt"
	"net/http"
)

func main() {
	port := ":3000"

	http.HandleFunc("/", handler.ErrorNotFound)
	http.HandleFunc("/get/books", handler.GetAllBooks)

	fmt.Println(utils.IdGEnerator(16))
    fmt.Printf("Start web server at http://localhost%s", port)
	http.ListenAndServe(port, nil)
}
