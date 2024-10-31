package handlers

import (
	helper "example.com/first/cmd/app/internal"
	"example.com/first/cmd/app/internal/request_structs"
	"example.com/first/cmd/app/internal/services"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(services.GetAllPosts())
}
func StorePostHandler(w http.ResponseWriter, r *http.Request) {
	var postData request_structs.PostRequest
	if err := helper.GetBodyData(&postData, r); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	if success := services.StorePost(&postData); !success {
		http.Error(w, "Uncorrected post data", http.StatusBadRequest)
		return
	}
	w.WriteHeader(200)
}
func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi((mux.Vars(r))["id"])
	if err != nil {
		http.Error(w, "Id must be a number", http.StatusBadRequest)
		return
	}

	res, err := services.DeletePost(id)
	if res == false && err == nil {
		http.Error(w, "Post not found", http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}
func UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi((mux.Vars(r))["id"])
	if err != nil {
		http.Error(w, "Id must be integer", http.StatusBadRequest)
		return
	}

	var postData request_structs.PostRequest
	if err := helper.GetBodyData(&postData, r); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	if !services.UpdatePost(id, &postData) {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}
