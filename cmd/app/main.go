package main

import (
	"log"
	"net/http"

	"example.com/first/api"
)

func main() {
	api := api.New("localhost:8080", http.NewServeMux())

	api.AddRoutes()

	if err := api.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
