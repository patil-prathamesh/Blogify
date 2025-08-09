package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
_id	ObjectId	Unique identifier for the post.
title	string	Title of the blog post.
content	string	Main content of the blog post.
author_id	ObjectId	Reference to the user who created the post.
tags	array of strings	Tags or categories associated with the post.
likes	number	Number of likes the post has received.
comments	number	Number of comments on the post.
created_at	datetime	Timestamp when the post was created.
updated_at	datetime	Timestamp when the post was last updated.
*/
type Post struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title"`
	Content   string             `json:"content"`
	AuthorId  primitive.ObjectID `json:"author_id" bson:"author_id"`
	Tags      []string           `json:"tags"`
	Likes     int                `json:"likes"`
	Comments  int                `json:"comments"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}
