package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
ield	Type	Description
_id	ObjectId	Unique identifier for the user.
name	string	Full name of the user.
email	string	Email address of the user.
password	string	Hashed password for authentication.
profile_pic	string	URL to the user's profile picture.
bio	string	Short biography or description of the user.
created_at	datetime	Timestamp when the user was created.
updated_at	datetime	Timestamp when the user was last updated
*/
type User struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName  string             `json:"first_name" validate:"required"`
	LastName   string             `json:"last_name" validate:"required"`
	Email      string             `json:"email" validate:"required"`
	Password   string             `json:"password" validate:"required"`
	ProfilePic string             `json:"profile_pic"`
	Bio        string             `json:"bio"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
}

/*
_id	ObjectId	Unique identifier for the comment.
post_id	ObjectId	Reference to the blog post being commented on.
author_id	ObjectId	Reference to the user who wrote the comment.
content	string	Content of the comment.
created_at	datetime	Timestamp when the comment was created.

4.2 likes Collection
If you want to track likes individually, you can create a likes collection.

Field	Type	Description
_id	ObjectId	Unique identifier for the like.
post_id	ObjectId	Reference to the blog post being liked.
user_id	ObjectId	Reference to the user who liked the post.
created_at	datetime	Timestamp when the like was created.
*/
