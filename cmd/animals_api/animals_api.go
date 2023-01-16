package animals_api

import (
	api "apidemo/internal/http"
	"net/http"

	_ "github.com/lib/pq"
)

func Run() {
	http.HandleFunc("/", api.HandleRequest)
	http.ListenAndServe(":8080", nil)
}
