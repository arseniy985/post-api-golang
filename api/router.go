package api

import (
	"example.com/first/cmd/app/internal/handlers"
	"net/http"
)

func (api *api) AddRoutes() {
	api.router.HandleFunc("/posts", handlers.GetPostsHandler).Methods(http.MethodGet)
	api.router.HandleFunc("/post/store", handlers.StorePostHandler).Methods(http.MethodPost)
	api.router.HandleFunc("/post/{id}", handlers.DeletePostHandler).Methods(http.MethodDelete)
	api.router.HandleFunc("/post/{id}", handlers.UpdatePostHandler).Methods(http.MethodPatch)
}
