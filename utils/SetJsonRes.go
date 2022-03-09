package utils

import "net/http"

func SetJSONRes(w http.ResponseWriter, message []byte, httpCode int)  {
	w.Header().Set("Content-type", "application/JSON")
	w.WriteHeader(httpCode)
	w.Write(message)
}