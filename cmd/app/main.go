package main

import (
	"example.com/first/cmd/app/api"
	"example.com/first/cmd/app/internal/database"
	"github.com/gorilla/mux"
)

func main() {
	api := api.New("localhost:8080", mux.NewRouter())
	db := database.NewDatabase(
		"mysql", "localhost", "goDatabase",
		"root", "root", 3306,
	)
	db.OpenConnection()
	defer db.CloseConnection()

	api.AddRoutes()

	api.ListenAndServe()
}
