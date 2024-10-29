package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	address string
	router  *mux.Router
}

func New(address string, router *mux.Router) *api {
	return &api{
		address: address,
		router:  router,
	}
}

func (api *api) ListenAndServe() {
	if err := http.ListenAndServe(api.address, api.router); err != nil {
		log.Fatal(err.Error())
	}
}
