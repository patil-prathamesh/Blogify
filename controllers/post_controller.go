package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patil-prathamesh/Blogify/database"
	"github.com/patil-prathamesh/Blogify/models"
	"go.mongodb.org/mongo-driver/bson"
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
	postId := c.Param("post_id")
	id := c.GetString("id")
	userId, _ := primitive.ObjectIDFromHex(id)

	count, _ := database.GetUsersCollection().CountDocuments(context.Background(), bson.M{"_id": userId})

	if count < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unauthorized access"})
		return
	}

	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	err := validate.Struct(post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation error"})
		return
	}

	pId, _ := primitive.ObjectIDFromHex(postId)
	filter := bson.M{"_id": pId}
	update := bson.M{"$set": bson.M{"title": post.Title, "content": post.Content, "tags": post.Tags, "updated_at": time.Now()}}

	result, err := database.GetPostsCollection().UpdateOne(context.Background(), filter, update)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "post not updated"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": result})
}

func DeletePost(c *gin.Context) {
	id := c.GetString("id")
	userId, _ := primitive.ObjectIDFromHex(id)

	count, _ := database.GetUsersCollection().CountDocuments(context.Background(), gin.H{"_d": userId})

	if count < 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unauthorized access"})
		return
	}

	postId := c.Param("post_id")
	fmt.Print(postId, " ------")
	objectPostId, _ := primitive.ObjectIDFromHex(postId)
	fmt.Print(objectPostId, " ------")

	result, err := database.GetPostsCollection().DeleteOne(context.Background(), bson.M{"_id": objectPostId})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "post not deleted"})
		return
	}

	if result.DeletedCount > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "post deleted successfully.", "result": result})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "post not availabe.", "result": result})
		return
	}

}

func DeleteAllPosts(c *gin.Context) {
	authorId := c.GetString("id")
	objectAuthorId, _ := primitive.ObjectIDFromHex(authorId)
	count , _ := database.GetUsersCollection().CountDocuments(context.Background(), gin.H{"_id": objectAuthorId})

	if count < 1{
		c.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized access"})
		return
	}

	result, _ := database.GetPostsCollection().DeleteMany(context.Background(), gin.H{"author_id": objectAuthorId})

	if result.DeletedCount > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "posts deleted", "count": result.DeletedCount})
		return
	}

}

func ListAllPosts(c *gin.Context) {
	var posts []models.Post
	id := c.GetString("id")

	authorId, _ := primitive.ObjectIDFromHex(id)
	cursor, _ := database.GetPostsCollection().Find(context.Background(), bson.M{"author_id": authorId})

	cursor.All(context.Background(), &posts)
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}
