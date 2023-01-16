package auth

import (
	"net/http"
	"os"
)

var valid_user = os.Getenv("ANIMALS_API_USERNAME")
var valid_pass = os.Getenv("ANIMALS_API_PASSWORD")

func Authorize(req *http.Request) bool {
	user, pass, _ := req.BasicAuth()
	if user == valid_user && pass == valid_pass {
		return true
	}
	return false
}
