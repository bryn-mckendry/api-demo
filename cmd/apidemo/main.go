package main

import (
	"encoding/json"
	"net/http"
)

type Payload struct {
	Message string `json:"Message"`
}

func main() {
	http.HandleFunc("/", handleResponse)
	http.ListenAndServe(":8080", nil)
}

func handleResponse(writer http.ResponseWriter, req *http.Request) {
	message := Payload{"Hello world!"}
	json.NewEncoder(writer).Encode(message)
}
