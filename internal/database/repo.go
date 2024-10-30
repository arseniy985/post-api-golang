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
		rows.Scan(&post.ID, &post.Title, &post.Content)
		posts = append(posts, post)
	}
	defer rows.Close()
	return posts
}

func StorePost(title, content string) error {
	sql, _ := DB.Prepare("INSERT INTO posts (title, content) VALUES (?, ?)")
	_, execErr := sql.Exec(title, content)
	if execErr != nil {
		return execErr
	}
	return nil
}

func DeletePost(postId int) error {
	sql, _ := DB.Prepare("DELETE FROM posts WHERE id = ?")
	_, err := sql.Exec(postId)
	if err != nil {
		return err
	}
	return nil
}
