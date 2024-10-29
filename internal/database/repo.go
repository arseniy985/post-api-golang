package database

import (
	"log"
)

func GetAllPosts() []Post {
	rows, err := DB.Query("SELECT * FROM posts")
	if err != nil {
		log.Fatal(err.Error())
	}
	var posts []Post
	for rows.Next() {
		var post Post
		posts = append(posts, post)
	}
	return posts
}
