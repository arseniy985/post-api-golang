package api

import (
	"github.com/gorilla/mux"
)

func (api *api) AddRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/posts")
}
