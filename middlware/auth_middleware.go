package middlware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/patil-prathamesh/Blogify/database"
	"github.com/patil-prathamesh/Blogify/helpers"
	"go.mongodb.org/mongo-driver/bson"
)

func Authenticate(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	token := ""
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		token = authHeader[7:]
	}

	if token == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unauthorized access"})
		c.Abort()
		return
	}

	claims, err := helpers.ValidateToken(token)
	if err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		c.Abort()
		return
	}

	c.Set("email", claims.Email)
	c.Set("first_name", claims.FirstName)
	c.Set("last_name", claims.LastName)
	c.Set("id", claims.ID)
	c.Next()
}

func IsValid(c *gin.Context) {
	userEmail := c.GetString("email")
	count, _ := database.GetUsersCollection().CountDocuments(context.Background(), bson.M{"email": userEmail})

	fmt.Println(count)

	if count < 1 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		c.Abort()
		return
	}

	c.Next()

}
