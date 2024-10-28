package api

import (
	"net/http"
)

type api struct {
	address string
	r       *http.ServeMux
}

func New(address string, r *http.ServeMux) *api {
	return &api{
		address: address,
		r:       r,
	}
}

func (api *api) ListenAndServe() error {
	return http.ListenAndServe(api.address, api.r)
}
