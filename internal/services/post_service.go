package services

import (
	"encoding/json"
	"example.com/first/cmd/app/internal/database"
	"example.com/first/cmd/app/internal/request_structs"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllPosts() []byte {
	posts, _ := json.Marshal(database.GetAllPosts())
	return posts
}

func StorePost(postData *request_structs.StorePostRequest) bool {
	if success := database.ValidatePostData(postData.Title, postData.Content); !success {
		return false
	}
	database.StorePost(postData.Title, postData.Content)
	return true
}

func DeletePost(postId int) bool {
	if err := database.DeletePost(postId); err != nil {
		return false
	}
}
