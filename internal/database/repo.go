package database

import (
	"example.com/first/cmd/app/internal/request_structs"
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

func DeletePost(postId int) (bool, error) {
	sql, _ := DB.Prepare("DELETE FROM posts WHERE id = ?")
	result, err := sql.Exec(postId)
	if err != nil {
		return false, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	} else if rowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func UpdatePost(postId int, postData *request_structs.PostRequest) (bool, error) {
	sql, _ := DB.Prepare("UPDATE posts SET title = ?, content = ? WHERE id = ?")
	_, err := sql.Exec(postData.Title, postData.Content, postId)
	if err != nil {
		return false, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	} else if rowsAffected == 0 {
		return false, nil
	}

	return true, nil
}
