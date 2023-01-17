package auth

import (
	"net/http"
	"os"
)

var valid_user, valid_pass string

func init() {
	valid_user = os.Getenv("ANIMALS_API_USERNAME")
	valid_pass = os.Getenv("ANIMALS_API_PASSWORD")

	if valid_user == "" || valid_pass == "" {
		panic("ANIMALS_API_USERNAME and ANIMALS_API_PASSWORD environment variables not set")
	}
}

func Authorize(req *http.Request) bool {
	user, pass, _ := req.BasicAuth()
	if user == valid_user && pass == valid_pass {
		return true
	}
	return false
}
