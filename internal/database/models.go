package database

import "time"

type Post struct {
	id        uint
	title     string
	content   string
	createdAt time.Time
	updatedAt time.Time
}
