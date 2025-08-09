package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patil-prathamesh/Blogify/database"
	"github.com/patil-prathamesh/Blogify/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePost(c *gin.Context) {
	var post models.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := validate.Struct(post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation error"})
		return
	}

	authorId := c.GetString("id")

	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()
	post.AuthorId, _ = primitive.ObjectIDFromHex(authorId)

	inserted, err := database.GetPostsCollection().InsertOne(context.Background(), post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message":     "post created",
		"inserted_id": inserted.InsertedID,
		"post":        post,
	})

}

func UpdatePost(c *gin.Context) {

}

func DeletePost(c *gin.Context) {

}

func DeleteAllPosts(c *gin.Context) {

}

func ListAllPosts(c *gin.Context) {
	
}
