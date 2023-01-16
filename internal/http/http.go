package http

import (
	"apidemo/internal/auth"
	"apidemo/internal/database"
	"apidemo/internal/domain"
	"encoding/json"
	"net/http"
)

func HandleRequest(res http.ResponseWriter, req *http.Request) {
	if !auth.Authorize(req) {
		http.Error(res, "Unauthorized", http.StatusUnauthorized)
		return
	}

	switch req.Method {
	case "GET":
		get(res, req)
	case "POST":
		post(res, req)
	default:
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func get(res http.ResponseWriter, req *http.Request) {
	animals, err := database.GetAnimals()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(animals)
}

func post(res http.ResponseWriter, req *http.Request) {
	var a domain.Animal
	if err := json.NewDecoder(req.Body).Decode(&a); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := database.CreateAnimal(a); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusCreated)
}
