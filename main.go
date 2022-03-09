package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := ":3000"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := []byte(
			`{"status": "pong"}`,
		)
		w.Write(message)
	})

    fmt.Printf("Start web server at http://localhost%s", port)
	http.ListenAndServe(port, nil)
}
