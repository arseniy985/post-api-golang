package handlers

import (
	"encoding/json"
	"example.com/first/cmd/app/internal/request_structs"
	"example.com/first/cmd/app/internal/services"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(services.GetAllPosts())
}
func StorePostHandler(w http.ResponseWriter, r *http.Request) {
	var postData request_structs.StorePostRequest
	body, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		http.Error(w, "", http.StatusBadRequest)
	}
	defer r.Body.Close()
	if err := json.Unmarshal(body, &postData); err != nil {
		http.Error(w, "", http.StatusBadRequest)
	}
	if success := services.StorePost(&postData); !success {
		http.Error(w, "Uncorrected post data", http.StatusBadRequest)
	}
	w.WriteHeader(200)
}
func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi((mux.Vars(r))["id"])
	if err != nil {
		http.Error(w, "Id must be integer", http.StatusBadRequest)
	}
	if !services.DeletePost(id) {
		http.Error(w, "", http.StatusInternalServerError)
	}
	w.WriteHeader(200)
}
