package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateServer() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		for k, v := range r.URL.Query() {
			fmt.Fprintf(w, "%v: %v\n", k, v)
		}
	})
	return r
}
