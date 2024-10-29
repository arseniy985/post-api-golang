package api

import (
	"example.com/first/cmd/app/internal/database"
	"net/http"
)

func (api *api) AddRoutes() {
	api.router.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Саси письку"))
	})
	api.router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		posts := string(database.GetAllPosts())
		w.Write([]byte(posts))
	})
}
