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

func StorePost(postData *request_structs.PostRequest) bool {
	if success := database.ValidatePostData(postData.Title, postData.Content); !success {
		return false
	}
	database.StorePost(postData.Title, postData.Content)
	return true
}

func DeletePost(postId int) (bool, error) {
	res, err := database.DeletePost(postId)
	if err != nil {
		return false, err
	} else if res == false && err == nil {
		return false, nil
	}
	return true, nil
}

func UpdatePost(postId int, postData *request_structs.PostRequest) bool {
	if err := database.UpdatePost(postId, postData); err != nil {
		return false
	}
	return true
}
