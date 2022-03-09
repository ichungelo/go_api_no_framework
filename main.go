package main

import (
	"api_go_no_framework/utils"
	"fmt"
	"net/http"
)

func main() {
	port := ":3000"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := []byte(
			`{"status": "pong"}`,
		)
		utils.SetJSONRes(w, message, http.StatusOK)
	})

    fmt.Printf("Start web server at http://localhost%s", port)
	http.ListenAndServe(port, nil)
}
