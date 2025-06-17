package domain

import "time"

type Post struct {
	Id        string    `json:"id" bson:"_id,omitempty"`
	Title     string    `json:"title" bson:"title"`
	Content   string    `json:"content" bson:"content"`
	AuthorID  string    `json:"author_id" bson:"author_id"`
	Tags      []string  `json:"tags" bson:"tags"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
