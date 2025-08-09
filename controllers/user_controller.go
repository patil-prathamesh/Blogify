package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/patil-prathamesh/Blogify/database"
	"github.com/patil-prathamesh/Blogify/helpers"
	"github.com/patil-prathamesh/Blogify/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	err := validate.Struct(req)	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "validation error"})
		return
	}

	// find user by email
	var user models.User
	err = database.GetUsersCollection().FindOne(context.Background(), bson.M{"email": req.Email}).Decode(&user)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// compare password

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, refreshtoken, err := helpers.GenerateAllTokens(user.ID.Hex(), user.FirstName, user.LastName, user.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not able to generate tokens"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  token,
		"refresh_token": refreshtoken,
		"user":          user,
	})
}

func SignUp(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	err := validate.Struct(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "validation error"})
		return
	}
	
	count , _ := database.GetUsersCollection().CountDocuments(context.Background(), bson.M{"email": user.Email})

	if count >= 1 {
		c.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
		return
	}

	user.Password = HashPassword(user.Password)

	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	token, refreshToken, err := helpers.GenerateAllTokens(user.ID.Hex(), user.FirstName, user.LastName, user.Email)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	inserted, err := database.GetUsersCollection().InsertOne(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":       "User created",
		"inserted id":   inserted.InsertedID,
		"token":         token,
		"refresh_token": refreshToken,
	})
}

func HashPassword(plainPassword string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(plainPassword), 12)
	return string(hashedPassword)
}
